package application

import (
	"github.com/mark-by/little-busy-back/crm/internal/domain/entity"
	"github.com/mark-by/little-busy-back/crm/internal/domain/repository"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type UserI interface {
	GetByID(id int) (*entity.User, error)
	Create(username string, password string) (*entity.User, error)
	CheckCredentials(username string, password string) (*entity.User, error)
}

type User struct {
	repo repository.User
}

func (u User) GetByID(id int) (*entity.User, error) {
	return u.repo.GetByID(id)
}

func (u User) Create(username string, password string) (*entity.User, error) {
	if !u.repo.IsStorageEmpty() {
		return nil, errors.New("User already exists")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}
	return u.repo.Create(&entity.User{
		Username: username,
		Password: string(hashedPassword),
		IsAdmin:  true,
	})
}

func (u User) CheckCredentials(username string, password string) (*entity.User, error) {
	user, err := u.repo.Get(username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		if u.repo.IsStorageEmpty() {
			return u.Create(username, password)
		}
		return nil, errors.New("no user")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}

	user.Password = ""

	return user, nil
}

func NewUser(repo repository.User) *User {
	return &User{
		repo: repo,
	}
}

var _ UserI = &User{}
