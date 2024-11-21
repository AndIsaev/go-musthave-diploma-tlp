package storage

import (
	"context"

	"github.com/AndIsaev/go-musthave-diploma-tlp/internal/model"
)

type Storage interface {
	System() SystemRepository
	User() UserRepository
}

type SystemRepository interface {
	Close(ctx context.Context) error
	Ping(ctx context.Context) error
	RunMigrations(ctx context.Context) error
}

type UserRepository interface {
	GetUserByLogin(context.Context, *model.UserLogin) (*model.User, error)
	CreateUser(ctx context.Context, user *model.AuthParams) (*model.UserWithToken, error)
	Login(ctx context.Context, params *model.AuthParams) (*model.UserWithToken, error)
	SetUserOrder(ctx context.Context, params *model.UserOrder) (*model.Order, error)
	GetOrderByNumber(ctx context.Context, params *model.UserOrder) (*model.Order, error)
	ListOrders(ctx context.Context) ([]model.Order, error)
	UpdateOrder(ctx context.Context, order *model.Order) error
	ListOrdersByUserID(ctx context.Context, userID int) (orders []model.Order, err error)
	GetBalance(ctx context.Context, userID int) (balance *model.Balance, err error)
	CreateBalance(ctx context.Context, current float64, userID int) (*model.Balance, error)
	UpdateBalance(ctx context.Context, current float64, userID int) error
	CreateWithdraw(ctx context.Context, withdraw *model.Withdraw, userID int) (*model.Withdraw, error)
	GetListWithdrawnBalance(ctx context.Context, userID int) (values []model.Withdrawal, err error)
}
