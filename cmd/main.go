package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/matiasnoriega-uala/poc-module-access-list-lambda/pkg/handler"
)

func main() {
	handler := handler.ModuleAccessListHandler{}
	lambda.Start(handler.HandleRequest)
}
