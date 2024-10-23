package storage

import (
	"context"
)

type Storage interface {
	Close(ctx context.Context) error
	Insert(ctx context.Context) error
	Create(ctx context.Context) error
	InsertBatch(ctx context.Context) error
	Get(ctx context.Context) error
	List(ctx context.Context) error
}
