package postgres

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

type PostgresStorage struct {
	conn *pgx.Conn
}

func NewPostgresStorage(connString string) (*PostgresStorage, error) {
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		log.Println("unable to connect to database")
		return nil, err
	}

	return &PostgresStorage{conn: conn}, nil
}
