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
					id serial  primary key,
					login varchar(255) not null unique,
					password varchar(255) not null,
					created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
					);
				create table if not exists orders (
					id serial primary key,
					user_id integer not null,
					number bigint not null ,
					foreign key (user_id) references users (id),
					unique (number, user_id)
				    );

`)

	if err != nil {
		return fmt.Errorf("can't run migratons: %v", err.Error())
	}
	return nil

}
