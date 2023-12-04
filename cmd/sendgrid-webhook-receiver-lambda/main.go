package main

import (
	"github.com/Arthur1/sendgrid-events-to-mackerel/internal/server"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
)

func main() {
	h := server.NewHTTPHandler()
	lambda.Start(httpadapter.NewV2(h).ProxyWithContext)
}
