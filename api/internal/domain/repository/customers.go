package repository

import "github.com/mark-by/little-busy-back/api/internal/domain/entity"

type Customers interface {
	GetByID(id int64) (*entity.Customer, error)
	GetCustomers(ids []int64) ([]entity.Customer, error)
	CreateCustomer(customer *entity.Customer) (*entity.Customer, error)
	DeleteCustomer(id int64) error
	SearchCustomer(searchText, since string, limit int) ([]entity.Customer, error)
	UpdateCustomer(customer *entity.Customer) error
}
