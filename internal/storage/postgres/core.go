package postgres

import (
	"github.com/AndIsaev/go-musthave-diploma-tlp/internal/storage"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"log"
)

type PgStorage struct {
	db *sqlx.DB
}

func NewPgStorage(connString string) (*PgStorage, error) {
	conn, err := sqlx.Connect("pgx", connString)
	if err != nil {
		log.Printf("unable to connect to database: %v\n", err)
		return nil, err
	}

	return &PgStorage{db: conn}, nil
}

func (p *PgStorage) System() storage.SystemRepository {
	return p
}

func (p *PgStorage) User() storage.UserRepository {
	return p
}
