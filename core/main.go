package main

import (
	"root/shared/middleware"
	"root/shared/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/ysugimoto/grpc-graphql-gateway/runtime"

	auth "root/shared/proto/out"
)

func main() {
	const port = ":6069"
	log := utils.Logger()
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})
	app.Use(middleware.Logger(log))

	mux := runtime.NewServeMux()
	if err := auth.RegisterAuthServiceGraphql(mux); err != nil {
		log.WithError(err).Fatal("❌ Failed to register GraphQL handler")
		return
	}
	app.All("/graphql", adaptor.HTTPHandler(mux))
	log.Infof("🌐 Server is running on %s", port)
	if err := app.Listen(port); err != nil {
		log.Fatal("❌ Server startup error:", err)
	}

	log.Info("✅ Server started successfully")
}
