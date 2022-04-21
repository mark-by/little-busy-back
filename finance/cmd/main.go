package main

import (
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/mark-by/little-busy-back/api/internal/infrastructure/postgresql"
	"github.com/mark-by/little-busy-back/finance/internal/application"
	"github.com/mark-by/little-busy-back/finance/internal/interfaces/grpc"
	"github.com/mark-by/little-busy-back/pkg/utils"
	"go.uber.org/zap"
	"time"
)

func initDB(logger *zap.SugaredLogger) clickhouse.Conn {
	options := &utils.Options{
		MigrationsDir: "finance/internal/infrastructure/clickhouse/migrations",
		User:          "finance",
		Password:      "159357",
		Host:          "127.0.0.1",
		Port:          "9000",
		Type:          "clickhouse",
		Name:          "finance",
		Logger:        logger,
	}

	utils.Migrate(options)

	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{"127.0.0.1:9000"},
		Auth: clickhouse.Auth{
			Database: "finance",
			Username: "finance",
			Password: "159357",
		},
		//Debug:           true,
		DialTimeout:     time.Second,
		MaxOpenConns:    10,
		MaxIdleConns:    5,
		ConnMaxLifetime: time.Hour,
		Compression: &clickhouse.Compression{
			Method: clickhouse.CompressionLZ4,
		},
	})

	if err != nil {
		logger.Fatal("fail to connect database: %s", err)
	}

	return conn
}

func main() {
	logger, _ := zap.NewDevelopment()

	conn := initDB(logger.Sugar())

	recordsRepository := postgresql.NewRecord(conn)
	valuesRepository := postgresql.NewValue(conn)

	recordsApp := application.NewRecord(recordsRepository)
	valuesApp := application.NewValue(valuesRepository)

	logger.Info("Start finance server")
	grpc.NewFinanceServer(recordsApp, valuesApp).Start("0.0.0.0:8003")
}
