#!/bin/bash
DOCKERCOMPOSE_PATH="../docker-compose.yml"


start_dbpsq() {
  docker-compose -f $DOCKERCOMPOSE_PATH up -d
}

if [ $# -eq 1 ]; then
    "$1"
fi