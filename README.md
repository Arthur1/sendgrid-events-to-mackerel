# sendgrid-events-to-mackerel

![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/Arthur1/sendgrid-events-to-mackerel)

> [!WARNING]
> これは実験的な実装です。本番環境での利用は各自の責任でお願いします。

このリポジトリは以下の2つのアプリケーションを含みます：

- SendGrid Event Webhookのリクエストを受けて、Delivery EventをパースするHTTPサーバ
- 出力したログを[cloudwatch-logs-aggregator](https://github.com/mackerelio-labs/mackerel-monitoring-modules/tree/main/cloudwatch-logs-aggregator#readme)でサービスメトリック化してMackerelに投稿する

```console
$ make build/sendgrid-webhook-receiver-lambda/lambda.zip
$ cd terraform
$ terraform init
$ terraform apply
```
