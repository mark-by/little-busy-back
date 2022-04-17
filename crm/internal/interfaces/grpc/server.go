package grpc

import (
	"fmt"
	protoCRM "github.com/mark-by/little-busy-back/crm/pkg/proto/crm"
	"google.golang.org/grpc"
	"log"
	"net"
)

type CRMServer struct {
	crmService *CRMService
}

type Options struct {
	Host string
	Port string
}

func NewCRMServer(service *CRMService) *CRMServer {
	return &CRMServer{crmService: service}
}

func (c CRMServer) Start(options *Options) {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s",
		options.Host,
		options.Port))
	if err != nil {
		log.Fatalf("fail to licten tcp for host %s:%s: %v", options.Host, options.Port, err)
	}

	server := grpc.NewServer()

	protoCRM.RegisterCrmServiceServer(server, c.crmService)

	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}