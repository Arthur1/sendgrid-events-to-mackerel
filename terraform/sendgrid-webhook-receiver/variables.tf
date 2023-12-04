variable "region" {
  description = "AWS region"
  type        = string
  default     = "ap-northeast-1"
}

variable "name" {
  description = "Name of resources"
  type        = string
  default     = "sendgrid-webhook-receiver"
}

variable "tags" {
  description = "Tags of resources"
  type        = map(string)
  default     = {}
}
