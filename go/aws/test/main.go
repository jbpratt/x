package main

import "github.com/aws/aws-lambda-go/lambda"

func main() {
	lambda.Start(handler)
}

func handler() (string, error) {
	return "Welcome to Serverless world", nil
}
