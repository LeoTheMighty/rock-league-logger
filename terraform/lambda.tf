data "archive_file" "lambda_file" {
  type        = "zip"
  source_file = "${path.module}/../lambda/main"
  output_path = "${path.module}/../lambda/lambda.zip"
}

resource "aws_lambda_function" "rock_league_logger" {
  function_name    = "RockLeagueLogger"
  filename          = "${path.module}/../lambda/lambda.zip"
  handler          = "main"  # assuming you're exporting a function named "main" in your Go code
  runtime          = "go1.x"
  role             = aws_iam_role.assume_lambda_role.arn

  source_code_hash = data.archive_file.lambda_file.output_base64sha256

  environment {
    variables = {
      TABLE_NAME = aws_dynamodb_table.logs.name
    }
  }

  tags = {
    Name = "RockLeagueLogger"
  }
}