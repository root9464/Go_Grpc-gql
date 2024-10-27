package main

import (
	"net"
	auth "root/services/auth/rpc"
	"root/shared/database"
	"root/shared/proto/out"
	"root/shared/utils"

	"google.golang.org/grpc"
)

func main() {
	log := utils.Logger()
	database, err := database.ConnectDb()
	if err != nil {
		log.WithError(err).Fatal("❌ Failed to connect to database")
	}

	log.Info("✅ Database connected successfully")

	conn, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.WithError(err).Fatal("❌ Failed to listen port")
	}
	defer conn.Close()

	server := grpc.NewServer()
	authService := auth.NewAuthService(database.Db)    // Use the constructor to create the service
	out.RegisterAuthServiceServer(server, authService) // Register the service

	log.Info("✅ Server started successfully")

	if err := server.Serve(conn); err != nil {
		log.WithError(err).Fatal("❌ Failed to start server")
	}
}
