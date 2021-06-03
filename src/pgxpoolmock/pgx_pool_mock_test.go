package pgxpoolmock_test

import (
	"context"
	"pgxpoolmock/src/pgxpoolmock"
	"testing"

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
	// make sure you pass the `mockPool` to your data access object.
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
