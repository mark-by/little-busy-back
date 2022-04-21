package application

import "github.com/mark-by/little-busy-back/api/internal/domain/entity"

func getCustomerIdsForEvents(events []entity.Event) []int64 {
	var ids []int64
	for _, event := range events {
		if event.CustomerID == nil {
			continue
		}
		ids = append(ids, *event.CustomerID)
	}
	return ids
}

func createCustomerMap(customers []entity.Customer) map[int64]entity.Customer {
	customerMap := make(map[int64]entity.Customer)

	for _, customer := range customers {
		customerMap[customer.ID] = customer
	}

	return customerMap
}

func mergeCustomersEvents(customerMap map[int64]entity.Customer, events []entity.Event) []entity.Event {
	for idx := 0; idx < len(events); idx++ {
		customer, ok := customerMap[*events[idx].CustomerID]
		if !ok {
			continue
		}
		events[idx].Customer = &customer
	}
	return events
}
