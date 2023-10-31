.PHONY: mock
mock:
	@echo "Running mockgen..."
	@mockgen -source=webook/internal/service/user.go -package=svcmocks -destination=webook/internal/service/mocks/user.mock.go
	@go mod tidy