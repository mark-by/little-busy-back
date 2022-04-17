package application

import (
	"github.com/mark-by/little-busy-back/api/internal/domain/entity"
)

func (e Events) addCustomers(events []entity.Event) ([]entity.Event, error) {
	customerIds := getCustomerIdsForEvents(events)
	if len(customerIds) == 0 {
		return events, nil
	}
	customers, err := e.repoCustomers.GetCustomers(customerIds)
	if err != nil {
		return nil, err
	}

	customerMap := createCustomerMap(customers)
	newEvents := mergeCustomersEvents(customerMap, events)

	return newEvents, nil
}
