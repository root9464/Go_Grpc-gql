package main

import (
	"net"
	rpc "root/services/auth/rpc"
	"root/shared/database"
	auth "root/shared/proto/out" // Correctly import the auth service
	"root/shared/utils"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

// функция main запускает сервер gRPC
// функция register регистрирует все сервисы gRPC
// функция newServices создает новый экземпляр Services и возвращает указатель на него

type Services struct {
	AuthService auth.AuthServiceServer
}

func NewServices(db *gorm.DB) *Services {
	return &Services{
		AuthService: rpc.NewAuthService(db),
	}
}

func (s *Services) Register(server *grpc.Server) {
	auth.RegisterAuthServiceServer(server, s.AuthService)
}

func main() {
	log := utils.Logger()
	database, err := database.ConnectDb()
	if err != nil {
		log.WithError(err).Fatal("❌ Failed to connect to database")
	}

	log.Info("✅ Database connected successfully")

	conn, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.WithError(err).Fatal("❌ Failed to listen on port")
	}
	defer conn.Close()

	server := grpc.NewServer()
	services := NewServices(database.Db)

	// Register all services
	services.Register(server)

	log.Info("✅ Server started successfully")

	if err := server.Serve(conn); err != nil {
		log.WithError(err).Fatal("❌ Failed to start server")
	}
}
