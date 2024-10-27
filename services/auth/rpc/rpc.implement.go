package auth

import (
	"context"
	"fmt"

	"root/shared/proto/out"
	"root/shared/utils"

	"gorm.io/gorm"
)

var log = utils.Logger()

type Server struct {
	out.UnimplementedAuthServiceServer
	db *gorm.DB
}

func NewAuthService(db *gorm.DB) *Server {
	return &Server{db: db}
}

func (s *Server) SayHello(ctx context.Context, req *out.HelloRequest) (*out.HelloReply, error) {
	log.Printf("Received SayHello request for name: %s", req.GetName())
	return &out.HelloReply{
		Message: fmt.Sprintf("Hello, %s!", req.GetName()),
	}, nil
}
