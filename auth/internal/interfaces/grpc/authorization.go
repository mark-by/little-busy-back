package grpc

import (
	"context"
	"github.com/mark-by/little-busy-back/auth/internal/application"
	protoAuth "github.com/mark-by/little-busy-back/auth/pkg/proto/authorization"
	"go.uber.org/zap"
)

type AuthService struct {
	sessionApp application.SessionI
	logger     *zap.SugaredLogger
}

func NewAuthService(sessionApp application.SessionI, logger *zap.SugaredLogger) *AuthService {
	wrappedLogger := logger.With(zap.String("grpc_service", "auth"))
	return &AuthService{
		sessionApp: sessionApp,
		logger:     wrappedLogger,
	}
}

func (a AuthService) Create(ctx context.Context, id *protoAuth.UserID) (*protoAuth.Session, error) {
	session, err := a.sessionApp.Create(int(id.GetID()))
	if err != nil {
		return nil, err
	}
	return &protoAuth.Session{
		ID:             session.ID,
		ExpirationDate: session.Expiration.Unix(),
	}, nil
}

func (a AuthService) Check(ctx context.Context, id *protoAuth.SessionID) (*protoAuth.UserID, error) {
	session, err := a.sessionApp.Get(id.GetID())
	if err != nil {
		return nil, err
	}
	if session == nil {
		return nil, nil
	}

	return &protoAuth.UserID{
		ID: int64(session.UserID),
	}, nil
}

func (a AuthService) Delete(ctx context.Context, id *protoAuth.SessionID) (*protoAuth.Empty, error) {
	err := a.sessionApp.Delete(id.GetID())
	if err != nil {
		return nil, err
	}

	return nil, err
}

var _ protoAuth.AuthorizationServiceServer = &AuthService{}
