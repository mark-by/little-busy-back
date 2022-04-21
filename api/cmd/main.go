package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/mark-by/little-busy-back/api/internal/application"
	"github.com/mark-by/little-busy-back/api/internal/config"
	"github.com/mark-by/little-busy-back/api/internal/infrastructure/microservices"
	"github.com/mark-by/little-busy-back/api/internal/infrastructure/postgresql"
	"github.com/mark-by/little-busy-back/api/internal/interfaces/rest"
	"github.com/mark-by/little-busy-back/pkg/utils"
	"go.uber.org/zap"
	"log"
	"strconv"
)

func initDB(logger *zap.SugaredLogger) *pgxpool.Pool {
	config := pgconn.Config{
		Host: "127.0.0.1",
		Port: 5432, Database: "postgres",
		User:     "postgres",
		Password: "123",
	}
	utils.Migrate(&utils.Options{
		MigrationsDir: "api/internal/infrastructure/postgresql/migrations",
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

func initConfig() *config.Config {
	configFilename := flag.String("f", "config.yaml", "config file")
	flag.Parse()

	return config.ParseConfig(*configFilename)
}

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal("fail to create logger: ", err)
	}

	config := initConfig()

	authClient := microservices.NewAuthorizationClient("127.0.0.1:8000")
	connDB := initDB(logger.Sugar())

	authorizationRepository := microservices.NewAuthorization(authClient)
	userRepository := postgresql.NewUser(connDB, logger.Sugar().With("user repo"))
	customerRepository := postgresql.NewCustomers(connDB)
	eventsRepository := postgresql.NewEvent(connDB)

	userApp := application.NewUser(userRepository, authorizationRepository)
	authApp := application.NewAuthorization(authorizationRepository)
	customerApp := application.NewCustomers(customerRepository, eventsRepository)
	eventsApp := application.NewEvents(eventsRepository, customerRepository)

	log.Print(rest.NewServer(&rest.ServerOptions{
		UserApp:     userApp,
		AuthApp:     authApp,
		CustomerApp: customerApp,
		EventsApp:   eventsApp,
		Logger:      logger.Sugar(),
		Config:      config,
	}).Start())
}
