package model

import "time"

type Balance struct {
	Current   *float64 `json:"current" db:"current"`
	Withdrawn *float64 `json:"withdrawn" db:"withdrawn"`
}

type BalanceWithTime struct {
	Order       string    `json:"order" db:"number"`
	Sum         float64   `json:"sum" db:"price"`
	ProcessedAt time.Time `json:"processed_at" db:"processed_at"`
}
