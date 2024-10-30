package postgres

import (
	"context"
	"log"

	"github.com/AndIsaev/go-musthave-diploma-tlp/internal/model"
)

func (p *PgStorage) GetUserByLogin(ctx context.Context, userLogin *model.UserLogin) (*model.User, error) {
	val := model.User{}
	query := "SELECT id, login FROM users WHERE login = $1;"

	err := p.db.GetContext(ctx, &val, query, userLogin.Username)
	if err != nil {
		return nil, err
	}
	return &val, nil
}

func (p *PgStorage) CreateUser(ctx context.Context, user *model.User) (*model.UserResponse, error) {
	val := model.UserResponse{}
	query := `INSERT INTO users (login, password) VALUES ($1, $2) RETURNING id, login;`

	err := p.db.QueryRowContext(ctx, query, user.Login, user.Password).Scan(&val.ID, &val.Login)
	if err != nil {
		log.Println("can't insert data of user")
		return nil, err
	}
	return &val, nil
}
