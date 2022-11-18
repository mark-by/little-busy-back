package crmClient

import "github.com/mark-by/little-busy-back/bot/internal/entity"

type CrmClient interface {
	GetTomorrowEvents() ([]entity.Event, error)
	GetFutureEventsForCustomer(customerTel string) ([]entity.Event, error)
	GetUser(tel string) (*entity.User, error)
}
