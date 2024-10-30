package postgres

import (
	"context"
	"errors"
	"fmt"
	"log"
)

func (p *PgStorage) Close(_ context.Context) error {
	err := p.db.Close()
	if err != nil {
		log.Println(errors.Unwrap(err))
	}
	return err
}

func (p *PgStorage) Ping(_ context.Context) error {
	return p.db.Ping()
}

func (p *PgStorage) RunMigrations(ctx context.Context) error {
	log.Println("run migrations")

	_, err := p.db.ExecContext(ctx,
		`create table if not exists users (
			id SERIAL PRIMARY KEY,
			login VARCHAR(255) NOT NULL UNIQUE,
			password VARCHAR(255) NOT NULL,
			created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
			);`,
	)
	if err != nil {
		return fmt.Errorf("can't run migratons: %v", err.Error())
	}
	return nil

}
