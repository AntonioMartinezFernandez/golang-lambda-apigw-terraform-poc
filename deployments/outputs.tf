output "rest_api_gateway_base_url" {
  value = "http://${aws_api_gateway_rest_api.rest_api_gateway.id}.execute-api.localhost.localstack.cloud:4566/lambda"
}
