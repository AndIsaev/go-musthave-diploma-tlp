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

func (p *PgStorage) GetListWithdrawnBalance(ctx context.Context, userID int) (values []model.Withdrawal, err error) {
	query := `select number, price, processed_at from withdraw
				where user_id = $1
				order by processed_at desc;`
	err = p.db.SelectContext(ctx, &values, query, userID)
	if err != nil {
		return nil, err
	}
	return
}
