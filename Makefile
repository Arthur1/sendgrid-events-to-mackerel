build/sendgrid-webhook-receiver-lambda/bootstrap:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o ./build/sendgrid-webhook-receiver-lambda/bootstrap ./cmd/sendgrid-webhook-receiver-lambda

build/sendgrid-webhook-receiver-lambda/lambda.zip: build/sendgrid-webhook-receiver-lambda/bootstrap
	cd build/sendgrid-webhook-receiver-lambda; zip lambda.zip bootstrap

.PHONY: clean
clean:
	$(RM) -r build/
