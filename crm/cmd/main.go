package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/mark-by/little-busy-back/crm/internal/application"
	"github.com/mark-by/little-busy-back/crm/internal/infrastructure/postgresql"
	"github.com/mark-by/little-busy-back/crm/internal/interfaces/grpc"
	"github.com/mark-by/little-busy-back/pkg/utils"
	"go.uber.org/zap"
	"strconv"
)

func initDB(logger *zap.SugaredLogger) *pgxpool.Pool {
	config := pgx.ConnConfig{
		Host:     "127.0.0.1",
		Port:     5432,
		Database: "crm",
		User:     "crm",
		Password: "123",
	}
	utils.Migrate(&utils.Options{
		MigrationsDir: "crm/internal/infrastructure/postgresql/migrations",
		User:          config.User,
		Password:      config.Password,
		Host:          config.Host,
		Port:          strconv.Itoa(int(config.Port)),
		Type:          "postgresql",
		Name:          config.Database,
		Logger:        logger,
	})

	pool, err := pgxpool.Connect(context.Background(),
		fmt.Sprintf("postgres://%s:%s@%s:%v/%s",
			config.User, config.Password, config.Host, config.Port, config.Database))

	if err != nil {
		logger.Fatal("fail to connect database: %s", err)
	}

	return pool
}

func main() {
	logger, _ := zap.NewDevelopment()

	conn := initDB(logger.Sugar())
	userRepository := postgresql.NewUser(conn, logger.Sugar().With("where", "user repository"))
	customerRepository := postgresql.NewCustomers(conn)

	userApplication := application.NewUser(userRepository)
	customerApplication := application.NewCustomers(customerRepository)

	crmService := grpc.NewCRMService(userApplication, customerApplication, logger.Sugar())

	logger.Info("Started crm server")
	grpc.NewCRMServer(crmService).Start(&grpc.Options{
		Host: "0.0.0.0",
		Port: "8001",
	})
}
