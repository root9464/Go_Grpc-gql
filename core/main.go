package main

import (
	"context"
	"root/shared/middleware"
	"root/shared/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/sirupsen/logrus"
	"github.com/ysugimoto/grpc-graphql-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	auth "root/shared/proto/out"
)

func main() {
	// other
	const port = ":6069"
	log := utils.Logger()
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	app.Use(middleware.LoggerHTTP(log))

	// Register GraphQL handler
	mux := runtime.NewServeMux()
	if err := auth.RegisterAuthServiceGraphql(mux); err != nil {
		log.WithError(err).Fatal("❌ Failed to register GraphQL handler")
		return
	}

	// Connect to gRPC server
	client, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.WithError(err).Error("не удалось подключиться к gRPC серверу")
	}
	defer client.Close()

	app.All("/graphql", adaptor.HTTPHandler(mux))
	app.Post("/hello", handleHelloRequest(log, client))

	log.Infof("🌐 Server is running on %s", port)
	if err := app.Listen(port); err != nil {
		log.Fatal("❌ Server startup error:", err)
	}

	log.Info("✅ Server started successfully")
}

func handleHelloRequest(log *logrus.Logger, client *grpc.ClientConn) fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := new(auth.HelloRequest)
		if err := c.BodyParser(&req); err != nil {
			log.WithError(err).Error("не удалось распарсить запрос")
			return c.Status(fiber.StatusBadRequest).SendString("Invalid request")
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		authClient := auth.NewAuthServiceClient(client)
		r, err := authClient.SayHello(ctx, &auth.HelloRequest{Name: req.GetName()})
		if err != nil {
			log.WithError(err).Error("не удалось обработать запрос")
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to process request")
		}

		return c.SendString(r.GetMessage())
	}
}
