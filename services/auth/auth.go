package main

import (
	"net"
	"root/shared/database"
	auth "root/shared/proto/out"
	"root/shared/utils"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type Server struct {
	auth.UnimplementedAuthServiceServer
	db *gorm.DB
}

func main() {
	log := utils.Logger()

	conn, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	db, err := database.ConnectDb()
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	config := &Server{
		UnimplementedAuthServiceServer: auth.UnimplementedAuthServiceServer{},
		db:                             db.Db,
	}

	auth.RegisterAuthServiceServer(server, config)
	err = server.Serve(conn)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Starting server on :50051")
	log.Info("Database connected")
	log.Info("Server started")
}
