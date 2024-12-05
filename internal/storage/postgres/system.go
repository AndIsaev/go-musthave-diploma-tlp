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
					number varchar(255) not null ,
					status varchar(255) not null,
				    accrual float,
				    uploaded_at timestamp with time zone, 
					foreign key (user_id) references users (id),
					unique (number, user_id)
				    );
				create table if not exists balance (
					id serial primary key,
					user_id integer not null unique,
				    current float default 0 not null check (current > 0),
					foreign key (user_id) references users (id)
				    );
				create table if not exists withdraw (
					id serial primary key,
					number varchar(255),
				    user_id integer not null,
				    price float default 0 not null check (price > 0),
				    processed_at timestamp with time zone,
				    foreign key (user_id) references users (id)
				    );
`)

	if err != nil {
		return fmt.Errorf("can't run migratons: %v", err.Error())
	}
	return nil
}
