.PHONY: all build run test clean watch docker-run docker-down itest

# Regenerate the mock
mock:
	@echo "Generating mock..."
	@echo "Installing mockgen from go.uber.org/mock/mockgen..."
	@go install go.uber.org/mock/mockgen@latest
	@mockgen -source=pgx_pool.go -destination="pgx_pool_mock.go" -package=pgxpoolmock