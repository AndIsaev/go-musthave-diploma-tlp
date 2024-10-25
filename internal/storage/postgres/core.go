package postgres

import (
	"context"
	"log"

	"errors"

	"github.com/AndIsaev/go-musthave-diploma-tlp/internal/storage"
	"github.com/jackc/pgx/v5"
)

type PostgresStorage struct {
	conn *pgx.Conn
}

func NewPostgresStorage(connString string) (*PostgresStorage, error) {
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		log.Println(errors.Unwrap(err))
		return nil, err
	}

	return &PostgresStorage{conn: conn}, nil
}

func (p *PostgresStorage) System() storage.Inside {
	return p
}
