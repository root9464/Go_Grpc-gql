package rpc

import (
	"context"
	"fmt"
	auth "root/shared/proto/out"
	"root/shared/utils"

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
