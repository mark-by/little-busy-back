package grpc

import (
	"fmt"
	protoAuth "github.com/mark-by/little-busy-back/auth/pkg/proto/authorization"
	"google.golang.org/grpc"
	"log"
	"net"
)

type AuthServer struct {
	authService *AuthService
}

type Options struct {
	Host string
	Port string
}

func NewAuthServer(service *AuthService) *AuthServer {
	return &AuthServer{authService: service}
}

func (a AuthServer) Start(options *Options) {

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s",
		options.Host,
		options.Port))
	if err != nil {
		log.Fatalf("fail to licten tcp for host %s:%s: %v", options.Host, options.Port, err)
	}

	server := grpc.NewServer()

	protoAuth.RegisterAuthorizationServiceServer(server, a.authService)

	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
