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

func (dao *OrderDAO) InsertOrder(order *Order) error {
	_, err := dao.Pool.Exec(context.Background(), "insert into order (id, price) values ($1, $2)", order.ID, order.Price)
	if err != nil {
		return err
	}
	return nil
}

func (dao *OrderDAO) GetOneOrderByID(id int) (*Order, error) {
	order := &Order{}
	err := dao.Pool.QueryRow(context.Background(), "select id, price from order where id = $1", id).
		Scan(&order.ID, &order.Price)
	if err != nil {
		return nil, err
	}
	return order, nil
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

func (dao *OrderDAO) GetOrderMapByID(id int) map[string]interface{} {
	rows, _ := dao.Pool.Query(context.Background(), "SELECT ID,Price FROM order WHERE ID =$1", id)
	for rows.Next() {
		var id, price interface{}
		rows.Scan(&id, &price)
		order := map[string]interface{}{
			"ID":    id,
			"Price": price,
		}
		return order
	}

	return nil
}
