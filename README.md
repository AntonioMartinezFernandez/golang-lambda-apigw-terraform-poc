# golang-lambda-apigw-terraform-poc

Running and testing AWS Lambda (Golang), API Gateway and DynamoDB in local environment using Terraform and LocalStack

## Requirements

1. [Golang](https://go.dev/dl/)
2. [Docker Desktop](https://www.docker.com/products/docker-desktop/)
3. [Terraform](https://developer.hashicorp.com/terraform/tutorials/aws-get-started/install-cli)
4. [Localstack](https://docs.localstack.cloud/user-guide/aws/feature-coverage/)
5. [AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html)
6. [Just](https://github.com/casey/just)

## Getting Started

Clone the repo:

```
git clone https://github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc && cd golang-lambda-apigw-terraform-poc
```

Start Serverless Mode (it will provide an output with the lambda URL):

```
# You need to have Docker started and then execute the next command
just start-serverless
```

Run unit tests:

```
just test-unit
```

Run integration tests:

```
# You need to start serverless mode first in order to have deployed all the necessary components and then execute the next command
just test-integration
```

Start Independent Service Mode:

```
# You need to start serverless mode first in order to have deployed all the necessary components and then execute the next command
just start-local
```

ℹ️ The OpenAPI definition for this project is available [here](https://app.swaggerhub.com/apis/ANTONIO_22/golang-lambda-apigw-terraform-poc/1.0.0) and is also available in the `api\swagger.yaml` file of this repository.

## Resources

- [LocalStack API Gateway docs](https://docs.localstack.cloud/user-guide/aws/apigateway/)
- [LocalStack Lambda docs](https://docs.localstack.cloud/user-guide/aws/lambda/)
- [Locally debug functions with AWS SAM](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-using-debugging.html)
- [Test AWS Lambdas on localhost](https://prabhakar-borah.medium.com/localstack-test-your-lambda-on-your-localhost-5cce066c967c)
- [Testing and Running Go API GW Lambdas Locally](https://boyter.org/posts/testing-running-api-gw-lambda-locally/)
- [Serverless Applications with AWS Lambda and API Gateway](https://registry.terraform.io/providers/hashicorp/aws/2.34.0/docs/guides/serverless-with-aws-lambda-and-api-gateway)
- [DynamoDB Developer Guide](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/SQLtoNoSQL.html)
- [DynamoDB examples using SDK for Go V2](https://github.com/awsdocs/aws-doc-sdk-examples/tree/main/gov2/dynamodb/actions)
- [DynamoDB Table Terraform module](https://registry.terraform.io/modules/terraform-aws-modules/dynamodb-table/aws/latest)
- [DynamoDB Hash Key (Partition Key -PK-) and Range Key (Sort Key -SK-)](https://stackoverflow.com/questions/27329461/what-is-hash-and-range-primary-key)
- [Best Practices in API Design](https://swagger.io/resources/articles/best-practices-in-api-design/)
