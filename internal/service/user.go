package service

import (
	"context"
	"database/sql"
	"errors"
	"github.com/AndIsaev/go-musthave-diploma-tlp/internal/exception"
	"github.com/AndIsaev/go-musthave-diploma-tlp/internal/model"
	"github.com/AndIsaev/go-musthave-diploma-tlp/internal/storage"
	"log"
)

type Service interface {
	Register(ctx context.Context, params *model.AuthParams) (*model.UserWithToken, error)
	Login(ctx context.Context, params *model.AuthParams) (*model.UserWithToken, error)
	SetOrder(ctx context.Context, params *model.UserOrder) (*model.Order, error)
	GetUserOrders(ctx context.Context, login *model.UserLogin) (orders []model.Order, err error)
	//SetNewBalance(ctx context.Context, login *model.UserLogin, balance *model.Balance) (*model.Balance, error)
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

func (s *UserMethods) SetOrder(ctx context.Context, params *model.UserOrder) (*model.Order, error) {
	user, err := s.Storage.User().GetUserByLogin(ctx, &params.UserLogin)
	if err != nil {
		log.Println("user with such login not found")
		return nil, err
	}
	params.UserID = user.ID

	existsOrder, err := s.Storage.User().GetOrderByNumber(ctx, params)
	if errors.Is(err, sql.ErrNoRows) {
		params.Status = model.NEW
		newOrder, err := s.Storage.User().SetUserOrder(ctx, params)
		if err != nil {
			return nil, err
		}
		return newOrder, nil
	}

	if existsOrder.UserID != params.UserID {
		log.Println("the order already set for another user")
		return nil, exception.ErrOrderAlreadyExistsAnotherUser

	}

	if existsOrder.UserID == params.UserID {
		log.Println("the order already set for this user")
		return nil, exception.ErrOrderAlreadyExists
	}
	return nil, err

}

func (s *UserMethods) GetUserOrders(ctx context.Context, login *model.UserLogin) (orders []model.Order, err error) {
	user, err := s.Storage.User().GetUserByLogin(ctx, login)
	if err != nil {
		return nil, err
	}

	orders, err = s.Storage.User().ListOrdersById(ctx, user.ID)
	if err != nil {
		log.Println("error when trying to receive user orders")
		return
	}
	return
}

func (s *UserMethods) GetUserBalance(ctx context.Context, login *model.UserLogin) (*model.Balance, error) {
	user, err := s.Storage.User().GetUserByLogin(ctx, login)
	if err != nil {
		return nil, err
	}

	balance, err := s.Storage.User().GetBalance(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	return balance, nil
}
