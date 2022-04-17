package grpc

import (
	"fmt"
	protoScheduler "github.com/mark-by/little-busy-back/scheduler/pkg/proto/scheduler"
	"google.golang.org/grpc"
	"log"
	"net"
)

type SchedulerServer struct {
	schedulerService *SchedulerService
}

type Options struct {
	Host string
	Port string
}

func NewSchedulerServer(service *SchedulerService) *SchedulerServer {
	return &SchedulerServer{schedulerService: service}
}

func (c SchedulerServer) Start(options *Options) {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s",
		options.Host,
		options.Port))
	if err != nil {
		log.Fatalf("fail to listen tcp for host %s:%s: %v", options.Host, options.Port, err)
	}

	server := grpc.NewServer()

	protoScheduler.RegisterSchedulerServer(server, c.schedulerService)

	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
