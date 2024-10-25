package model

type User struct {
	ID       int    `json:"id,omitempty" db:"id,omitempty"`
	Login    string `json:"login" db:"login"`
	Password string `json:"password" db:"password"`
}
