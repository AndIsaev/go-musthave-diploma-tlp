package postgres

import (
	"context"
	"errors"
	"log"
)

func (p *PostgresStorage) Close(ctx context.Context) error {
	err := p.conn.Close(ctx)
	if err != nil {
		log.Println(errors.Unwrap(err))
	}
	return err
}
