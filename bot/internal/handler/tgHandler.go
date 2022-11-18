package handler

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/mark-by/little-busy-back/bot/internal/crmClient"
	"github.com/mark-by/little-busy-back/bot/internal/entity"
	"github.com/mark-by/little-busy-back/bot/internal/storage"
	"github.com/nleeper/goment"
	"log"
)

var (
	CommandEvents               = "Ближайшие сеансы"
	CommandTurnOnNotifications  = "Включить уведомления"
	CommandTurnOffNotifications = "Выключить уведомления"

	errorMessage = "К сожалению, что-то пошло не так попробуйте позже 🧐"
)

type TgHandler struct {
	crm     crmClient.CrmClient
	storage storage.UserStorage
}

func NewTgHandler(crm crmClient.CrmClient, userStorage storage.UserStorage) *TgHandler {
	return &TgHandler{
		crm:     crm,
		storage: userStorage,
	}
}

func (t TgHandler) Start(incomeMessage *tgbotapi.Message) tgbotapi.MessageConfig {
	newMessage := tgbotapi.NewMessage(incomeMessage.Chat.ID,
		fmt.Sprintf("Привет, %s!"+
			"\nЯ помощник массажиста Ирины Быховец. "+
			"Я могу напоминать, о предстоящих сеансах и рассказывать, какие сеансы запланированы."+
			"\nДля продолжения работы мне нужен ваш телефон. Нажмите ниже \"Отправить номер телефона\"", incomeMessage.From.FirstName))

	newMessage.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButtonContact("Отправить номер телефона"),
		),
	)

	return newMessage
}

func (t TgHandler) Contact(incomeMessage *tgbotapi.Message) tgbotapi.MessageConfig {
	phoneNumber := incomeMessage.Contact.PhoneNumber[1:]
	user, err := t.crm.GetUser(phoneNumber)
	if err != nil {
		log.Println("fail to get contact: ", err)
		return tgbotapi.NewMessage(incomeMessage.Chat.ID, "К сожалению, мы не знакомы или у меня нет вашего номера. Обратитесь к Ирине 🧐")
	}

	user.ChatID = incomeMessage.Chat.ID
	err = t.storage.SaveUser(user)
	if err != nil {
		log.Println("err: fail to save: ", err)
		return tgbotapi.NewMessage(incomeMessage.Chat.ID, errorMessage)
	}

	msg := tgbotapi.NewMessage(incomeMessage.Chat.ID, fmt.Sprintf(
		"Отлично! Мы успешно подключились!"+
			"\nЯ могу показать вам ваши ближайшие сеансы и напоминать о приближающихся событиях."+
			"\nВоспользуйтесь кнопками ниже."))

	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(CommandEvents),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(CommandTurnOnNotifications),
		),
	)

	return msg
}

func (t TgHandler) renderEventsMessage(events []entity.Event) string {
	if len(events) == 0 {
		return "В ближайший месяц сеансов нет"
	}
	result := "Ближайшие сеансы:"
	goment.SetLocale("ru")
	for _, event := range events {
		g1, _ := goment.New(event.StartTime)
		g2, _ := goment.New(event.EndTime)
		result += fmt.Sprintf("\n%s до %s", g1.Format("Do в H:mm"), g2.Format("H:mm, MMMM, ddd"))
	}
	return result
}

func (t TgHandler) ShowEvents(incomeMessage *tgbotapi.Message) tgbotapi.MessageConfig {
	user, err := t.storage.UserByChatID(incomeMessage.Chat.ID)
	if err != nil {
		log.Println("err: fail to get user ", incomeMessage.Chat.ID, errorMessage)
		return tgbotapi.NewMessage(incomeMessage.Chat.ID, errorMessage)
	}
	events, err := t.crm.GetFutureEventsForCustomer(user.Tel)
	if err != nil {
		log.Println("err: fail to get events ", errorMessage)
		return tgbotapi.NewMessage(incomeMessage.Chat.ID, errorMessage)
	}
	msg := tgbotapi.NewMessage(incomeMessage.Chat.ID, t.renderEventsMessage(events))

	toggleNotify := CommandTurnOnNotifications
	if user.NotificationIsEnabled {
		toggleNotify = CommandTurnOffNotifications
	}

	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(CommandEvents),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(toggleNotify),
		),
	)

	return msg
}

func (t TgHandler) EnableNotifications(incomeMessage *tgbotapi.Message) tgbotapi.MessageConfig {
	var turnOnNotification bool
	switch incomeMessage.Text {
	case CommandTurnOnNotifications:
		turnOnNotification = true
	case CommandTurnOffNotifications:
		turnOnNotification = false
	default:
		turnOnNotification = false
	}

	err := t.storage.SetNotification(incomeMessage.Chat.ID, turnOnNotification)
	if err != nil {
		log.Println("fail to set notification", err)
		return tgbotapi.NewMessage(incomeMessage.Chat.ID, errorMessage)
	}

	msg := tgbotapi.NewMessage(incomeMessage.Chat.ID, "Сделано!")

	toggleNotify := CommandTurnOnNotifications
	if turnOnNotification {
		msg.Text = "Сделано!\nБуду напоминать за день до сеанса в 18:00"
		toggleNotify = CommandTurnOffNotifications
	}

	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(CommandEvents),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(toggleNotify),
		),
	)

	return msg
}

var _ Handler = &TgHandler{}
