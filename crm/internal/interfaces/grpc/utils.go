package grpc

import (
	"github.com/mark-by/little-busy-back/crm/internal/domain/entity"
	"github.com/mark-by/little-busy-back/crm/pkg/proto/crm"
)

func convertCustomer(customer *entity.Customer) *crm.Customer {
	return &crm.Customer{
		ID:   customer.ID,
		Name: customer.Name,
		Tel:  dropNullString(customer.Tel),
	}
}

func convertCustomers(customers []entity.Customer) []*crm.Customer {
	var newCustomers []*crm.Customer
	for _, customer := range customers {
		newCustomers = append(newCustomers, convertCustomer(&customer))
	}
	return newCustomers
}

func dropNullString(value *string) string {
	if value == nil {
		return ""
	}
	return *value
}
