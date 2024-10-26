package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"root/shared/database"
	auth "root/shared/proto/out"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

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
	log.Println("start gRPC server at :50051")
	conn, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	server := grpc.NewServer()
	auth.RegisterAuthServiceServer(server, &Server{db: database.DB.Db}) // Используем mongo.Conn
	server.Serve(conn)

}
