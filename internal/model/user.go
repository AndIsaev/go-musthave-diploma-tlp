package model

type User struct {
	ID       int    `json:"id,omitempty" db:"id,omitempty"`
	Login    string `json:"login" db:"login" validate:"required"`
	Password string `json:"password" db:"password" validate:"required,min=8"`
}

type UserLogin struct {
	Username string
}

type UserResponse struct {
	ID    int    `json:"id,omitempty" db:"id,omitempty"`
	Login string `json:"login" db:"login"`
}
