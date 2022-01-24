package repository

import "github.com/mark-by/little-busy-back/src/domain/entity"

type User interface {
	Save(user entity.User) error
}
