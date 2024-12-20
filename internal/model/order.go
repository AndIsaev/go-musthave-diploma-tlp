package model

import "time"

type Status string

const (
	NEW       Status = "NEW"
	INVALID   Status = "INVALID"
	PROCESSED Status = "PROCESSED"
)

type Order struct {
	ID         *int      `json:"id,omitempty" db:"id"`
	UserID     int       `json:"user_id,omitempty" db:"user_id"`
	Number     string    `json:"number" db:"number"`
	Status     Status    `json:"status" db:"status"`
	Accrual    *float64  `json:"accrual,omitempty" db:"accrual,omitempty"`
	UploadedAt time.Time `json:"uploaded_at" db:"uploaded_at"`
}

type UserOrder struct {
	ID        int    `json:"id" db:"id"`
	UserID    int    `json:"user_id" db:"user_id"`
	Number    string `json:"number" db:"number"`
	Status    Status `json:"status" db:"status"`
	UserLogin UserLogin
}
