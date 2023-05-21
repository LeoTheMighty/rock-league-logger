
resource "aws_api_gateway_rest_api" "rock_league_logger_api" {
  name        = "RockLeagueLoggerAPI"
  description = "This is the API Gateway for the logger microservice of Rock League"
}

resource "aws_api_gateway_resource" "rock_league_logger_resource" {
  rest_api_id = aws_api_gateway_rest_api.rock_league_logger_api.id
  parent_id   = aws_api_gateway_rest_api.rock_league_logger_api.root_resource_id
  path_part   = "logs"
}

resource "aws_api_gateway_method" "rock_league_logger_get_method" {
  rest_api_id   = aws_api_gateway_rest_api.rock_league_logger_api.id
  resource_id   = aws_api_gateway_resource.rock_league_logger_resource.id
  http_method   = "GET"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "rock_league_logger_get_integration" {
  rest_api_id   = aws_api_gateway_rest_api.rock_league_logger_api.id
  resource_id   = aws_api_gateway_resource.rock_league_logger_resource.id
  http_method = aws_api_gateway_method.rock_league_logger_get_method.http_method

  integration_http_method = "GET"
  type        = "AWS_PROXY"
  uri         = aws_lambda_function.rock_league_logger.invoke_arn
}

resource "aws_api_gateway_method" "rock_league_logger_post_method" {
  rest_api_id   = aws_api_gateway_rest_api.rock_league_logger_api.id
  resource_id   = aws_api_gateway_resource.rock_league_logger_resource.id
  http_method   = "POST"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "rock_league_logger_post_integration" {
  rest_api_id   = aws_api_gateway_rest_api.rock_league_logger_api.id
  resource_id   = aws_api_gateway_resource.rock_league_logger_resource.id
  http_method = aws_api_gateway_method.rock_league_logger_post_method.http_method

  integration_http_method = "POST"
  type        = "AWS_PROXY"
  uri         = aws_lambda_function.rock_league_logger.invoke_arn
}

resource "aws_api_gateway_method" "rock_league_logger_patch_method" {
  rest_api_id   = aws_api_gateway_rest_api.rock_league_logger_api.id
  resource_id   = aws_api_gateway_resource.rock_league_logger_resource.id
  http_method   = "PATCH"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "rock_league_logger_patch_integration" {
  rest_api_id   = aws_api_gateway_rest_api.rock_league_logger_api.id
  resource_id   = aws_api_gateway_resource.rock_league_logger_resource.id
  http_method = aws_api_gateway_method.rock_league_logger_patch_method.http_method

  integration_http_method = "PATCH"
  type        = "AWS_PROXY"
  uri         = aws_lambda_function.rock_league_logger.invoke_arn
}

resource "aws_api_gateway_method" "rock_league_logger_delete_method" {
  rest_api_id   = aws_api_gateway_rest_api.rock_league_logger_api.id
  resource_id   = aws_api_gateway_resource.rock_league_logger_resource.id
  http_method   = "DELETE"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "rock_league_logger_delete_integration" {
  rest_api_id   = aws_api_gateway_rest_api.rock_league_logger_api.id
  resource_id   = aws_api_gateway_resource.rock_league_logger_resource.id
  http_method = aws_api_gateway_method.rock_league_logger_delete_method.http_method

  integration_http_method = "DELETE"
  type        = "AWS_PROXY"
  uri         = aws_lambda_function.rock_league_logger.invoke_arn
}

resource "aws_api_gateway_deployment" "rock_league_logger_gateway_deployment" {
  depends_on = [
    aws_api_gateway_integration.rock_league_logger_get_integration,
    aws_api_gateway_integration.rock_league_logger_post_integration,
    aws_api_gateway_integration.rock_league_logger_patch_integration,
    aws_api_gateway_integration.rock_league_logger_delete_integration,
  ]

  rest_api_id = aws_api_gateway_rest_api.rock_league_logger_api.id
  stage_name  = "v1"
}
