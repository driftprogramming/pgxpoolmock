package testdata

import (
	"context"

	"github.com/driftprogramming/pgxpoolmock"
)

type Order struct {
	ID    int
	Price float64
}

type OrderDAO struct {
	Pool pgxpoolmock.PgxPool
}

func (dao *OrderDAO) GetOrderByID(id int) *Order {
	rows, _ := dao.Pool.Query(context.Background(), "SELECT ID,Price FROM order WHERE ID =$1", id)
	for rows.Next() {
		order := &Order{}
		rows.Scan(&order.ID, &order.Price)
		return order
	}

	return nil
}
