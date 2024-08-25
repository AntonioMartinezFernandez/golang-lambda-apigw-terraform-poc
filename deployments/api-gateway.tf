resource "aws_api_gateway_rest_api" "rest_api_gateway" {
  name        = "api"
  description = "API Gateway V1"
}

resource "aws_api_gateway_resource" "proxy" {
  rest_api_id = aws_api_gateway_rest_api.rest_api_gateway.id
  parent_id   = aws_api_gateway_rest_api.rest_api_gateway.root_resource_id
  path_part   = "{proxy+}" // The path_part 'proxy+' value activate proxy behaviour, so it willl match any request path
}

resource "aws_api_gateway_method" "proxy" {
  rest_api_id   = aws_api_gateway_rest_api.rest_api_gateway.id
  resource_id   = aws_api_gateway_resource.proxy.id
  http_method   = "ANY"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "lambda" {
  rest_api_id             = aws_api_gateway_rest_api.rest_api_gateway.id
  resource_id             = aws_api_gateway_method.proxy.resource_id
  http_method             = aws_api_gateway_method.proxy.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.rest_api_golang_lambda.invoke_arn
}

resource "aws_api_gateway_deployment" "apigw_deployment" {
  depends_on = [
    aws_api_gateway_integration.lambda,
  ]
  rest_api_id = aws_api_gateway_rest_api.rest_api_gateway.id
  stage_name  = "test"
}
