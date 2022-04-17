package microservices

import (
	"github.com/mark-by/little-busy-back/auth/pkg/proto/authorization"
	"github.com/mark-by/little-busy-back/scheduler/pkg/proto/scheduler"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func NewAuthorizationClient(address string) authorization.AuthorizationServiceClient {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("fail to connect auth service on ", address)
	}

	return authorization.NewAuthorizationServiceClient(conn)
}

func NewSchedulerClient(address string) scheduler.SchedulerClient {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("fail to connect crm service on ", address)
	}

	return scheduler.NewSchedulerClient(conn)
}
