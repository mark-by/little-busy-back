package telegramListener

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/mark-by/little-busy-back/bot/internal/handler"
	"log"
)

type TelegramListener struct {
	bot     *tgbotapi.BotAPI
	handler handler.Handler
}

func NewTelegramListener(bot *tgbotapi.BotAPI, handler handler.Handler) *TelegramListener {
	return &TelegramListener{
		bot:     bot,
		handler: handler,
	}
}

func (l *TelegramListener) Start() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := l.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "start":
				msg = l.handler.Start(update.Message)
			default:
				msg.Text = "Я не знаю такой команды"
			}
		} else {
			switch update.Message.Text {
			case handler.CommandEvents:
				msg = l.handler.ShowEvents(update.Message)
			case handler.CommandTurnOnNotifications:
				fallthrough
			case handler.CommandTurnOffNotifications:
				msg = l.handler.EnableNotifications(update.Message)
			}

			if update.Message.Contact != nil {
				msg = l.handler.Contact(update.Message)
			}
		}

		if _, err := l.bot.Send(msg); err != nil {
			log.Printf("fail to send message: %s\n", err)
		}
	}

	return nil
}
