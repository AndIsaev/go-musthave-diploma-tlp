package model

type Balance struct {
	Current   float64 `json:"current,omitempty" db:"current,omitempty"`
	Withdrawn float64 `json:"withdrawn,omitempty" db:"withdrawn,omitempty"`
}
