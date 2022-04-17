package microservices

import (
	"context"
	"github.com/mark-by/little-busy-back/api/internal/domain/entity"
	"github.com/mark-by/little-busy-back/api/internal/domain/repository"
	"github.com/mark-by/little-busy-back/auth/pkg/proto/authorization"
)

type Authorization struct {
	authClient authorization.AuthorizationServiceClient
}

func NewAuthorization(authClient authorization.AuthorizationServiceClient) *Authorization {
	return &Authorization{authClient: authClient}
}

func (a Authorization) Create(userID int) (*entity.AuthSession, error) {
	session, err := a.authClient.Create(context.Background(), &authorization.UserID{ID: int64(userID)})
	if err != nil {
		return nil, err
	}

	return &entity.AuthSession{
		ID:             session.ID,
		UserID:         userID,
		ExpirationDate: session.ExpirationDate,
	}, nil
}

func (a Authorization) Check(session string) (int, error) {
	userID, err := a.authClient.Check(context.Background(), &authorization.SessionID{ID: session})
	if err != nil {
		return 0, err
	}

	return int(userID.ID), nil
}

var _ repository.Authorization = &Authorization{}
