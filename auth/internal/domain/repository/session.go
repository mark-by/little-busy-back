package repository

import "github.com/mark-by/little-busy-back/auth/internal/domain/entity"

type Session interface {
	Get(sessionID string) (*entity.Session, error)
	Delete(sessionID string) error
	Create(session *entity.Session) error
}
