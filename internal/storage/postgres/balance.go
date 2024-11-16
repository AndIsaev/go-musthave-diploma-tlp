package postgres

import (
	"context"
	"fmt"
	"github.com/AndIsaev/go-musthave-diploma-tlp/internal/model"
	"log"
)

func (p *PgStorage) GetBalance(ctx context.Context, userID int) (*model.Balance, error) {
	var val model.Balance
	query := `SELECT current FROM balance WHERE user_id = $1`
	err := p.db.GetContext(ctx, &val, query, userID)
	if err != nil {
		fmt.Println("00000000000")
		fmt.Println(err)

		fmt.Println("00000000000")

		return nil, err
	}
	return &val, nil
}

func (p *PgStorage) CreateBalance(ctx context.Context, current float64, userID int) (*model.Balance, error) {
	var val model.Balance
	fmt.Println(current)
	query := `INSERT INTO balance (current, user_id) VALUES ($1, $2) RETURNING current;`

	err := p.db.QueryRowContext(ctx, query, current, userID).Scan(&val.Current)
	if err != nil {
		log.Println("can't insert data of balance")
		return nil, err
	}

	return &val, nil
}

func (p *PgStorage) UpdateBalance(ctx context.Context, current float64, userID int) error {

	query := `UPDATE balance SET current = $1 WHERE user_id = $2;`

	_, err := p.db.ExecContext(ctx, query, current, userID)
	if err != nil {
		log.Println("can't update data of balance")
		return err
	}

	return nil
}
