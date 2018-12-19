envfile = ./scripts/.env
include $(envfile)
export $(shell sed 's/=.*//' $(envfile))

# Sets up containers needed for test
test-containers-setup:
	../../scripts/startcontainer.sh redis '-p 6379:6379 redis:4.0'

# Runs the tests on a CI environment. Use if you don't want to set up containers
test-ci:
	go clean --testcache ./...
	go test -parallel 5 -cover -coverprofile cover.out ./...

# Sets up the dependencies for tests and run them
test:
	make test-containers-setup
	go clean --testcache ./...
	go test -parallel 5 -cover -coverprofile cover.out ./...
	go tool cover -func cover.out


# Run the application
run:
	go run internal/app.go

# Only build the application
build:
	go build -o start internal/app.go

# Create the container and deploy the application
deploy:
	./scripts/deploy.sh

# Removing junk
clean:
	rm -f cover.out start deployments/_k8s_objects.yml