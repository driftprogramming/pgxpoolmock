### PGX POOL MOCK

This repo is mock `pgxpool` in https://github.com/jackc/pgx. so that you can test your data access code locally, no real database is needed.
I also strongly recommend you use `pgxpool` rather than `pgx` only. see https://github.com/jackc/pgx/tree/master/pgxpool

### How to install

```
go get -u github.com/driftprogramming/pgxpoolmock
```

### How to Use

See file `order_dao_example_test.go` to figure out how to use it. Or see the below:
```go
package pgxpoolmock_test

import (
	"testing"

	"github.com/driftprogramming/pgxpoolmock"
	"github.com/driftprogramming/pgxpoolmock/testdata"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)


func TestInsertOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPool := pgxpoolmock.NewMockPgxPool(ctrl)
	commandTag := pgconn.CommandTag("INSERT 0 1")
	mockPool.EXPECT().Exec(gomock.Any(), "insert into order (id, price) values ($1, $2)", 1, 2.3).Return(commandTag, nil)
	orderDao := testdata.OrderDAO{
		Pool: mockPool,
	}

	order := &testdata.Order{
		ID:    1,
		Price: 2.3,
	}
	err := orderDao.InsertOrder(order)
	if err != nil {
		t.Fatalf("%v", err)
	}
}

func TestGetOneOrderByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockPool := pgxpoolmock.NewMockPgxPool(ctrl)
	columns := []string{"id", "price"}
	pgxRow := pgxpoolmock.NewRow(columns, 1, 2.3)
	mockPool.EXPECT().QueryRow(gomock.Any(), "select id, price from order where id = $1", 1).Return(pgxRow)
	orderDao := testdata.OrderDAO{
		Pool: mockPool,
	}
	actualOrder, err := orderDao.GetOneOrderByID(1)
	if err != nil {
		t.Fatalf("%v", err)
	}
	assert.NotNil(t, actualOrder)
	assert.Equal(t, 1, actualOrder.ID)
	assert.Equal(t, 2.3, actualOrder.Price)
}

func TestName(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// given
	mockPool := pgxpoolmock.NewMockPgxPool(ctrl)
	columns := []string{"id", "price"}
	pgxRows := pgxpoolmock.NewRows(columns).AddRow(100, 100000.9).ToPgxRows()
	mockPool.EXPECT().Query(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgxRows, nil)
	orderDao := testdata.OrderDAO{
		Pool: mockPool,
	}

	// when
	actualOrder := orderDao.GetOrderByID(1)

	// then
	assert.NotNil(t, actualOrder)
	assert.Equal(t, 100, actualOrder.ID)
	assert.Equal(t, 100000.9, actualOrder.Price)
}

```
