data "aws_iam_policy_document" "assume_role" {
  statement {
    effect = "Allow"
    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
    actions = ["sts:AssumeRole"]
  }
}

resource "aws_iam_role" "server" {
  name               = "${var.name}-lambda-role"
  assume_role_policy = data.aws_iam_policy_document.assume_role.json
  tags               = var.tags
}

data "aws_iam_policy" "lambda_basic_execution" {
  arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

resource "aws_iam_role_policy_attachment" "this" {
  role       = aws_iam_role.server.name
  policy_arn = data.aws_iam_policy.lambda_basic_execution.arn
}

resource "aws_lambda_function" "server" {
  function_name    = var.name
  runtime          = "provided.al2023"
  handler          = "main"
  role             = aws_iam_role.server.arn
  architectures    = ["arm64"]
  memory_size      = 256
  timeout          = 10
  package_type     = "Zip"
  filename         = "${path.root}/../build/sendgrid-webhook-receiver-lambda/lambda.zip"
  source_code_hash = filebase64sha256("${path.root}/../build/sendgrid-webhook-receiver-lambda/lambda.zip")
  tags             = var.tags
}

resource "aws_lambda_permission" "api" {
  statement_id  = "AllowAPIGatewayInvoke"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.server.function_name
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${aws_apigatewayv2_api.api.execution_arn}/*/*"
}

resource "aws_cloudwatch_log_group" "server" {
  name              = "/aws/lambda/${resource.aws_lambda_function.server.function_name}"
  retention_in_days = 7
  tags              = var.tags
}
