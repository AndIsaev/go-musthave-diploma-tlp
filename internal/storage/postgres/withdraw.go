package postgres

import (
	"context"
	"github.com/AndIsaev/go-musthave-diploma-tlp/internal/model"
	"log"
)

func (p *PgStorage) CreateWithdraw(ctx context.Context, withdraw *model.Withdraw, userID int) (*model.Withdraw, error) {
	val := &model.Withdraw{}
	query := `INSERT INTO withdraw (number, price, user_id, processed_at) VALUES ($1, $2, $3, now()) RETURNING number, price;`

	err := p.db.QueryRowContext(ctx, query, withdraw.Order, withdraw.Price, userID).Scan(&val.Order, &val.Price)
	if err != nil {
		log.Println("can't insert data of withdraw")
		return nil, err
	}

	return val, nil
}
