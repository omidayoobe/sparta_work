locals {
  name = "health-dashboard-notifier"
}

#------------------------------------------------------------------------------
# CloudWatch Event
#------------------------------------------------------------------------------
resource "aws_cloudwatch_event_rule" "trigger_event" {
  name          = local.name
  description   = "captures health dashboard events"
  event_pattern = jsonencode({ "source" = ["aws.health"] })
}

resource "aws_cloudwatch_event_target" "target" {
  rule      = aws_cloudwatch_event_rule.trigger_event.name
  target_id = aws_lambda_function.health_dashboard_notifier_lambda.function_name
  arn       = aws_lambda_function.health_dashboard_notifier_lambda.arn
}