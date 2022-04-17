package repository

import "github.com/mark-by/little-busy-back/crm/internal/domain/entity"

type User interface {
	Get(username string) (*entity.User, error)
	GetByID(id int) (*entity.User, error)
	Create(user *entity.User) (*entity.User, error)
	IsStorageEmpty() bool
}
