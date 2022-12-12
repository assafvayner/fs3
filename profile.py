"""cse453 default profile to run DeathStarBench.

This profile is sourced from the cse453 gitlab repo at
https://gitlab.cs.washington.edu/syslab/cse453-cloud-project/. The
entire repo is checked out into /local/repository on each created machine.
"""

import geni.portal as portal
import geni.rspec.pg as pg
import geni.rspec.emulab as emulab
from lxml import etree as ET

### Configuration

# Boilerplate setup
pc = portal.Context()
request = pc.makeRequestRSpec()

# Lab selection
lablist = [
    ('lab0', 'Lab 0'),
    ('lab1', 'Lab 1'),
    ('lab2', 'Lab 2'),
    ('lab3', 'Lab 3'),
    ('lab4', 'Lab 4 - Advanced options available')]
pc.defineParameter("lab", "Select the lab you are working on",
                   portal.ParameterType.STRING, lablist[0], lablist)

# Node type to reserve (lab 1 only)
pc.defineParameter("node_type", "Node type (for lab 1 only)",
                   portal.ParameterType.NODETYPE, "c6525-25g", longDescription="These are known to work: Utah: c6525-25g, c6525-100g; APT: r320; Wisconin: c220g2, c220g1")

# Number of server machines
pc.defineParameter("n_servers", "Number of server machines",
                   portal.ParameterType.INTEGER, 3, advanced=True)

# Number of client machines
pc.defineParameter("n_clients", "Number of client machines",
                   portal.ParameterType.INTEGER, 1, advanced=True)

# Parameter to set virtualized mode or not
modelist = [
    ('default', 'default - Any x86 machine'),
    ('passthru', 'passthru - Only machines that support device pass-through')]
pc.defineParameter("mode", "Select default or device pass-through mode for servers",
                   portal.ParameterType.STRING, modelist[0], modelist, advanced=True)

# Retrieve the values the user specifies during instantiation
params = pc.bindParameters()

# Check parameter validity
if params.n_servers < 1 or params.n_servers > 10:
    pc.reportError(portal.ParameterError("You must choose at least 1 and no more than 10 servers.", ["n_servers"]))
if params.n_clients < 1 or params.n_clients > 3:
    pc.reportError(portal.ParameterError("You must choose at least 1 and no more than 3 clients.", ["n_clients"]))

if params.lab == 'lab0':
        params.n_servers = 3
        params.n_clients = 1
        params.mode = 'default'
elif params.lab == 'lab1':
        params.n_servers = 3
        params.n_clients = 1
        params.mode = 'passthru'
elif params.lab == 'lab2':
        params.n_servers = 1
        params.n_clients = 1
        params.mode = 'default'
elif params.lab == 'lab3':
        params.n_servers = 2
        params.n_clients = 1
        params.mode = 'default'
else:
    if params.lab != 'lab4':
        pc.reportError(portal.ParameterError("Invalid lab selected!", ["lab"]))

# Abort execution if there are any errors, and report them
portal.context.verifyParameters()

# Add rewritten parameters to the manifest
class Parameters(pg.Resource):
    def _write(self, root):
        ns = "{http://www.protogeni.net/resources/rspec/ext/profile-parameters/1}"
        paramXML = "%sdata_item" % (ns,)
        el = ET.SubElement(root,"%sdata_set" % (ns,))

        param = ET.SubElement(el,paramXML, name='n_servers')
        param.text = '%u' % int(params.n_servers)
        param = ET.SubElement(el,paramXML, name='n_clients')
        param.text = '%u' % int(params.n_clients)
        param = ET.SubElement(el,paramXML, name='mode')
        param.text = params.mode
        param = ET.SubElement(el,paramXML, name='node_type')
        param.text = params.node_type

        return el

parameters = Parameters()
request.addResource(parameters)

# Create starfish network topology
mylink = request.Link('mylink')
mylink.Site('undefined')

for i in range(params.n_servers):
    # Create node
    n = request.RawPC('server%u' % i)
    n.disk_image = 'urn:publicid:IDN+emulab.net+image+emulab-ops//UBUNTU20-64-STD'
    iface = n.addInterface('interface-%u' % i)
    if params.mode == "passthru":
        # We know that the AMD machines support device pass-through.
        # XXX: This is restrictive, as other machine types might support it, too. Not clear how to constrain the hardware type to a set of machines, rather than just a single type.
        n.hardware_type = params.node_type
    n.addService(pg.Execute(shell="bash", command="/local/repository/init.sh"))
    mylink.addInterface(iface)

for i in range(params.n_clients):
    # Create node
    n = request.RawPC('client%u' % i)
    n.disk_image = 'urn:publicid:IDN+emulab.net+image+emulab-ops//UBUNTU20-64-STD'
    iface = n.addInterface('interface-%u' % (params.n_servers + i))
    n.addService(pg.Execute(shell="bash", command="/local/repository/init.sh"))
    mylink.addInterface(iface)

# Print the generated rspec
pc.printRequestRSpec(request)
