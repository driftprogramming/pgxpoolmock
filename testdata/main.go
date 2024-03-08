package testdata

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	connectionString := getConnectionString()
	realPool, _ := pgxpool.New(context.Background(), connectionString)
	orderDao := OrderDAO{
		Pool: realPool,
	}
	order := orderDao.GetOrderByID(100)
	fmt.Print(order)
}

func getConnectionString() string {
	databaseHost := "10.0.20.101"
	databaseUser := "admin"
	databasePassword := "admin_password"
	databaseName := "orderdatabase"
	databaseConnectionPoolMaxSize := 100
	databaseConnectionPoolMinSize := 10
	defaultPoolHealthCheckPeriod := "10s"
	connectionString := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s pool_max_conns=%s pool_min_conns=%s pool_health_check_period=%s",
		databaseHost,
		databaseUser,
		databasePassword,
		databaseName,
		databaseConnectionPoolMaxSize,
		databaseConnectionPoolMinSize,
		defaultPoolHealthCheckPeriod,
	)

	return connectionString
}
