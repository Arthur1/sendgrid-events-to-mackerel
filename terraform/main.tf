locals {
  region = "ap-northeast-1"
}

module "sendgrid_webhook_receiver" {
  source = "./sendgrid-webhook-receiver"
  region = local.region
}

module "sendgrid_webhook_logs_aggregator" {
  source                = "./sendgrid-webhook-logs-aggregator"
  region                = local.region
  target_log_group_name = module.sendgrid_webhook_receiver.lambda_log_group_name
  # 以下の値は利用者の環境に合わせて設定すること
  mackerel_service_name = "SendGrid"
  mackerel_api_key_name = "/mackerel.io/arthur-1-test/apikey"
}
