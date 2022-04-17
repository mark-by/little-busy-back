package application

import (
	"github.com/mark-by/little-busy-back/api/internal/domain/entity"
	"github.com/mark-by/little-busy-back/api/internal/domain/repository"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type UserI interface {
	GetUserBySession(session string) (*entity.User, error)
	LogIn(username, password string) (*entity.User, *entity.AuthSession, error)
	SignUp(username, password string) (*entity.User, *entity.AuthSession, error)
	CheckCredentials(username, password string) (*entity.User, error)
}

type User struct {
	repoUser repository.User
	repoAuth repository.Authorization
}

func NewUser(user repository.User, authorization repository.Authorization) *User {
	return &User{
		repoUser: user,
		repoAuth: authorization,
	}
}

func (u User) CheckCredentials(username, password string) (*entity.User, error) {
	user, err := u.repoUser.GetByUsername(username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("no user")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}

	user.Password = ""

	return user, nil
}

func (u User) GetUserBySession(session string) (*entity.User, error) {
	userID, err := u.repoAuth.Check(session)
	if err != nil {
		return nil, err
	}

	user, err := u.repoUser.GetByID(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u User) LogIn(username, password string) (*entity.User, *entity.AuthSession, error) {
	user, err := u.CheckCredentials(username, password)
	if err != nil {
		return nil, nil, err
	}

	session, err := u.repoAuth.Create(user.ID)
	if err != nil {
		return nil, nil, err
	}

	return user, session, nil
}

func (u User) SignUp(username, password string) (*entity.User, *entity.AuthSession, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return nil, nil, err
	}

	user, err := u.repoUser.Create(&entity.User{
		Username: username,
		Password: string(hashedPassword),
	})
	if err != nil {
		return nil, nil, err
	}

	session, err := u.repoAuth.Create(user.ID)
	if err != nil {
		return nil, nil, err
	}

	return user, session, err
}

var _ UserI = &User{}
