package grpc

import (
	"context"
	"fmt"
	"github.com/mark-by/little-busy-back/api/internal/application"
	"github.com/mark-by/little-busy-back/api/internal/domain/entity"
	"github.com/mark-by/little-busy-back/api/pkg/proto/crm"
	"github.com/mark-by/little-busy-back/pkg/utils"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"time"
)

var (
	ErrorNilUser      = errors.New("user is nil")
	ErrorUserNotFound = errors.New("user not found")
)

type CRMService struct {
	customerApp application.CustomersI
	eventApp    application.EventsI
	logger      *zap.SugaredLogger
}

func NewCRMService(customerApp application.CustomersI, eventApp application.EventsI, logger *zap.SugaredLogger) *CRMService {
	return &CRMService{
		customerApp: customerApp,
		eventApp:    eventApp,
		logger:      logger,
	}
}

func (c CRMService) getCustomerByPhoneNumber(tel string) (*entity.Customer, error) {
	customers, err := c.customerApp.Search(tel, "", 1)
	if err != nil {
		c.logger.With("tel", tel).Error("fail to search user: ", err)
		return nil, errors.Wrap(err, "fail to search user")
	}

	if len(customers) == 0 {
		return nil, nil
	}

	return &customers[0], nil
}

func (c CRMService) GetUserByPhoneNumber(ctx context.Context, user *crm.User) (*crm.User, error) {
	if user == nil {
		return nil, ErrorNilUser
	}

	c.logger.Info(fmt.Sprintf("search tel :%s", user.Tel))

	customer, err := c.getCustomerByPhoneNumber(user.Tel)
	if customer == nil {
		return nil, err
	}

	return &crm.User{Tel: *customer.Tel}, nil
}

func (c CRMService) convertEvents(events entity.Events) *crm.Events {
	protoEvents := &crm.Events{}
	for _, event := range events {
		protoEvents.Result = append(protoEvents.Result, &crm.Event{
			ClientTel: utils.DropNil(event.Customer.Tel),
			StartTime: event.StartTime.Unix(),
			EndTime:   event.EndTime.Unix(),
			Price:     utils.DropNil(event.Price),
		})
	}

	return protoEvents
}

func (c CRMService) GetFutureEventsForUser(ctx context.Context, user *crm.User) (*crm.Events, error) {
	if user == nil {
		return nil, ErrorNilUser
	}

	customer, err := c.getCustomerByPhoneNumber(user.Tel)
	if customer == nil && err == nil {
		return nil, ErrorUserNotFound
	}

	events, err := c.eventApp.GetForCustomer(customer.ID, time.Now(), 30)
	if err != nil {
		c.logger.With("customer_id", customer.ID).Error("fail to get events: ", err)
		return nil, err
	}

	return c.convertEvents(events), nil
}

func (c CRMService) GetLastRecordsForUser(ctx context.Context, user *crm.User) (*crm.Events, error) {
	//TODO implement me
	panic("implement me")
}

func (c CRMService) GetTomorrowEvents(ctx context.Context, empty *crm.Empty) (*crm.Events, error) {
	events, err := c.eventApp.GetForDay(time.Now().Year(), int(time.Now().Month()), time.Now().AddDate(0, 0, 1).Day())

	if err != nil {
		c.logger.With("date", time.Now()).Error("fail to get events for day")
		return nil, err
	}

	return c.convertEvents(events), nil
}

var _ crm.CrmServiceServer = &CRMService{}
