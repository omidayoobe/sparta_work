#------------------------------------------------------------------------------
# Lambda Role
#------------------------------------------------------------------------------
data "aws_iam_policy_document" "role_policy" {
  statement {
    actions = [
      "sts:AssumeRole",
    ]
    principals {
      type = "Service"
      identifiers = [
        "lambda.amazonaws.com",
      ]
    }
  }
}

resource "aws_iam_role" "lambda_role" {
  name               = local.name
  assume_role_policy = data.aws_iam_policy_document.role_policy.json
}

resource "aws_iam_role_policy_attachment" "basic_execution" {
  role       = aws_iam_role.lambda_role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

#------------------------------------------------------------------------------
# Lambda
#------------------------------------------------------------------------------
data "archive_file" "lambda" {
  source_file = "../${local.name}.zip"
  output_path = "${path.module}/${local.name}.zip"
  type        = "zip"
}

resource "aws_lambda_function" "health_dashboard_notifier_lambda" {
  function_name = local.name
  handler       = local.name
  role          = aws_iam_role.lambda_role.arn
  runtime       = "go1.x"

  memory_size = var.memory_size
  timeout     = var.timeout

  filename         = data.archive_file.lambda.output_path
  source_code_hash = data.archive_file.lambda.output_base64sha256

  environment {
    variables = {
      SLACK_WEBHOOK = var.slack_warning
    }
  }

}

resource "aws_lambda_permission" "allow_cloudwatch_execute_lambda" {
  statement_id  = "AllowExecutionFromCloudWatch"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.health_dashboard_notifier_lambda.function_name
  principal     = "events.amazonaws.com"
  source_arn    = aws_cloudwatch_event_rule.trigger_event.arn
}