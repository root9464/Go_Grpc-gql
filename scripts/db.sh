#!/bin/bash
DOCKERCOMPOSE_PATH="../docker-compose.yml"


start_dbpsq() {
  docker-compose -f $DOCKERCOMPOSE_PATH up -d
}

stop_dbpsq() {
  docker-compose -f $DOCKERCOMPOSE_PATH down
}

rm_dbpsq() {
  docker rm postgres &&
  docker rmi postgres
}

if [ $# -eq 1 ]; then
    "$1"
fi