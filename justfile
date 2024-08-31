# Run all tests, or any tests specified by the path with its extra parameters
test path="./..." *params="":
    go test {{path}} -race {{params}}

# Runs all tests located at ./test
test-integration *params:
    @just test ./test/... -timeout 300s {{params}}

# Runs all tests, except integration tests located at ./test
test-unit *params:
    go test -p 2 $(go list ./... | grep -v ./test) -race {{params}}

# Clears the test cache
clear-cache:
    go clean -testcache

# Formats all go files
lint:
    go fmt ./...

# Build app
build:
  @CGO_ENABLED=0 go build -o ./build/app ./cmd/user-api/main.go

# Build app and zip binary and assets
zip-lambda:
  @go mod tidy
  @CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o ./app ./cmd/user-api
  @zip ./build/app.zip app ./schemas/*
  @rm app

# Start necessary services to run the integration tests or run the application in development mode
start-infra-services:
  @docker compose up -d

# Removes the setup created by the start-infra-services task
stop-infra-services:
  @docker stop $(docker ps -a -q  --filter ancestor=localstack/localstack:3.6)

# Deploy infra dependencies
terraform-deploy:
  @terraform -chdir=deployments init
  @terraform -chdir=deployments apply --auto-approve

# Remove infra dependencies
terraform-teardown:
  @terraform -chdir=deployments destroy --auto-approve
  @rm -r deployments/.terraform
  @rm deployments/.terraform*
  @rm deployments/terraform*

# Run app locally as independent service
start-local:
  @cp .env.DEVELOPMENT .env
  @go mod tidy
  @go run cmd/user-api/main.go

# Run app locally as lambda
start-serverless:
  @just zip-lambda
  @just start-infra-services
  @just terraform-deploy

# Update deployed lambda with latest changes
update-serverless:
  @just zip-lambda
  @aws --endpoint-url http://localhost:4566 lambda update-function-code --function-name rest_api_golang_lambda --zip-file fileb://build/app.zip

# Stop lambda and associated infra
stop-serverless:
  @just terraform-teardown
  @just stop-infra-services
  @rm build/app.zip
