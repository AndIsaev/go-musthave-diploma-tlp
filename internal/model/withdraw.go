package model

type Withdraw struct {
	Order string   `json:"order" db:"number"`
	Price *float64 `json:"sum" db:"price"`
}
