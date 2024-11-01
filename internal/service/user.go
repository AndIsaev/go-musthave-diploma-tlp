package service

import (
	"context"
	"database/sql"
	"errors"
	"github.com/AndIsaev/go-musthave-diploma-tlp/internal/model"
	"github.com/AndIsaev/go-musthave-diploma-tlp/internal/storage"
	"log"
)

type Service interface {
	Register(ctx context.Context, params *model.AuthParams) (*model.UserWithToken, error)
	Login(ctx context.Context, params *model.AuthParams) (*model.UserWithToken, error)
}

type UserMethods struct {
	Storage storage.Storage
}

func (s *UserMethods) Register(ctx context.Context, params *model.AuthParams) (*model.UserWithToken, error) {
	user, err := s.Storage.User().GetUserByLogin(ctx, &model.UserLogin{Username: params.Login})
	if user != nil || !errors.Is(err, sql.ErrNoRows) {
		return nil, errors.New("can't register new user")
	}

	createdUser, err := s.Storage.User().CreateUser(ctx, params)
	if err != nil {
		log.Println(createdUser, err)
		return nil, err
	}

	return createdUser, nil
}

func (s *UserMethods) Login(ctx context.Context, params *model.AuthParams) (*model.UserWithToken, error) {
	user, err := s.Storage.User().Login(ctx, params)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errors.New("user doesn't exists")
	} else if err != nil {
		log.Println("can't select query")
		return nil, err
	}

	return user, nil
}
