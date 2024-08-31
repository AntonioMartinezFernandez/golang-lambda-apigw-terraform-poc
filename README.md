# golang-lambda-apigw-terraform-poc

Deploying AWS Lambda (Golang) and API Gateway in local environment with Terraform and LocalStack

## Requirements

1. [Golang](https://go.dev/dl/)
2. [Just](https://github.com/casey/just)
3. [Docker Desktop](https://www.docker.com/products/docker-desktop/)
4. [Terraform](https://developer.hashicorp.com/terraform/tutorials/aws-get-started/install-cli)
5. [Localstack](https://docs.localstack.cloud/user-guide/aws/feature-coverage/)
6. [AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html)

## Getting Started

Clone the repo:

```
git clone https://github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc && cd golang-lambda-apigw-terraform-poc
```

Start Serverless Mode (it will provide an output with the lambda URL):

```
# You need to have Docker started
just start-serverless
```

Start Independent Service Mode:

```
# Start serverless mode first to deploy the necessary infrastructure
just start-local
```

## Checking Serverless Mode

- Start serverless mode
- Take the `rest_api_gateway_base_url` value from terraform output when the deploy ends
- Execute `curl -X GET [rest_api_gateway_base_url]/healthcheck`
  - Example: `curl -X GET http://ia58t50h0q.execute-api.localhost.localstack.cloud:4566/lambda/healthcheck`
- Execute `curl -X POST -H 'Content-Type: application/json' -d '{"id":"01J63630X372YYYR4CTFP1ZGGZ","name":"Antonio"}' [rest_api_gateway_base_url]/user`
  - Example: `curl -X POST -H 'Content-Type: application/json' -d '{"id":"01J63630X372YYYR4CTFP1ZGGZ","name":"Antonio", "birthdate":"1984-11-25 17:04:12"}' http://ia58t50h0q.execute-api.localhost.localstack.cloud:4566/lambda/user`
- Execute `curl -X GET [rest_api_gateway_base_url]/user/[USER_ID]`
  - Example: `curl -X GET http://ia58t50h0q.execute-api.localhost.localstack.cloud:4566/lambda/user/01J63630X372YYYR4CTFP1ZGGZ`

## Resources

- [LocalStack API Gateway docs](https://docs.localstack.cloud/user-guide/aws/apigateway/)
- [LocalStack Lambda docs](https://docs.localstack.cloud/user-guide/aws/lambda/)
- [Locally debug functions with AWS SAM](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-using-debugging.html)
- [Test AWS Lambdas on localhost](https://prabhakar-borah.medium.com/localstack-test-your-lambda-on-your-localhost-5cce066c967c)
- [Testing and Running Go API GW Lambdas Locally](https://boyter.org/posts/testing-running-api-gw-lambda-locally/)
- [Serverless Applications with AWS Lambda and API Gateway](https://registry.terraform.io/providers/hashicorp/aws/2.34.0/docs/guides/serverless-with-aws-lambda-and-api-gateway)
- [DynamoDB Developer Guide](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/SQLtoNoSQL.html)
- [DynamoDB examples using SDK for Go V2](https://docs.aws.amazon.com/code-library/latest/ug/go_2_dynamodb_code_examples.html)
- [DynamoDB Table Terraform module](https://registry.terraform.io/modules/terraform-aws-modules/dynamodb-table/aws/latest)
- [DynamoDB Hash Key (Partition Key -PK-) and Range Key (Sort Key -SK-)](https://stackoverflow.com/questions/27329461/what-is-hash-and-range-primary-key)
