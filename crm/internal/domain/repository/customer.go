package repository

import "github.com/mark-by/little-busy-back/crm/internal/domain/entity"

type Customer interface {
	CreateCustomer(customer *entity.Customer) (*entity.Customer, error)
	DeleteCustomer(customerID int64) error
	GetCustomer(customerID int64) (*entity.Customer, error)
	GetCustomers(ids []int64) ([]entity.Customer, error)
	SearchCustomers(searchText string, searchField string, since string, limit int) ([]entity.Customer, error)
	Update(customer *entity.Customer) error
}
