package application

import "github.com/mark-by/little-busy-back/api/internal/domain/repository"

type AuthorizationI interface {
	CheckAuthorization(session string) (int, error)
}

type Authorization struct {
	repo repository.Authorization
}

func NewAuthorization(authRepo repository.Authorization) *Authorization {
	return &Authorization{repo: authRepo}
}

func (a Authorization) CheckAuthorization(session string) (int, error) {
	return a.repo.Check(session)
}

var _ AuthorizationI = &Authorization{}
