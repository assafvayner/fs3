#!/bin/bash -x

. config.sh

# This script runs only on server0
test x`hostname -s` == xserver0 || { echo "This script has to run on server0"; exit 1; }

ssh -oStrictHostKeyChecking=no client0 "
sudo mkdir -p /mnt/reservation
sudo bash -c 'echo -e \"/mnt/reservation\t*(rw,no_root_squash)\" >> /etc/exports'
sudo /etc/init.d/nfs-kernel-server reload"

for i in `seq 0 $((n_servers - 1))`; do
    ssh -oStrictHostKeyChecking=no server$i "
sudo mkdir -p /mnt/reservation
sudo mount -t nfs client0:/mnt/reservation /mnt/reservation"
done

