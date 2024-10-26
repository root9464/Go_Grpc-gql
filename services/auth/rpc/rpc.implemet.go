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
	return &auth.HelloReply{
		Message: fmt.Sprintf("Hello, %s!", req.GetName()),
	}, nil
}

func (s *Server) CreateUser(ctx context.Context, req *auth.CreateUserRequest) *auth.CreateUserResponse {
	if s.db == nil {
		log.Error("Database is not connected")
		return nil
	}

	user := &auth.CreateUserRequest{
		Name: req.GetName(),
	}
	if err := s.db.Create(user).Error; err != nil {
		log.Error(err)
		return nil
	}

	log.Info(user)

	s.db.Create(user)

	return &auth.CreateUserResponse{Status: true}
}
