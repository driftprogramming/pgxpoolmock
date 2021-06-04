### PGX POOL MOCK

This repo is mock `pgxpool` in https://github.com/jackc/pgx. so that you can test your data access code locally, no real database is needed.
I also strongly recommend you use `pgxpool` rather than `pgx` only. see https://github.com/jackc/pgx/tree/master/pgxpool

### How to install

```
go get github.com/driftprogramming/pgxpoolmock
```

### How to Use

See file `pgx_pool_mock_test.go` to figure out how to use it. Or see the below:
```go
package pgxpoolmock_test

import (
	"context"
	"testing"

	"github.com/driftprogramming/pgxpoolmock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type user struct {
	UserName string
	Age      int
}

func TestName(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockPool := pgxpoolmock.NewMockPgxPool(ctrl)
	columns := []string{"user_name", "age"}
	pgxRows := pgxpoolmock.NewRows(columns).AddRow("Leo", 99).AddRow("Tom", 100).ToPgxRows()
	mockPool.EXPECT().Query(gomock.Any(), gomock.Any()).Return(pgxRows, nil)
	// you can pass the `mockPool` into your data access object here.
	actualRows, _ := mockPool.Query(context.Background(), "SELECT MOCK SQL")
	var users []user
	for actualRows.Next() {
		u := &user{}
		actualRows.Scan(&u.UserName, &u.Age)
		users = append(users, *u)
	}

	assert.NotNil(t, users)
	assert.Equal(t, 2, len(users))
}

```
