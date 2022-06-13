package microservices

import (
	"github.com/mark-by/little-busy-back/auth/pkg/proto/authorization"
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
