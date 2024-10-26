#!/bin/bash

#dka - docker all 
#rmc - remove containers
#rmi - remove images

dka_remove() {
  docker stop $(docker ps -q) && 
  docker rm $(docker ps -a -q) && 
  docker rmi $(docker images -q)
}

dka_stop() {
  docker stop $(docker ps -q)
}

dka_rmc() {
  docker rm $(docker ps -a -q)
}

dka_rmi() {
  docker rmi $(docker images -q)
}

if [ $# -eq 1 ]; then
    "$1"
fi