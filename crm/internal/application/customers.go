package application

import (
	"github.com/mark-by/little-busy-back/crm/internal/domain/entity"
	"github.com/mark-by/little-busy-back/crm/internal/domain/repository"
	"strings"
)

type CustomersI interface {
	Create(name, tel string) (*entity.Customer, error)
	Get(customerID int64) (*entity.Customer, error)
	GetCustomers(ids []int64) ([]entity.Customer, error)
	Delete(customerID int64) error
	Search(searchText string, since string, limit int) ([]entity.Customer, error)
	Update(customer *entity.Customer) error
}

type Customers struct {
	repo repository.Customer
}

func (c Customers) GetCustomers(ids []int64) ([]entity.Customer, error) {
	return c.repo.GetCustomers(ids)
}

func NewCustomers(repo repository.Customer) *Customers {
	return &Customers{repo: repo}
}

func (c Customers) Create(name, tel string) (*entity.Customer, error) {
	return c.repo.CreateCustomer(&entity.Customer{
		Name: name,
		Tel:  &tel,
	})
}

func (c Customers) Get(customerID int64) (*entity.Customer, error) {
	return c.repo.GetCustomer(customerID)
}

func (c Customers) Delete(customerID int64) error {
	return c.repo.DeleteCustomer(customerID)
}

func (c Customers) Search(searchText string, since string, limit int) ([]entity.Customer, error) {
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

	return c.repo.SearchCustomers(searchText, searchField, since, limit)
}

func (c Customers) Update(customer *entity.Customer) error {
	return c.repo.Update(customer)
}

var _ CustomersI = &Customers{}
