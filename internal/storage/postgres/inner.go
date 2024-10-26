package postgres

import (
	"context"
	"errors"
	"fmt"
	"log"
)

func (p *PostgresStorage) Close(ctx context.Context) error {
	err := p.conn.Close(ctx)
	if err != nil {
		log.Println(errors.Unwrap(err))
	}
	return err
}

func (p *PostgresStorage) Ping(ctx context.Context) error {
	return p.conn.Ping(ctx)
}

func (p *PostgresStorage) RunMigrations(ctx context.Context) error {
	log.Println("run migrations")

	_, err := p.conn.Exec(ctx,
		`create table if not exists users (
			id SERIAL PRIMARY KEY,
			login VARCHAR(255) NOT NULL,
			password VARCHAR(255) NOT NULL,
			created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
			);`,
	)
	if err != nil {
		return fmt.Errorf("can't run migratons: %v", err.Error())
	}
	return nil

}
