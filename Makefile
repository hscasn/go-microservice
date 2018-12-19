envfile = ./scripts/.env
include $(envfile)
export $(shell sed 's/=.*//' $(envfile))

# Sets up containers needed for test
test-containers-setup:
	./scripts/startcontainer.sh redis '-p 6379:6379 redis:4.0'

# Runs the tests on a CI environment. Use if you don't want to set up containers
test-ci:
	go clean --testcache ./pkg/...
	go test -parallel 5 -cover -coverprofile cover.out ./pkg/...
	./scripts/testapps.sh ci

# Sets up the dependencies for tests and run them
test:
	make test-containers-setup
	go clean --testcache ./pkg/...
	go test -parallel 5 -cover -coverprofile cover.out ./pkg/...
	./scripts/testapps.sh
	go tool cover -func cover.out

# Create the containers and deploy the apps modified on the last commit
deploy:
	./scripts/deploychangedapps.sh

# Removing junk
clean:
	rm -f cover.out
	./scripts/cleanapps.sh