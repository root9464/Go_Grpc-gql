package main

import (
	"log"
	"net/http"
	auth "root/shared/proto/out"

	"github.com/ysugimoto/grpc-graphql-gateway/runtime"
)

func main() {
	mux := runtime.NewServeMux()

	if err := auth.RegisterAuthServiceGraphql(mux); err != nil {
		log.Printf("register grpc service to graphql error: %v\n", err)
		return
	}

	log.Println("start graphql server at :8888")
	http.Handle("/graphql", mux)
	log.Fatalln(http.ListenAndServe(":8888", nil))
}
