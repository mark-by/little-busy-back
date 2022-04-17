package postgresql

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/mark-by/little-busy-back/crm/internal/domain/entity"
	"github.com/mark-by/little-busy-back/crm/internal/domain/repository"
	"github.com/mark-by/little-busy-back/pkg/utils"
	"github.com/pkg/errors"
	"strings"
)

type Customers struct {
	db *pgxpool.Pool
}

func (c Customers) GetCustomers(ids []int64) ([]entity.Customer, error) {
	if len(ids) == 0 {
		return nil, errors.New("empty list for get customers")
	}
	var customers []entity.Customer
	args := make([]interface{}, 0, len(ids))
	for _, arg := range ids {
		args = append(args, arg)
	}
	err := pgxscan.Select(context.Background(), c.db, &customers,
		"select * from customers where id in ("+utils.SQLSlice(ids)+")", args...)
	if err != nil {
		return nil, errors.Wrap(err, "fail to select customers")
	}

	return customers, nil
}

func (c Customers) Update(customer *entity.Customer) error {
	_, err := c.db.Exec(context.Background(),
		"update customers set name = $1, tel = $2 where id = $3", customer.Name, customer.Tel, customer.ID)
	if err != nil {
		return errors.Wrap(convertPgxError(err), "fail to update")
	}
	return nil
}

func (c Customers) GetCustomer(customerID int64) (*entity.Customer, error) {
	customer := entity.Customer{
		ID: customerID,
	}
	err := pgxscan.Get(context.Background(), c.db, &customer,
		"select name, tel from customers where id = $1", customerID)
	if err != nil {
		return nil, errors.Wrap(err, "fail to ger customer by id")
	}
	return &customer, nil
}

func (c Customers) CreateCustomer(customer *entity.Customer) (*entity.Customer, error) {
	err := c.db.QueryRow(context.Background(),
		"insert into customers (name, tel) values ($1, $2) returning id",
		customer.Name, customer.Tel).Scan(&customer.ID)
	if err != nil {
		return nil, errors.Wrap(convertPgxError(err), "create customer fail")
	}

	return customer, nil
}

func (c Customers) DeleteCustomer(customerID int64) error {
	_, err := c.db.Exec(context.Background(), "delete from customers where id = $1", customerID)
	if err != nil {
		return errors.Wrap(err, "delete customer fail: ")
	}
	return nil
}

func (c Customers) SearchCustomers(searchText string, searchField string, since string, limit int) ([]entity.Customer, error) {
	var args []interface{}

	sqlFilterByName := "name > $1 "
	args = append(args, since)

	sqlFilterBySearchField := ""
	if searchText != "" {
		words := strings.Split(searchText, " ")
		convertedWords := make([]string, 0, len(words))
		for _, word := range words {
			if word == "" {
				continue
			}
			convertedWords = append(convertedWords, fmt.Sprintf("%s:*|%s:*",
				strings.ToLower(word),
				strings.Title(word)))
		}
		searchText = strings.Join(convertedWords, "|")
		args = append(args, searchText)
		sqlFilterBySearchField = fmt.Sprintf("and to_tsvector(%s) @@ to_tsquery('russian',$2) ", searchField)
	}

	args = append(args, limit)
	sqlOrder := fmt.Sprintf(" order by name, id limit $%d", len(args))

	var customers []entity.Customer

	err := pgxscan.Select(context.Background(), c.db, &customers, "select id, name, tel from customers where "+
		sqlFilterByName+sqlFilterBySearchField+sqlOrder, args...)

	if err != nil {
		return nil, errors.Wrap(err, "select search customers fail: ")
	}

	return customers, nil
}

func NewCustomers(db *pgxpool.Pool) *Customers {
	return &Customers{db}
}

var _ repository.Customer = &Customers{}
