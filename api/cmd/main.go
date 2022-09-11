package main

import (
	"context"
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
	"time"
)

func initDB(logger *zap.SugaredLogger, sysConfig *config.Config) *pgxpool.Pool {
	config := pgconn.Config{
		Host: "postgres",
		Port: 5432,
		Database: "postgres",
		User:     "postgres",
		Password: "123",
	}

	migrateFunc := func () {
		utils.Migrate(&utils.Options{
			MigrationsDir: sysConfig.Migrations,
			User:          config.User,
			Password:      config.Password,
			Host:          config.Host,
			Port:          strconv.Itoa(int(config.Port)),
			Type:          "postgresql",
			Name:          config.Database,
			Logger:        logger,
		})
	}

	utils.Retry(migrateFunc, 4, 1 * time.Second)

	pool, err := pgxpool.Connect(context.Background(),
		fmt.Sprintf("postgres://%s:%s@%s:%v/%s",
			config.User, config.Password, config.Host, config.Port, config.Database))

	if err != nil {
		logger.Fatal("fail to connect database: %s", err)
	}

	return pool
}

func initConfig() *config.Config {
	return config.ParseConfig()
}

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal("fail to create logger: ", err)
	}

	config := initConfig()

	authClient := microservices.NewAuthorizationClient("auth:8000")
	connDB := initDB(logger.Sugar(), config)

	authorizationRepository := microservices.NewAuthorization(authClient)
	userRepository := postgresql.NewUser(connDB, logger.Sugar().With("user repo"))
	customerRepository := postgresql.NewCustomers(connDB)
	eventsRepository := postgresql.NewEvent(connDB)
	recordsRepository := postgresql.NewRecord(connDB)
	settingsRepository := postgresql.NewSettings(connDB)

	userApp := application.NewUser(userRepository, authorizationRepository)
	authApp := application.NewAuthorization(authorizationRepository)
	customerApp := application.NewCustomers(customerRepository, eventsRepository)
	eventsApp := application.NewEvents(eventsRepository, customerRepository, recordsRepository)
	settingsApp := application.NewSettings(settingsRepository)
	recordsApp := application.NewRecords(recordsRepository, settingsRepository)

	application.NewScheduler(recordsApp, eventsApp, logger.Sugar()).Start()

	log.Print(rest.NewServer(&rest.ServerOptions{
		UserApp:     userApp,
		AuthApp:     authApp,
		CustomerApp: customerApp,
		EventsApp:   eventsApp,
		RecordsApp:  recordsApp,
		SettingsApp: settingsApp,
		Logger:      logger.Sugar(),
		Config:      config,
	}).Start())
}
