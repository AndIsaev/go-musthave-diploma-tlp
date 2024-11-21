package postgres

import (
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"

	"github.com/AndIsaev/go-musthave-diploma-tlp/internal/storage"
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

func (p *PgStorage) Order() storage.OrderRepository {
	return p
}

func (p *PgStorage) Balance() storage.BalanceRepository {
	return p
}

func (p *PgStorage) Withdraw() storage.WithdrawRepository {
	return p
}
