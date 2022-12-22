#!/bin/bash -x

. config.sh

# This script runs only on server0
test x`hostname -s` == xserver0 || { echo "This script has to run on server0"; exit 1; }

# Start Docker on all servers
for i in `seq 0 $((n_servers - 1))`; do
    ssh -oStrictHostKeyChecking=no server$i "curl -fsSL https://get.docker.com -o get-docker.sh && sudo sh get-docker.sh" &
done
wait

IP=`ip ad show | grep -s '10.10.' | awk '{ print $2 }'` # Local network IP
sudo docker swarm init --advertise-addr ${IP%/*}
# echo "Add other servers via the join command presented by docker"

# Add other servers as workers
JOIN_COMMAND=$(sudo docker swarm join-token worker | awk '/docker/ {print $0}')
# Start Docker on all servers
for i in `seq 1 $((n_servers - 1))`; do
    ssh -oStrictHostKeyChecking=no server$i "sudo $JOIN_COMMAND"
done

