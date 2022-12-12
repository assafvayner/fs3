#!/bin/bash -x

function geni_get_parameter()
{
    geni-get manifest | xpath -e "rspec/ns0:data_set/ns0:data_item[attribute::name = \"$1\"]/text()" -q
}

# Number of servers
n_servers=`geni_get_parameter n_servers`

# Number of clients
n_clients=`geni_get_parameter n_clients`

# Number of VMs (same as servers)
nvms=$n_servers

# XXX: Example getting old emulab parameters
# n_servers=`geni-get manifest | xpath -e 'rspec/data_set/data_item[attribute::name = "emulab.net.parameter.n_servers"]/text()' -q`

