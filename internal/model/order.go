package model

type Status string

const (
	NEW        Status = "NEW"
	PROCESSING Status = "PROCESSING"
	INVALID    Status = "INVALID"
	PROCESSED  Status = "PROCESSED"
)

type Order struct {
	ID      int    `json:"id" db:"id"`
	UserID  int    `json:"user_id" db:"user_id"`
	Number  string `json:"number" db:"number"`
	Status  Status `json:"status" db:"status"`
	Accrual *int   `json:"accrual,omitempty" db:"accrual,omitempty"`
}

type UserOrder struct {
	ID        int    `json:"id" db:"id"`
	UserID    int    `json:"user_id" db:"user_id"`
	Number    string `json:"number" db:"number"`
	Status    Status `json:"status" db:"status"`
	UserLogin UserLogin
}
