package main

import (
	"github.com/mark-by/little-busy-back/auth/internal/application"
	"github.com/mark-by/little-busy-back/auth/internal/infrastructure/redis"
	"github.com/mark-by/little-busy-back/auth/internal/interfaces/grpc"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()

	sessionRepo := redis.NewSession(&redis.Options{
		Host: "127.0.0.1",
		Port: "6379",
		User: "auth",
	})

	sessionApp := application.NewSession(sessionRepo)

	authService := grpc.NewAuthService(sessionApp, logger.Sugar().With("where", "auth_service"))

	logger.Info("Started auth server")
	grpc.NewAuthServer(authService).Start(&grpc.Options{
		Host: "0.0.0.0",
		Port: "8000",
	})
}
