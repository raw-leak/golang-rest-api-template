#!/bin/bash
echo ">> Running APP in dev mode..."

# Determine the root directory of your project
ROOT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )/.." && pwd )"

# Set up env vars
source "$ROOT_DIR/env.sh"

# Navigate to the directory where main.go is located
cd "$ROOT_DIR/cmd/golang-rest-api-template"

# Run your Go application by executing main.go directly
go run -race main.go

# Optionally, you can pass any additional arguments or flags to your Go application
# go run -race main.go arg1 arg2 ...
