#!/bin/bash
#
# This runs as geniuser in /.

# All output and commands go to logfile in /tmp
set -x
exec 1>/tmp/init.log 2>&1

HOSTNAME=`hostname -s`
export MAKEFLAGS=-j

echo "init.sh credentials:"
echo PWD=$PWD
echo USER=$USER
echo GROUP=$GROUP
echo HOSTNAME=$HOSTNAME

# Install a bunch of useful tools
sudo apt-get update
sudo apt-get install --yes libxml-xpath-perl joe nfs-kernel-server

case $HOSTNAME in
    server*)
	# Install virtual machine
	sudo apt-get install --yes snapd
	sudo snap install lxd
	;;

    client*)
	# Build wrk2
	cd /local/repository/DeathStarBench
	sudo apt-get install --yes python3-aiohttp libssl-dev libz-dev luarocks lua-socket
	sudo luarocks install luasocket
	make -C hotelReservation/wrk2
	;;

    *)
	echo "Unknown hostname $HOSTNAME!"
	;;
esac

echo "init.sh done"

