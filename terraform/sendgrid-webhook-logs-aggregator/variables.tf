variable "region" {
  description = "AWS region"
  type        = string
  default     = "ap-northeast-1"
}

variable "mackerel_service_name" {
  description = "Mackerel Service name to post metrics"
  type        = string
}

variable "mackerel_api_key_name" {
  description = "Name of Parameter which stores Mackerel API Key"
  type        = string
}

variable "name" {
  description = "Name of resources"
  type        = string
  default     = "sendgrid-webhook-logs-aggregator"
}

variable "target_log_group_name" {
  description = "target log group name"
  type        = string
}

variable "tags" {
  description = "Tags of resources"
  type        = map(string)
  default     = {}
}
