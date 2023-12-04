locals {
  metric_name_prefix = replace(var.name, "-", "_")
}

module "cw_logs_aggregator_lambda" {
  source = "github.com/mackerelio-labs/mackerel-monitoring-modules//cloudwatch-logs-aggregator/lambda?ref=v0.2.1"

  region        = var.region
  tags          = var.tags
  function_name = var.name
  iam_role_name = "${var.name}-lambda"
}

module "cw_logs_aggregator_rule_delivery_events" {
  source = "github.com/mackerelio-labs/mackerel-monitoring-modules//cloudwatch-logs-aggregator/rule?ref=v0.2.1"

  region       = var.region
  rule_name    = "${var.name}-delivery-events"
  function_arn = module.cw_logs_aggregator_lambda.function_arn

  api_key_name = var.mackerel_api_key_name
  service_name = var.mackerel_service_name

  log_group_name     = var.target_log_group_name
  query              = <<-EOT
    filter level = "INFO" and msg = "sendgrid delivery events count"
    | stats sum(processedCount) as `~processed`, sum(droppedCount) as `~dropped`, sum(deliveredCount) as `~delivered`, sum(deferredCount) as `~deferred`, sum(bounceCount) as `~bounce`
  EOT
  metric_name_prefix = "${local.metric_name_prefix}.delivery_events"
  default_metrics = {
    "${local.metric_name_prefix}.delivery_events.processed" = 0
    "${local.metric_name_prefix}.delivery_events.dropped"   = 0
    "${local.metric_name_prefix}.delivery_events.delivered" = 0
    "${local.metric_name_prefix}.delivery_events.deferred"  = 0
    "${local.metric_name_prefix}.delivery_events.bounce"    = 0
  }
  schedule_expression = "rate(1 minute)"
  interval_in_minutes = 1
  offset_in_minutes   = 5
}
