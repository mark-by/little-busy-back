package entity

type Customer struct {
	ID                  int64   `json:"id"`
	Name                string  `json:"name" validate:"required"`
	Tel                 *string `json:"tel" validate:"omitempty,numeric,startswith=9"`
	SpecialPricePerHour *int    `json:"special_price_per_hour"`
}
