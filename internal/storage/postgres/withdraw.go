package postgres

import (
	"context"
	"github.com/AndIsaev/go-musthave-diploma-tlp/internal/model"
	"log"
)

//
//func (p *PgStorage) GetWithdraw(ctx context.Context, userID int) (*model.Balance, error) {
//	var val model.Balance
//	query := `SELECT current FROM balance WHERE user_id = $1`
//	err := p.db.GetContext(ctx, &val, query, userID)
//	if err != nil {
//		fmt.Println("00000000000")
//		fmt.Println(err)
//
//		fmt.Println("00000000000")
//
//		return nil, err
//	}
//	return &val, nil
//}

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
