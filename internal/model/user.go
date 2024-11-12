package model

import "github.com/golang-jwt/jwt/v5"

type User struct {
	ID       int    `json:"id,omitempty" db:"id,omitempty"`
	Login    string `json:"login" db:"login" validate:"required"`
	Password string `json:"password" db:"password" validate:"required,min=8"`
}

type UserWithToken struct {
	ID    int    `json:"id" db:"id,omitempty"`
	Login string `json:"login" db:"login" validate:"required"`
	Token string `json:"token"`
}

type UserLogin struct {
	Username string
}

type UserResponse struct {
	ID    int    `json:"id,omitempty" db:"id,omitempty"`
	Login string `json:"login" db:"login"`
}

type Claims struct {
	Login string `json:"login"`
	jwt.RegisteredClaims
}

type UserToken struct {
	Token string
}

type AuthParams struct {
	Login    string
	Password string
}
