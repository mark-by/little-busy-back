package application

import (
	"github.com/mark-by/little-busy-back/auth/internal/domain/entity"
	"github.com/mark-by/little-busy-back/auth/internal/domain/repository"
	"github.com/pkg/errors"
	"time"
)

type SessionI interface {
	Get(sessionID string) (*entity.Session, error)
	Delete(sessionID string) error
	Create(userID int) (*entity.Session, error)
}

type Session struct {
	repo repository.Session
}

func NewSession(repoSession repository.Session) *Session {
	return &Session{repoSession}
}

func (s Session) Get(sessionID string) (*entity.Session, error) {
	return s.repo.Get(sessionID)
}

func (s Session) Delete(sessionID string) error {
	return s.repo.Delete(sessionID)
}

func (s Session) Create(userID int) (*entity.Session, error) {
	value := s.createSessionValue(userID)

	newSession := &entity.Session{
		ID:         value,
		UserID:     userID,
		Expiration: time.Now().Add(30 * 24 * time.Hour),
	}

	err := s.repo.Create(newSession)
	if err != nil {
		return nil, errors.Wrap(err, "fail to create authorization")
	}

	return newSession, nil
}

var _ SessionI = &Session{}
