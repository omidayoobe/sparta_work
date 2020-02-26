locals {
  name      = "health-dashboard-notifier"
  full_name = "${var.env}-${local.name}"
  tags = {
    Name        = local.full_name
    Environment = var.env
  }
}

#------------------------------------------------------------------------------
# CloudWatch Event
#------------------------------------------------------------------------------
resource "aws_cloudwatch_event_rule" "trigger_event" {
  name          = local.full_name
  description   = "captures health dashboard events"
  event_pattern = jsonencode({ "source" = ["aws.health"] })
  tags          = local.tags
}

resource "aws_cloudwatch_event_target" "target" {
  rule      = aws_cloudwatch_event_rule.trigger_event.name
  target_id = aws_lambda_function.health_dashboard_notifier_lambda.function_name
  arn       = aws_lambda_function.health_dashboard_notifier_lambda.arn
}