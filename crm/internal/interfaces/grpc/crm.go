package grpc

import (
	context "context"
	"github.com/mark-by/little-busy-back/crm/internal/application"
	"github.com/mark-by/little-busy-back/crm/internal/domain/entity"
	protoCRM "github.com/mark-by/little-busy-back/crm/pkg/proto/crm"
	"go.uber.org/zap"
)

type CRMService struct {
	userApp     application.UserI
	customerApp application.CustomersI
	logger      *zap.SugaredLogger
}

func (c CRMService) GetCustomers(ctx context.Context, request *protoCRM.CustomersRequest) (*protoCRM.Customers, error) {
	customers, err := c.customerApp.GetCustomers(request.GetIds())
	if err != nil {
		return nil, err
	}

	return &protoCRM.Customers{Customers: convertCustomers(customers)}, nil
}

func (c CRMService) UpdateCustomer(ctx context.Context, customer *protoCRM.Customer) (*protoCRM.Empty, error) {
	return new(protoCRM.Empty), c.customerApp.Update(&entity.Customer{
		ID:   customer.ID,
		Name: customer.Name,
		Tel:  &customer.Tel,
	})
}

func (c CRMService) GetCustomer(ctx context.Context, id *protoCRM.CustomerID) (*protoCRM.Customer, error) {
	customer, err := c.customerApp.Get(id.GetID())
	if err != nil {
		return nil, err
	}
	return convertCustomer(customer), nil
}

func (c CRMService) GetUserByID(ctx context.Context, user *protoCRM.User) (*protoCRM.User, error) {
	foundUser, err := c.userApp.GetByID(int(user.ID))
	if err != nil {
		return nil, err
	}

	return &protoCRM.User{
		ID:       int64(foundUser.ID),
		Username: foundUser.Username,
	}, nil
}

func (c CRMService) CreateUser(ctx context.Context, credentials *protoCRM.Credentials) (*protoCRM.User, error) {
	user, err := c.userApp.Create(credentials.Username, credentials.Password)
	if err != nil {
		return nil, err
	}
	return &protoCRM.User{
		ID:       int64(user.ID),
		Username: user.Username,
	}, nil
}

func (c CRMService) CheckCredentials(ctx context.Context, credentials *protoCRM.Credentials) (*protoCRM.User, error) {
	user, err := c.userApp.CheckCredentials(credentials.Username, credentials.Password)
	if err != nil {
		return nil, err
	}
	return &protoCRM.User{
		ID:       int64(user.ID),
		Username: user.Username,
	}, nil
}

func (c CRMService) CreateCustomer(ctx context.Context, customer *protoCRM.Customer) (*protoCRM.Customer, error) {
	newCustomer, err := c.customerApp.Create(customer.Name, customer.Tel)

	if err != nil {
		return nil, err
	}

	return convertCustomer(newCustomer), nil
}

func (c CRMService) DeleteCustomer(ctx context.Context, id *protoCRM.CustomerID) (*protoCRM.Empty, error) {
	return new(protoCRM.Empty), c.customerApp.Delete(id.ID)
}

func (c CRMService) SearchCustomer(ctx context.Context, filter *protoCRM.SearchFilter) (*protoCRM.Customers, error) {
	customers, err := c.customerApp.Search(filter.SearchText, filter.Since, int(filter.Limit))
	if err != nil {
		return nil, err
	}

	protoCustomers := make([]*protoCRM.Customer, 0, len(customers))
	for _, customer := range customers {
		protoCustomers = append(protoCustomers, convertCustomer(&customer))
	}

	return &protoCRM.Customers{Customers: protoCustomers}, nil
}

func NewCRMService(userApp application.UserI, customerApp application.CustomersI, logger *zap.SugaredLogger) *CRMService {
	wrappedLogger := logger.With(zap.String("grpc_service", "crm"))
	return &CRMService{
		userApp:     userApp,
		customerApp: customerApp,
		logger:      wrappedLogger,
	}
}

var _ protoCRM.CrmServiceServer = &CRMService{}
