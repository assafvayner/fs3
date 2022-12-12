#!/bin/bash -x

sudo docker service create --name registry --publish published=4000,target=5000 registry:2
sudo docker compose -f $1 build
sudo docker compose -f $1 push
#sudo docker stack deploy --compose-file $1 hotelreservation

