package postgres

import (
	"context"
	"fmt"
	"github.com/AndIsaev/go-musthave-diploma-tlp/internal/model"
	"log"
)

// SetUserOrder - uploading the user's order number
func (p *PgStorage) SetUserOrder(ctx context.Context, params *model.UserOrder) (*model.Order, error) {
	var val model.Order

	query := `INSERT INTO orders (number, user_id, status, uploaded_at) VALUES ($1, $2, $3, now()) RETURNING id, number, user_id, status, uploaded_at;`

	err := p.db.QueryRowContext(ctx, query, params.Number, params.UserID, params.Status).
		Scan(&val.ID, &val.Number, &val.UserID, &val.Status, &val.UploadedAt)
	if err != nil {
		log.Printf("can't set order for user - %v", params.UserLogin.Username)
		return nil, err
	}

	return &val, nil
}

func (p *PgStorage) GetOrderByNumber(ctx context.Context, params *model.UserOrder) (*model.Order, error) {
	var val model.Order
	query := "SELECT id, number, user_id, status, accrual FROM orders WHERE number = $1;"

	err := p.db.GetContext(ctx, &val, query, params.Number)
	if err != nil {
		return nil, err
	}
	return &val, nil
}

func (p *PgStorage) ListOrdersByUserID(ctx context.Context, userID int) (orders []model.Order, err error) {
	query := `SELECT number, status, accrual, uploaded_at FROM orders WHERE user_id = $1 ORDER BY uploaded_at DESC`
	err = p.db.SelectContext(ctx, &orders, query, userID)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}
	return
}

func (p *PgStorage) ListOrders(ctx context.Context) (ids []model.Order, err error) {
	query := `SELECT id, number, status FROM orders WHERE status not in ($1, $2) ORDER BY uploaded_at DESC`
	err = p.db.SelectContext(ctx, &ids, query, model.PROCESSED, model.INVALID)
	if err != nil {
		return nil, err
	}
	return
}

func (p *PgStorage) UpdateOrder(ctx context.Context, order *model.Order) error {

	query := `UPDATE orders SET accrual = $1, status = $2 WHERE id = $3 RETURNING id, number, user_id, status;`

	err := p.db.QueryRowContext(ctx, query, order.Accrual, order.Status, order.ID).
		Scan(&order.ID, &order.Number, &order.UserID, &order.Status)
	if err != nil {
		fmt.Println(err)
		log.Printf("can't update order")
		return err
	}

	return nil
}
