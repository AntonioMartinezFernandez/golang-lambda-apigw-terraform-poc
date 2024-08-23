# golang-lambda-apigw-terraform-poc

Deploying Golang lambda and API Gateway in Localstack with Terraform

## Requirements

1. [Golang](https://go.dev/dl/)
2. [Just](https://github.com/casey/just)
3. [Docker Desktop](https://www.docker.com/products/docker-desktop/)
4. [Terraform](https://developer.hashicorp.com/terraform/tutorials/aws-get-started/install-cli)
5. [Localstack](https://docs.localstack.cloud/user-guide/aws/feature-coverage/)

## Resources

- [LocalStack API Gateway docs](https://docs.localstack.cloud/user-guide/aws/apigateway/)
- [LocalStack Lambda docs](https://docs.localstack.cloud/user-guide/aws/lambda/)
- [Test AWS Lambdas on localhost](https://prabhakar-borah.medium.com/localstack-test-your-lambda-on-your-localhost-5cce066c967c)

## Local deploy and lambda access

- `just deploy-local`
- Take the _api_gw_id_ from the terraform output when the deploy ends
- `curl -X GET http://[API-GW-ID].execute-api.localhost.localstack.cloud:4566/[SOME-STRING]/hello`
  - Example: `curl -X GET http://yt5foqe749.execute-api.localhost.localstack.cloud:4566/lambda/hello`
- `curl -X POST -H 'Content-Type: application/json' -d '{"name":"Antonio"}' http://[API-GW-ID].execute-api.localhost.localstack.cloud:4566/[SOME-STRING]/say-my-name`
  - Example: `curl -X POST -H 'Content-Type: application/json' -d '{"name":"Antonio"}' http://yt5foqe749.execute-api.localhost.localstack.cloud:4566/lambda/say-my-name`
