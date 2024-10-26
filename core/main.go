package main

import (
	"root/shared/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/ysugimoto/grpc-graphql-gateway/runtime"

	auth "root/shared/proto/out"
)

func main() {
	const port = ":6069"
	log := utils.Logger()
	app := fiber.New()
	app.Use(logger.New())
	mux := runtime.NewServeMux()

	if err := auth.RegisterAuthServiceGraphql(mux); err != nil {
		log.WithError(err).Fatal("Failed to register GraphQL handler")
		return
	}

	app.All("/graphql", adaptor.HTTPHandler(mux))

	log.Infof("Listening on %s", port)
	if err := app.Listen(port); err != nil {
		log.Fatal(err)
	}
}
