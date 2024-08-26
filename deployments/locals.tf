locals {
  rest_api_golang_lambda_env_vars = {
    # Common Tags
    APP_SERVICE_NAME = "rest_api_golang_lambda"
    APP_ENV          = "serverless"
    APP_VERSION      = "1.0.0"

    # App Configuration
    LOG_LEVEL             = "debug"
    JSON_SCHEMA_BASE_PATH = "./schemas/"
  }
}
