package storage

import (
	"context"
)

type Storage interface {
	System() Inside
}

type Inside interface {
	Close(ctx context.Context) error
	Ping(ctx context.Context) error
	RunMigrations(ctx context.Context) error
}
