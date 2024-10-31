package auth

import (
	"context"
	"fmt"

	"root/shared/database/models"
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

func (s *Server) SayHello(ctx context.Context, req *out.HelloRequest) (*out.HelloResponse, error) {
	log.Printf("Received SayHello request for name: %s", req.GetName())
	return &out.HelloResponse{
		Message: fmt.Sprintf("Hello epta, %s!", req.GetName()),
	}, nil
}

func (s *Server) CreateUser(ctx context.Context, req *out.CreateUserRequest) (*out.CreateUserResponse, error) {
	log.Printf("Received CreateUser request for name: %s", req.GetName())

	if req.GetName() == "" || req.GetEmail() == "" || req.GetPassword() == "" {
		log.Error("Invalid request")
		return &out.CreateUserResponse{Status: false}, nil
	}

	user := models.User{
		ID:       utils.GenerateRandomID(),
		Username: req.GetName(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}

	if err := s.db.Create(&user).Error; err != nil {
		log.Printf("Failed to create user: %v", err)
		return nil, err
	}
	return &out.CreateUserResponse{Status: true}, nil
}

func (s *Server) GetUserInfo(ctx context.Context, req *out.GetUserInfoRequest) (*out.GetUserInfoResponse, error) {
	log.Printf("Received GetUserInfo request for name: %s", req.GetName())

	if req.GetName() == "" {
		log.Error("Invalid request")
		return &out.GetUserInfoResponse{}, nil
	}

	user := new(models.User)

	if err := s.db.Where("username = ?", req.GetName()).First(&user).Error; err != nil {
		log.Printf("Failed to get user: %v", err)
		return nil, err
	}

	return &out.GetUserInfoResponse{
		Name:     user.Username,
		Password: user.Password,
		Email:    user.Email,
	}, nil
}
