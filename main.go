package main

import (
	"encoding/json"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mlabouardy/imdb"
)

func format(payload interface{}) string {
	response, _ := json.Marshal(payload)
	return string(response)
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	name := request.PathParameters["movie"]

	response, err := imdbClient.MovieByTitle(name, "")
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Headers: map[string]string{
				"Content-Type":                 "application/json",
				"Access-Control-Allow-Origin":  "*",
				"Access-Control-Allow-Headers": "*",
			},
			Body: "Couldn't find the movie",
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type":                 "application/json",
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Headers": "*",
		},
		Body: format(response),
	}, nil
}

var imdbClient *imdb.Client

func init() {
	imdbClient = imdb.New(os.Getenv("API_KEY"))
}

func main() {
	lambda.Start(handler)
}
