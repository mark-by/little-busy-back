package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/mark-by/little-busy-back/bot/internal/crmClient"
	"github.com/mark-by/little-busy-back/bot/internal/handler"
	"github.com/mark-by/little-busy-back/bot/internal/listener/telegramListener"
	"github.com/mark-by/little-busy-back/bot/internal/scheduler"
	"github.com/mark-by/little-busy-back/bot/internal/storage/postgres"
	"github.com/mark-by/little-busy-back/pkg/utils"
	"go.uber.org/zap"
	"log"
	"os"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("SECRET"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	client := crmClient.NewGrpcCrmClient(os.Getenv("CRM_ADDRESS"))

	logger, _ := zap.NewDevelopment()
	storage := postgres.NewStorage(&utils.Options{
		MigrationsDir: os.Getenv("MIGRATIONS"),
		User:          "postgres",
		Database:      "postgres",
		Password:      os.Getenv("DB_PASSWORD"),
		Host:          "postgres_bot",
		Port:          "5432",
		Type:          "postgresql",
		Logger:        logger.Sugar().With("bot db"),
	})

	tgHandler := handler.NewTgHandler(client, storage)

	listener := telegramListener.NewTelegramListener(bot, tgHandler)
	if listener == nil {
		log.Panic("fail create listener")
	}

	notifier := scheduler.NewScheduler(bot, storage, client, logger.Sugar().With("scheduler"))
	notifier.Start()

	err = listener.Start()
	if err != nil {
		log.Panic("fail to listen ", err)
	}
}
