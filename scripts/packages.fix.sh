#!/bin/bash

# в моменте у меня возникала проблема в том что не было пакета protoc-gen-go или protoc-gen-graphql
# я очень долго потратил время на решение данной проблемы, и доконца не понял в чем были пиколы но 
# данный скрипт создан мною для решения этой проблемы

# ./packages.fix.sh pwhich protoc-gen-go     
# ${HOME}/go/bin/protoc-gen-go
pwhich() {
  [ $# -eq 1 ] && which "$1" || echo "Usage: pwhich <command>"
}

ppath() {
  export PATH="$PATH:$(go env GOPATH)/bin"
  echo "Updated PATH: $PATH"
}

pingo() {
  echo "pong"
}

pupdate() {
  export GOROOT=/usr/local/go &&
  export GOPATH=$HOME/go &&
  export GOBIN=$GOPATH/bin &&
  export PATH=$PATH:$GOROOT:$GOPATH:$GOBIN
}

pinstall() {
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest &&
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest &&
  go get github.com/ysugimoto/grpc-graphql-gateway/protoc-gen-graphql &&
  go get github.com/graphql-go/graphql &&
  go get github.com/ysugimoto/grpc-graphql-gateway/graphql &&
  go get github.com/ysugimoto/grpc-graphql-gateway
}

# Main script logic
case "$1" in
    ppath)
        ppath
        ;;
    pingo)
        pingo
        ;;
    pinstall)
        pinstall
        ;;
    pwhich)
        [ $# -eq 2 ] && pwhich "$2" || echo "Usage: pwhich <command>"
        ;;
    *)
        echo "Invalid command: $1"
        echo "Usage: $0 {pwhich|ppath|pingo} [args]"
        exit 1
        ;;
esac
