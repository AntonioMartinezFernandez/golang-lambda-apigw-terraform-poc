resource "aws_lambda_function" "rest_api_golang_lambda" {
  filename      = "${path.module}/../build/app.zip"
  function_name = "rest_api_golang_lambda"
  role          = aws_iam_role.lambda_role.arn
  handler       = "app"
  runtime       = "go1.x"
  timeout       = 30

  environment {
    variables = {
      foo = "bar"
    }
  }
}

resource "aws_lambda_permission" "apigw_lambda" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.rest_api_golang_lambda.function_name
  principal     = "apigateway.amazonaws.com"

  source_arn = "${aws_api_gateway_rest_api.rest_api_gateway.execution_arn}/*/*/*"
}
