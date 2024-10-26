package main

import (
	"net/http"
	auth "root/shared/proto/out"
	"root/shared/utils"

	"github.com/ysugimoto/grpc-graphql-gateway/runtime"
)

func main() {
	const port = ":6069"
	log := utils.Logger()
	mux := runtime.NewServeMux()

	if err := auth.RegisterAuthServiceGraphql(mux); err != nil {
		log.WithError(err).Fatal("Failed to register graphql handler")
		return
	}

	log.Infof("Listening on %s", port)
	http.Handle("/graphql", mux)
	log.Fatalln(http.ListenAndServe(port, nil))
}
