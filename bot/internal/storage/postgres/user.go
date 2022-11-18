package postgres

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/mark-by/little-busy-back/bot/internal/entity"
	"github.com/pkg/errors"
)

func (s *Storage) SaveUser(user *entity.User) error {
	_, err := s.db.Exec(context.Background(),
		`insert into customer (username, chat_id, tel) values ($1, $2, $3)
			on conflict (chat_id) do update set tel = EXCLUDED.tel`,
		user.Username, user.ChatID, user.Tel)
	if err != nil {
		return errors.Wrap(err, "fail to save user into db")
	}

	return nil
}

func (s *Storage) DeleteUser(chatID int64) error {
	_, err := s.db.Exec(context.Background(), `delete from customer where chat_id = $1`, chatID)
	if err != nil {
		return errors.Wrap(err, "fail to delete user from db")
	}

	return nil
}

func (s *Storage) UserByTel(tel string) (*entity.User, error) {
	user := entity.User{}
	err := pgxscan.Get(context.Background(), s.db, &user, `select * from customer where tel = $1`, tel)
	if err != nil {
		return nil, errors.Wrap(err, "fail to get user from db")
	}

	return &user, nil
}

func (s *Storage) SetNotification(chatID int64, value bool) error {
	_, err := s.db.Exec(context.Background(), `update customer set notification_is_enabled = $1 where chat_id = $2`, value, chatID)
	if err != nil {
		return errors.Wrap(err, "fail to update enable notification in db")
	}

	return nil
}

func (s *Storage) UserByChatID(chatID int64) (*entity.User, error) {
	user := entity.User{}
	err := pgxscan.Get(context.Background(), s.db, &user, `select * from customer where chat_id = $1`, chatID)
	if err != nil {
		return nil, errors.Wrap(err, "fail to get user from db")
	}

	return &user, nil
}
