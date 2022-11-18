package crmClient

import (
	"context"
	"github.com/mark-by/little-busy-back/api/pkg/proto/crm"
	"github.com/mark-by/little-busy-back/bot/internal/entity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

type GrpcCrmClient struct {
	grpc crm.CrmServiceClient
}

func newGrpcCrmServiceClient(address string) crm.CrmServiceClient {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("fail to connect auth service on ", address)
	}

	return crm.NewCrmServiceClient(conn)
}

func NewGrpcCrmClient(address string) *GrpcCrmClient {
	return &GrpcCrmClient{grpc: newGrpcCrmServiceClient(address)}
}

func (g GrpcCrmClient) convertEvents(events []*crm.Event) []entity.Event {
	var convertedEvents []entity.Event
	loc, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		log.Println("err: fail to load location: ", err)
		return nil
	}

	for _, event := range events {
		convertedEvents = append(convertedEvents, entity.Event{
			ClientTel: event.ClientTel,
			StartTime: time.Unix(event.StartTime, 0).In(loc),
			EndTime:   time.Unix(event.EndTime, 0).In(loc),
			Price:     event.Price,
		})
	}

	return convertedEvents
}

func (g GrpcCrmClient) GetTomorrowEvents() ([]entity.Event, error) {
	events, err := g.grpc.GetTomorrowEvents(context.Background(), &crm.Empty{})
	if err != nil {
		return nil, err
	}

	return g.convertEvents(events.Result), nil
}

func (g GrpcCrmClient) GetFutureEventsForCustomer(customerTel string) ([]entity.Event, error) {
	events, err := g.grpc.GetFutureEventsForUser(context.Background(), &crm.User{Tel: customerTel})
	if err != nil {
		return nil, err
	}

	return g.convertEvents(events.Result), nil
}

func (g GrpcCrmClient) GetUser(tel string) (*entity.User, error) {
	user, err := g.grpc.GetUserByPhoneNumber(context.Background(), &crm.User{Tel: tel})
	if err != nil {
		return nil, err
	}

	return &entity.User{
		Tel: user.Tel,
	}, nil
}

var _ CrmClient = &GrpcCrmClient{}
