locals {
  rest_api_golang_lambda_env_vars = {
    # Common Tags
    APP_SERVICE_NAME = "rest_api_golang_lambda"
    APP_ENV          = "serverless"
    APP_VERSION      = "1.0.0"

    # App Configuration
    LOG_LEVEL             = "debug"
    JSON_SCHEMA_BASE_PATH = "./schemas/"

    # AWS Parameters
    AWS_REGION = "eu-central-1"

    # DynamoDB
    DYNAMO_DB_ENDPOINT = "http://localstack:4566"
  }
}
