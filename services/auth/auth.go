package main

import (
	"context"
	"fmt"
	"net"
	"root/shared/database"
	auth "root/shared/proto/out"
	"root/shared/utils"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

var log = utils.Logger()

type Server struct {
	auth.UnimplementedAuthServiceServer
	db *gorm.DB
}

func (s *Server) SayHello(ctx context.Context, req *auth.HelloRequest) (*auth.HelloReply, error) {
	log.Printf("Received SayHello request for name: %s", req.GetName())
	return &auth.HelloReply{
		Message: fmt.Sprintf("Hello, %s!", req.GetName()),
	}, nil
}

func main() {
	log := utils.Logger()
	database, err := database.ConnectDb()
	if err != nil {
		log.WithError(err).Fatal("Failed to connect database")
	}

	log.Info("Database connected successfully")

	conn, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.WithError(err).Fatal("Failed to listen")
	}
	defer conn.Close()

	server := grpc.NewServer()
	auth.RegisterAuthServiceServer(server, &Server{db: database.Db}) // Используем mongo.Conn
	log.Info("Server started successfully")

	if err := server.Serve(conn); err != nil {
		log.WithError(err).Fatal("Failed to serve")
	}

}
