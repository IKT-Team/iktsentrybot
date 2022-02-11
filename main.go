package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type LambdaEvent struct {
	Body string `json:"body"`
}

type Response struct {
	StatusCode int `json:"statusCode"`
}

func main() {
	lambda.Start(handler)
}

func handler(lambdaEvent LambdaEvent) (Response, error) {
	fmt.Println(lambdaEvent.Body)
	sentryEvent, err := DecodeSentryEvent(lambdaEvent.Body)
	if err != nil {
		return Response{StatusCode: 500}, err
	}
	err = SendTelegramMessage(sentryEvent)
	if err != nil {
		return Response{StatusCode: 500}, err
	}
	return Response{StatusCode: 200}, nil
}
