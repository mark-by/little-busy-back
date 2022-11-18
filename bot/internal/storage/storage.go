package storage

import "github.com/mark-by/little-busy-back/bot/internal/entity"

type Storage interface {
	UserStorage
}

type UserStorage interface {
	SaveUser(user *entity.User) error
	SetNotification(chatID int64, value bool) error
	DeleteUser(chatID int64) error
	UserByTel(tel string) (*entity.User, error)
	UserByChatID(chatID int64) (*entity.User, error)
}
