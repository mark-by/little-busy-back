package application

import (
	"github.com/mark-by/little-busy-back/api/internal/domain/entity"
	"github.com/mark-by/little-busy-back/api/internal/domain/repository"
	"strings"
)

type CustomersI interface {
	Create(customer *entity.Customer) (*entity.Customer, error)
	Get(id int64) (*entity.Customer, error)
	Search(searchText, since string, limit int) ([]entity.Customer, error)
	Delete(id int64) error
	Update(customer *entity.Customer) error
}

type Customers struct {
	customers repository.Customers
	events    repository.Events
}

func (c Customers) Create(customer *entity.Customer) (*entity.Customer, error) {
	return c.customers.CreateCustomer(customer)
}

func (c Customers) Get(id int64) (*entity.Customer, error) {
	return c.customers.GetByID(id)
}

func (c Customers) Search(searchText, since string, limit int) ([]entity.Customer, error) {
	if limit == 0 {
		limit = 10
	}
	if limit > 50 {
		limit = 50
	}

	searchField := "name"
	if strings.HasPrefix(searchText, "+7") || strings.HasPrefix(searchText, "8") {
		searchText = strings.TrimPrefix(searchText, "+7")
		searchText = strings.TrimPrefix(searchText, "8")
		searchField = "tel"
	}
	if strings.HasPrefix(searchText, "9") {
		searchField = "tel"
	}

	return c.customers.SearchCustomers(searchText, searchField, since, limit)
}

func (c Customers) Delete(id int64) error {
	return c.customers.DeleteCustomer(id)
}

func (c Customers) Update(customer *entity.Customer) error {
	return c.customers.UpdateCustomer(customer)
}

func NewCustomers(customers repository.Customers, events repository.Events) *Customers {
	return &Customers{customers, events}
}

var _ CustomersI = &Customers{}
