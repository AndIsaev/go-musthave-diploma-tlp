package service

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/AndIsaev/go-musthave-diploma-tlp/internal/exception"
	"github.com/AndIsaev/go-musthave-diploma-tlp/internal/model"
	"github.com/AndIsaev/go-musthave-diploma-tlp/internal/storage"
)

//go:generate mockgen -destination=mocks/mock_service.go -package=mocks github.com/AndIsaev/go-musthave-diploma-tlp/internal/service Service
type Service interface {
	Register(ctx context.Context, params *model.AuthParams) (*model.UserWithToken, error)
	Login(ctx context.Context, params *model.AuthParams) (*model.UserWithToken, error)
	SetOrder(ctx context.Context, params *model.UserOrder) (*model.Order, error)
	GetUserOrders(ctx context.Context, login *model.UserLogin) (orders []model.Order, err error)
	DeductPoints(ctx context.Context, withdraw *model.Withdraw, login *model.UserLogin) (*model.Withdraw, error)
	GetUserBalance(ctx context.Context, login *model.UserLogin) (*model.Balance, error)
	GetUserWithdrawals(ctx context.Context, login *model.UserLogin) ([]model.Withdrawal, error)
}

type Methods struct {
	Storage storage.Storage
}

// Register - need for register new user
func (s *Methods) Register(ctx context.Context, params *model.AuthParams) (*model.UserWithToken, error) {
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

// Login - check permissions for user and return token
func (s *Methods) Login(ctx context.Context, params *model.AuthParams) (*model.UserWithToken, error) {
	user, err := s.Storage.User().Auth(ctx, params)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errors.New("user doesn't exists")
	} else if err != nil {
		log.Println("can't select query")
		return nil, err
	}

	return user, nil
}

// SetOrder - create new order for check accrual system
func (s *Methods) SetOrder(ctx context.Context, params *model.UserOrder) (*model.Order, error) {
	user, err := s.Storage.User().GetUserByLogin(ctx, &params.UserLogin)
	if err != nil {
		log.Println("user with such login not found")
		return nil, err
	}
	params.UserID = user.ID

	existsOrder, err := s.Storage.Order().GetOrderByNumber(ctx, params)
	if errors.Is(err, sql.ErrNoRows) {
		params.Status = model.NEW
		newOrder, err := s.Storage.Order().SetUserOrder(ctx, params)
		if err != nil {
			log.Println("error when try to set new order of user")
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

// GetUserOrders - get orders of user with statuses
func (s *Methods) GetUserOrders(ctx context.Context, login *model.UserLogin) (orders []model.Order, err error) {
	user, err := s.Storage.User().GetUserByLogin(ctx, login)
	if err != nil {
		log.Println("error when try to select user by login")
		return nil, err
	}

	orders, err = s.Storage.Order().ListOrdersByUserID(ctx, user.ID)
	if err != nil {
		log.Println("error when trying to receive user orders")
		return
	}
	return
}

// GetUserBalance - get balance of user with current and withdrawn
func (s *Methods) GetUserBalance(ctx context.Context, login *model.UserLogin) (*model.Balance, error) {
	user, err := s.Storage.User().GetUserByLogin(ctx, login)
	if err != nil {
		log.Println("error when try to select user by login")
		return nil, err
	}

	balance, err := s.Storage.Balance().GetBalance(ctx, user.ID)
	if err != nil {
		log.Println("error when try to select balance")
		return nil, err
	}

	return balance, nil
}

// DeductPoints - purchase an order for points
func (s *Methods) DeductPoints(ctx context.Context, withdraw *model.Withdraw, login *model.UserLogin) (*model.Withdraw, error) {
	user, err := s.Storage.User().GetUserByLogin(ctx, login)
	if err != nil {
		log.Println("error when try to select user by login")
		return nil, err
	}

	balance, err := s.Storage.Balance().GetBalance(ctx, user.ID)
	if err != nil {
		log.Println("error when try to select balance")
		return nil, err
	}
	if *balance.Current < *withdraw.Price {
		return nil, exception.ErrNotEnoughBonuses
	}

	current := *balance.Current - *withdraw.Price
	err = s.Storage.Balance().UpdateBalance(ctx, current, user.ID)
	if err != nil {
		log.Println("error when try to create balance")
		return nil, err
	}

	newWithdraw, err := s.Storage.Withdraw().CreateWithdraw(ctx, withdraw, user.ID)
	if err != nil {
		log.Println("error when try to create withdraw")
		return nil, err
	}

	return newWithdraw, nil
}

// GetUserWithdrawals - get history withdrawals of user
func (s *Methods) GetUserWithdrawals(ctx context.Context, login *model.UserLogin) (values []model.Withdrawal, err error) {
	user, err := s.Storage.User().GetUserByLogin(ctx, login)
	if err != nil {
		return nil, err
	}

	values, err = s.Storage.Withdraw().GetListWithdrawnBalance(ctx, user.ID)
	if err != nil {
		log.Println("error when trying to get list withdrawn")
		return
	}
	return
}
