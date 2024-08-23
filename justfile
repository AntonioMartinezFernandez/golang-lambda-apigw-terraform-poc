test:
  go test -race ./...

cleancache:
  go clean -testcache

run:
  @go run cmd/api/main.go

build:
  @go build -o app cmd/api/main.go

zip:
	@CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o app ./cmd/api
	@zip build/app.zip app
	@rm app

start-infra:
	@docker compose up -d

terraform:
	@terraform -chdir=deployments init
	@terraform -chdir=deployments apply --auto-approve

start-local:
  just zip
  just start-infra
  just terraform

stop-local:
  @terraform -chdir=deployments destroy --auto-approve
  @rm -r deployments/.terraform
  @rm deployments/.terraform*
  @rm deployments/terraform*
  @docker stop $(docker ps -a -q  --filter ancestor=localstack/localstack:3.6)
  @rm build/app.zip
