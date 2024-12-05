package postgres

import (
	"context"
	"log"

	"github.com/AndIsaev/go-musthave-diploma-tlp/internal/model"
)

func (p *PgStorage) GetBalance(ctx context.Context, userID int) (*model.Balance, error) {
	balance := &model.Balance{}
	query := `SELECT
    			b.current,
    			SUM(w.price) as withdrawn
			  FROM balance b
			  LEFT JOIN withdraw w on b.user_id = w.user_id
			  WHERE b.user_id = $1
			  GROUP BY b.current
			  `
	err := p.db.GetContext(ctx, balance, query, userID)
	if err != nil {
		log.Printf("error when try select balance of user %v\n", userID)
		return nil, err
	}
	return balance, nil
}

func (p *PgStorage) CreateBalance(ctx context.Context, current float64, userID int) (*model.Balance, error) {
	balance := &model.Balance{}

	query := `INSERT INTO balance (current, user_id) VALUES ($1, $2) RETURNING current;`

	err := p.db.QueryRowContext(ctx, query, current, userID).Scan(&balance.Current)
	if err != nil {
		log.Println("can't insert data of balance")
		return nil, err
	}

	return balance, nil
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

func (p *PgStorage) GetWithdrawnBalance(ctx context.Context, userID int) (*model.Balance, error) {
	balance := &model.Balance{}
	query := `select b.current as current, sum(w.price) as withdrawn FROM balance b
				inner join withdraw as w on w.user_id  = b.user_id 
				where b.user_id = $1
				group by b.current;`
	err := p.db.GetContext(ctx, balance, query, userID)
	if err != nil {
		log.Printf("error when try select balance with withdraw of user %v\n", userID)
		return nil, err
	}
	return balance, nil
}
