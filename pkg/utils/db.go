package utils

import (
	"context"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/clickhouse"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v4"
	"go.uber.org/zap"
	"strconv"
)

type Options struct {
	MigrationsDir string
	User          string
	Password      string
	Host          string
	Port          string
	Name          string
	Type          string
	Logger        *zap.SugaredLogger
}

func Migrate(options *Options) {
	log := options.Logger

	var connectionString string

	switch options.Type {
	case "clickhouse":
		connectionString = fmt.Sprintf("clickhouse://%s:%s@%s:%s/%s?x-multi-statement=true",
			options.User,
			options.Password,
			options.Host,
			options.Port,
			options.Name,
		)
	case "postgresql":
		connectionString = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
			options.User,
			options.Password,
			options.Host,
			options.Port,
			options.Name,
		)
	}

	m, err := migrate.New(
		"file://"+options.MigrationsDir,
		connectionString,
	)
	if err != nil {
		s, _ := fmt.Printf("Fail to connect to database: %s", err)
		panic(s)
	}
	defer m.Close()

	err = m.Up()
	switch err {
	case nil:
		log.Info("Migrate status: migrations applied")
	case migrate.ErrNoChange:
		log.Info("Migrate status: no changes")
	default:
		log.Fatal("Fail to apply migrations: ", err)
	}
}

func EndTx(tx pgx.Tx, err error) {
	if err != nil {
		_ = tx.Rollback(context.Background())
		return
	}
	_ = tx.Commit(context.Background())
}

func SQLSlice[T any](slice []T) string {
	var paramrefs string
	for i, _ := range slice {
		paramrefs += `$` + strconv.Itoa(i+1) + `,`
	}
	paramrefs = paramrefs[:len(paramrefs)-1] // remove last ","
	return paramrefs
}
