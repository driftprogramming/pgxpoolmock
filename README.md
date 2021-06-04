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
