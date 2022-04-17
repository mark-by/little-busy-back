package microservices

import (
	"context"
	"github.com/mark-by/little-busy-back/api/internal/domain/entity"
	"github.com/mark-by/little-busy-back/api/internal/domain/repository"
	"github.com/mark-by/little-busy-back/crm/pkg/proto/crm"
)

type User struct {
	crmClient crm.CrmServiceClient
}

func NewUser(crmClient crm.CrmServiceClient) *User {
	return &User{crmClient: crmClient}
}

func (u User) GetByID(id int) (*entity.User, error) {
	user, err := u.crmClient.GetUserByID(context.Background(), &crm.User{
		ID: int64(id),
	})
	if err != nil {
		return nil, err
	}

	return &entity.User{
		ID:       int(user.ID),
		Username: user.Username,
	}, nil
}

func (u User) Create(user *entity.User) (*entity.User, error) {
	createdUser, err := u.crmClient.CreateUser(context.Background(), &crm.Credentials{
		Username: user.Username,
		Password: user.Password,
	})
	if err != nil {
		return nil, err
	}

	user.ID = int(createdUser.ID)
	user.Password = ""

	return user, nil
}

func (u User) CheckCredentials(username, password string) (*entity.User, error) {
	user, err := u.crmClient.CheckCredentials(context.Background(), &crm.Credentials{
		Username: username,
		Password: password,
	})
	if err != nil {
		return nil, err
	}

	return &entity.User{
		ID:       int(user.ID),
		Username: user.Username,
	}, nil
}

var _ repository.User = &User{}
