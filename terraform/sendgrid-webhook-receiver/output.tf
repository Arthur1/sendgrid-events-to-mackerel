output "lambda_log_group_name" {
  value = resource.aws_cloudwatch_log_group.server.name
}
