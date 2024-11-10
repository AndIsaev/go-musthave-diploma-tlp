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
}
