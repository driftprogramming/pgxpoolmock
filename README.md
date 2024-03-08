# Mock implementation for `pgx/v5/pgxpool.Pool`

> _Forked from [driftprogramming/pgxpoolmock](https://github.com/omnipeak/pgxpoolmock)_ and updated for pgx/v5

This repo implements a mock for `pgx/v5/pgxpool.Pool` in https://github.com/jackc/pgx, allowing you to test your data
access code locally without actually hitting a real database.

We strongly recommend you use `pgxpool.Pool` rather than `pgx.Conn` directly.

See https://github.com/jackc/pgx/tree/master/pgxpool for more details.

## Installation

You can add the library via go get.

```bash
go get -u github.com/omnipeak/pgxpoolmock/v5
```

### Usage

Basic usage with go's built-in testing.T is as follows:

```go
package pgxpoolmock_test

import (
  "testing"

  "github.com/golang/mock/gomock"
  "github.com/jackc/pgx/v5/pgconn"
  "github.com/stretchr/testify/assert"

  "github.com/omnipeak/pgxpoolmock/v5"
  "github.com/omnipeak/pgxpoolmock/v5/testdata"
)


func TestInsertOrder(t *testing.T) {
  ctrl := gomock.NewController(t)
  defer ctrl.Finish()

  mockPool := pgxpoolmock.NewMockPgxPool(ctrl)
  commandTag := pgconn.NewCommandTag("INSERT 0 1")
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

You can find a more complete example in [`order_dao_example_test.go`](order_dao_example_test.go).
