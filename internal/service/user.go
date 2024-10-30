package service

import (
	"context"
	"errors"
	"github.com/AndIsaev/go-musthave-diploma-tlp/internal/model"
	"github.com/AndIsaev/go-musthave-diploma-tlp/internal/storage"
)

type RegisterParams struct {
	Login    string
	Password string
}

type Service interface {
	Register(ctx context.Context, params *RegisterParams) (*model.UserResponse, error)
}

type UserMethods struct {
	Storage storage.Storage
}

func (s *UserMethods) Register(ctx context.Context, params *RegisterParams) (*model.UserResponse, error) {
	user, err := s.Storage.User().GetUserByLogin(ctx, &model.UserLogin{Username: params.Login})
	if user != nil && err != nil {
		return nil, errors.New("can't register new user")
	}
	hashedPassword, _ := HashPassword(params.Password)
	createdUser, err := s.Storage.User().CreateUser(ctx, &model.User{Login: params.Login, Password: hashedPassword})
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}
