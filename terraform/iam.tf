data "aws_iam_policy_document" "assume_lambda" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
  }
}

resource "aws_iam_role" "assume_lambda_role" {
  name = "AssumeLambdaRole"
  description = "Role for lambda to assume lambda"
  assume_role_policy = data.aws_iam_policy_document.assume_lambda.json
}

data "aws_iam_policy_document" "allow_lambda_logging" {
  statement {
    effect = "Allow"
    actions = [
      "logs:CreateLogStream",
      "logs:PutLogEvents",
    ]

    resources = [
      "arn:aws:logs:*:*:*",
    ]
  }
}

resource "aws_iam_policy" "lambda_logging_policy" {
  name        = "AllowLambdaLoggingPolicy"
  description = "Policy for Lambda Cloudwatch Logging"
  policy      = data.aws_iam_policy_document.allow_lambda_logging.json
}

resource "aws_iam_role_policy_attachment" "lambda_logging_policy_attachment" {
  role       = aws_iam_role.assume_lambda_role.id
  policy_arn = aws_iam_policy.lambda_logging_policy.arn
}

data "aws_iam_policy_document" "allow_dynamodb_operations" {
  statement {
    effect = "Allow"
    actions = [
      "dynamodb:PutItem",
      "dynamodb:GetItem",
      "dynamodb:UpdateItem",
      "dynamodb:DeleteItem",
      "dynamodb:Scan",
      "dynamodb:Query"
    ]

    resources = [aws_dynamodb_table.logs.arn]
  }
}

resource "aws_iam_policy" "dynamodb_operations_policy" {
  name        = "AllowDynamoDBOperationsPolicy"
  description = "Policy for DynamoDB operations"
  policy      = data.aws_iam_policy_document.allow_dynamodb_operations.json
}

resource "aws_iam_role_policy_attachment" "dynamodb_operations_policy_attachment" {
  role       = aws_iam_role.assume_lambda_role.id
  policy_arn = aws_iam_policy.dynamodb_operations_policy.arn
}
