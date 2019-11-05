package main

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandler(t *testing.T) {

	params := events.APIGatewayProxyRequest{
		PathParameters: map[string]string{
			"movie": "Titanic",
		},
	}
	_, err := handler(params)

	if err != nil {
		t.Error("Shouldn't be null")
	}
}
