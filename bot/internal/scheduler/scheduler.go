package scheduler

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/mark-by/little-busy-back/bot/internal/crmClient"
	"github.com/mark-by/little-busy-back/bot/internal/storage"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"time"
)

type Scheduler struct {
	bot     *tgbotapi.BotAPI
	storage storage.UserStorage
	crm     crmClient.CrmClient
	logger  *zap.SugaredLogger
}

func NewScheduler(bot *tgbotapi.BotAPI, storage storage.UserStorage, crm crmClient.CrmClient, logger *zap.SugaredLogger) *Scheduler {
	return &Scheduler{
		bot:     bot,
		storage: storage,
		crm:     crm,
		logger:  logger,
	}
}

func (s Scheduler) Start() {
	location, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		s.logger.Error("fail to create location: ", err)
		return
	}

	scheduler := cron.New(cron.WithLocation(location))
	_, err = scheduler.AddFunc("0 18 * * *", s.notify)
	if err != nil {
		s.logger.Error("fail to create job: ", err)
		return
	}

	go scheduler.Start()
}

func (s Scheduler) notify() {
	events, err := s.crm.GetTomorrowEvents()
	if err != nil {
		s.logger.Error("fail to get events: ", err)
		return
	}

	for _, event := range events {
		user, err := s.storage.UserByTel(event.ClientTel)
		if err != nil {
			s.logger.With("tel", event.ClientTel).Error("fail to get user by tel: ", err)
			continue
		}
		if user == nil || !user.NotificationIsEnabled {
			continue
		}

		_, err = s.bot.Send(tgbotapi.NewMessage(user.ChatID, fmt.Sprintf(
			"Добрый вечер!\n"+
				"Завтра в %s у вас назначен сеанс", event.StartTime.Format("15:04"))))
		if err != nil {
			s.logger.With("chat", user.ChatID).Error("fail to send message: ", err)
		}
	}
}
