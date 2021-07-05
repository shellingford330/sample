package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest() (string, error) {
	log.Printf("Accept access.")
	return "Hello!!!", nil
}

func main() {
	lambda.Start(HandleRequest)
}
