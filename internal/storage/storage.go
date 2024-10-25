package storage

import (
	"context"
)

type Storage interface {
	System() Inside
}

type Inside interface {
	Close(ctx context.Context) error
}
