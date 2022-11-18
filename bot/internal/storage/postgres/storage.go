package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/mark-by/little-busy-back/pkg/utils"
	"log"
	"time"
)

type Storage struct {
	db *pgxpool.Pool
}

func (s *Storage) initDB(config *utils.Options) {
	migrateFunc := func() {
		utils.Migrate(config)
	}

	utils.Retry(migrateFunc, 4, 1*time.Second)

	pool, err := pgxpool.Connect(context.Background(),
		fmt.Sprintf("postgres://%s:%s@%s:%v/%s",
			config.User, config.Password, config.Host, config.Port, config.Database))

	if err != nil {
		log.Fatal("fail to connect database: ", err)
	}

	s.db = pool
}

func NewStorage(config *utils.Options) *Storage {
	st := &Storage{}
	st.initDB(config)
	return st
}
