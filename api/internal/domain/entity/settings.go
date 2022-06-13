package entity

type Settings struct {
	StartWorkHour       int `json:"start_work_hour" validate:"required"`
	EndWorkHour         int `json:"end_work_hour" validate:"required"`
	DefaultPricePerHour int `json:"default_price_per_hour" validate:"required,gte=0"`
}
