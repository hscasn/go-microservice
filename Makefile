envfile = ./scripts/.env
include $(envfile)
export $(shell sed 's/=.*//' $(envfile))

# Sets up containers needed for test
test-containers-setup:
	./scripts/startcontainer.sh redis '-p 6379:6379 redis:4.0'

# Only runs the tests. Use if you don't want to set up containers
test-only:
	go clean --testcache ./pkg/...
	go test -cover -coverprofile cover.out ./...

# Sets up the dependencies for tests and run them
test:
	make test-containers-setup
	make test-only
	go tool cover -func cover.out


# Run the application
run:
	go run internal/app.go

# Only build the application
build:
	go mod tidy
	go build -o start internal/app.go

# Create the container and deploy the application
deploy:
	./scripts/deploy.sh