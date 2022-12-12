#!/bin/bash -x

. config.sh

function get_token()
{
    awk -F, "\$1 == \"server$1\" { print \$2 }" /tmp/tokens.csv
}

# This script runs only on server0
test x`hostname -s` == xserver0 || { echo "This script has to run on server0"; exit 1; }

# Parse commandline
while [[ $# -gt 0 ]]; do
    case "$1" in
	--tcg)
	    # LXC_OPTS='-c raw.qemu.conf="[machine] accel=\"tcg\"" -c raw.qemu="-cpu max"'
      LXC_OPTS=(-c raw.qemu.conf="[machine] accel=\"tcg\"" -c raw.qemu="-cpu max")
	    shift
	    ;;
	-h | --help)
	    echo "Usage: $0 [--tcg]"
	    exit 0
	    ;;
	*)
	    echo "Unexpected option: $1"
	    shift
	    ;;
    esac
done

# Setup master LXD instance
sudo lxd init --preseed <<EOF
config:
  core.https_address: `hostname -s`:8443
networks:
- config:
    bridge.mode: fan
    fan.underlay_subnet: 10.10.1.0/24
  description: ""
  name: lxdfan0
  type: ""
  project: default
storage_pools:
- config:
    size: 20GB
  description: ""
  name: local
  driver: zfs
profiles:
- config: {}
  description: ""
  devices:
    eth0:
      name: eth0
      network: lxdfan0
      type: nic
    root:
      path: /
      pool: local
      type: disk
  name: default
projects: []
cluster:
  server_name: `hostname -s`
  enabled: true
  member_config: []
  cluster_address: ""
  cluster_certificate: ""
  server_address: ""
  cluster_password: ""
  cluster_certificate_path: ""
  cluster_token: ""
EOF

# Generate cluster join tokens
for i in `seq $((n_servers - 1))`; do
    sudo lxc cluster add server$i
done
sudo lxc cluster list-tokens -f csv > /tmp/tokens.csv

# Join all cluster nodes
for i in `seq $((n_servers - 1))`; do
    ssh -oStrictHostKeyChecking=no server$i "sudo lxd init --preseed <<EOF
config: {}
networks: []
storage_pools: []
profiles: []
projects: []
cluster:
  server_name: server$i
  enabled: true
  member_config:
  - entity: storage-pool
    name: local
    key: size
    value: 20GB
    description: '\"size\" property for storage pool \"local\"'
  - entity: storage-pool
    name: local
    key: source
    value: \"\"
    description: '\"source\" property for storage pool \"local\"'
  - entity: storage-pool
    name: local
    key: zfs.pool_name
    value: \"\"
    description: '\"zfs.pool_name\" property for storage pool \"local\"'
  cluster_address: server0:8443
  server_address: server$i:8443
  cluster_password: \"\"
  cluster_certificate_path: \"\"
  cluster_token: \"`get_token $i`\"
EOF"
done

# Launch $nvms VMs on LXD cluster
for i in `seq $nvms`; do sudo lxc launch images:ubuntu/20.04/cloud vm$i --vm -c limits.memory=4GB -c limits.cpu=`getconf _NPROCESSORS_ONLN` "${LXC_OPTS[@]}"; done

# Wait until all VMs are up
echo "Waiting 120 seconds for VMs to start..."
sleep 120

# Install docker on all VMs
for i in `seq $nvms`; do
    sudo lxc exec vm$i -n -- sh -c "
apt-get update
apt-get install --yes curl net-tools iperf3
ifconfig enp5s0 mtu 1450
curl -fsSL https://get.docker.com -o get-docker.sh
sh get-docker.sh" &
done
wait

# Install DeathStarBench
sudo lxc file push hotelreservation.yml vm1/root/
sudo lxc exec vm1 -- sh -c "sudo docker swarm init"

JOIN_COMMAND=$(sudo lxc exec vm1 -- sh -c "sudo docker swarm join-token worker | awk '/docker/ {print $1}'")

# Start Docker on all servers
for i in `seq 2 $nvms`; do
    sudo lxc exec vm$i -n -- sh -c  "$JOIN_COMMAND"
done

