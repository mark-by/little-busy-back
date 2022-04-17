package postgresql

import (
	"context"
	"errors"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/mark-by/little-busy-back/api/internal/domain/entity"
	"github.com/mark-by/little-busy-back/api/internal/domain/repository"
	"go.uber.org/zap"
	"log"
)

type User struct {
	db     *pgxpool.Pool
	logger *zap.SugaredLogger
}

func NewUser(db *pgxpool.Pool, logger *zap.SugaredLogger) *User {
	return &User{
		db:     db,
		logger: logger,
	}
}

func (u User) GetByID(id int) (*entity.User, error) {

	user := entity.User{
		ID: id,
	}
	err := pgxscan.Get(context.Background(), u.db, &user,
		"select username from users where id = $1", id)

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u User) GetByUsername(username string) (*entity.User, error) {
	var user entity.User
	err := pgxscan.Get(context.Background(), u.db, &user,
		"select id, username, password from users where username = $1", username)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u User) Create(user *entity.User) (*entity.User, error) {
	var id int
	err := u.db.QueryRow(context.Background(), "insert into users (username, password) values ($1, $2) returning id", user.Username, user.Password).Scan(&id)
	if err != nil {
		return nil, err
	}

	user.ID = id

	return user, nil
}

func (u User) IsStorageEmpty() bool {

	var exists bool
	err := u.db.QueryRow(context.Background(), "select exists(select from users)").Scan(&exists)
	if err != nil {
		log.Print("storage: ", err)
		return false
	}

	return !exists
}

var _ repository.User = &User{}
