package handler

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Handler interface {
	Start(incomeMessage *tgbotapi.Message) tgbotapi.MessageConfig
	Contact(incomeMessage *tgbotapi.Message) tgbotapi.MessageConfig
	ShowEvents(incomeMessage *tgbotapi.Message) tgbotapi.MessageConfig
	EnableNotifications(incomeMessage *tgbotapi.Message) tgbotapi.MessageConfig
}
