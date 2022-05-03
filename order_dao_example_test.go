package pgxpoolmock_test

import (
	"fmt"
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

func TestMap(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// given
	mockPool := pgxpoolmock.NewMockPgxPool(ctrl)
	columns := []string{"id", "price"}
	pgxRows := pgxpoolmock.NewRows(columns).AddRow(100, 100000.9).ToPgxRows()
	assert.NotEqual(t, "with empty rows", fmt.Sprintf("%s", pgxRows))

	mockPool.EXPECT().Query(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgxRows, nil)
	orderDao := testdata.OrderDAO{
		Pool: mockPool,
	}

	// when
	actualOrder := orderDao.GetOrderMapByID(1)

	// then
	assert.NotNil(t, actualOrder)
	assert.Equal(t, 100, actualOrder["ID"])
	assert.Equal(t, 100000.9, actualOrder["Price"])
}
