package repository

import "github.com/mark-by/little-busy-back/src/domain/entity"

type Customer interface {
	Get(ID int64) (*entity.Customer, error)
	GetAll(since int64) ([]entity.Customer, error)
	Find(text string) ([]entity.Customer, error)
	Save(customer *entity.Customer) error
	Delete(ID int64) error
}
