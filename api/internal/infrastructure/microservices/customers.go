package microservices

import (
	"context"
	"github.com/mark-by/little-busy-back/api/internal/domain/entity"
	"github.com/mark-by/little-busy-back/api/internal/domain/repository"
	"github.com/mark-by/little-busy-back/crm/pkg/proto/crm"
)

type Customers struct {
	crmClient crm.CrmServiceClient
}

func (c Customers) GetCustomers(ids []int64) ([]entity.Customer, error) {
	customers, err := c.crmClient.GetCustomers(context.Background(), &crm.CustomersRequest{Ids: ids})
	if err != nil {
		return nil, err
	}
	return convertProtoCustomers(customers.GetCustomers()), nil
}

func (c Customers) GetByID(id int64) (*entity.Customer, error) {
	customer, err := c.crmClient.GetCustomer(context.Background(), &crm.CustomerID{ID: id})
	if err != nil {
		return nil, err
	}

	return &entity.Customer{
		ID:   customer.ID,
		Name: customer.Name,
		Tel:  customer.Tel,
	}, nil
}

func (c Customers) CreateCustomer(customer *entity.Customer) (*entity.Customer, error) {
	createdCustomer, err := c.crmClient.CreateCustomer(context.Background(), &crm.Customer{
		Name: customer.Name,
		Tel:  customer.Tel,
	})
	if err != nil {
		return nil, err
	}

	customer.ID = createdCustomer.ID
	return customer, nil
}

func (c Customers) DeleteCustomer(id int64) error {
	_, err := c.crmClient.DeleteCustomer(context.Background(), &crm.CustomerID{ID: id})
	return err
}

func (c Customers) SearchCustomer(searchText, since string, limit int) ([]entity.Customer, error) {
	foundCustomer, err := c.crmClient.SearchCustomer(context.Background(), &crm.SearchFilter{
		SearchText: searchText,
		Since:      since,
		Limit:      int32(limit),
	})
	if err != nil {
		return nil, err
	}
	customers := make([]entity.Customer, 0, len(foundCustomer.Customers))
	for _, customer := range foundCustomer.Customers {
		customers = append(customers, entity.Customer{
			ID:   customer.ID,
			Name: customer.Name,
			Tel:  customer.Tel,
		})
	}
	return customers, nil
}

func (c Customers) UpdateCustomer(customer *entity.Customer) error {
	_, err := c.crmClient.UpdateCustomer(context.Background(), &crm.Customer{
		ID:   int64(customer.ID),
		Name: customer.Name,
		Tel:  customer.Tel,
	})
	return err
}

func NewCustomers(client crm.CrmServiceClient) *Customers {
	return &Customers{crmClient: client}
}

var _ repository.Customers = &Customers{}
