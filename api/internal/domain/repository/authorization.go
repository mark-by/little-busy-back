package repository

import "github.com/mark-by/little-busy-back/api/internal/domain/entity"

type Authorization interface {
	Create(userID int) (*entity.AuthSession, error)
	Check(session string) (int, error)
}
