test:
  go test -race ./...

clean:
  go clean -testcache

build:
  @CGO_ENABLED=0 go build -o ./build/app ./cmd/user-api/main.go

zip-lambda:
  @CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o ./app ./cmd/user-api
  @zip ./build/app.zip app
  @rm app

start-infra:
  @docker compose up -d

stop-infra:
  @docker stop $(docker ps -a -q  --filter ancestor=localstack/localstack:3.6)

terraform-deploy:
  @terraform -chdir=deployments init
  @terraform -chdir=deployments apply --auto-approve

terraform-teardown:
  @terraform -chdir=deployments destroy --auto-approve
  @rm -r deployments/.terraform
  @rm deployments/.terraform*
  @rm deployments/terraform*

start-local:
  @cp .env.DEVELOPMENT .env
  @go run cmd/user-api/main.go -dev

start-serverless:
  @just zip-lambda
  @just start-infra
  @just terraform-deploy

update-serverless:
  @just zip-lambda
  @aws --endpoint-url http://localhost:4566 lambda update-function-code --function-name rest_api_golang_lambda --zip-file fileb://build/app.zip

stop-serverless:
  @just terraform-teardown
  @just stop-infra
  @rm build/app.zip
