package model

import "time"

type Withdraw struct {
	Order string   `json:"order" db:"number"`
	Price *float64 `json:"sum" db:"price"`
}

type Withdrawal struct {
	Order       string    `json:"order" db:"number"`
	Sum         float64   `json:"sum" db:"price"`
	ProcessedAt time.Time `json:"processed_at" db:"processed_at"`
}
