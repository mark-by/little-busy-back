package repository

import "github.com/mark-by/little-busy-back/api/internal/domain/entity"

type User interface {
	GetByID(id int) (*entity.User, error)
	GetByUsername(username string) (*entity.User, error)
	Create(*entity.User) (*entity.User, error)
}
