# Define the target to run your Go application in development mode
run-dev:
	./scripts/rundev.sh

lint:
	golangci-lint run

test:
	go test ./...
