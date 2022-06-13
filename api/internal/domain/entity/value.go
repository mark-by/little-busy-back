package entity

import "time"

type Value struct {
	Incomes  float32   `json:"incomes"`
	Costs    float32   `json:"costs"`
	DateTime time.Time `json:"date_time"`
}
