package main

import (
	"food-votes/api/getLocal"
	"food-votes/structs"

	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
)

func HandleLambdaEvent(event structs.Event) (structs.Response, error) {
	fmt.Println("Event received: ", event)

	// Load environment variable from .env if not set
	if os.Getenv("FOURSQUARE_API_KEY") == "" {
		println("Environment not set, checking .env.")
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
			println(err)
		}
	} else {
		println("Environment set.")
	}

	if event.RawPath == "/getLocal" {
		body := getLocal.Request(event.QueryStringParameters.Zipcode)
		fmt.Println("Returning a Response with ", len(body.Results), " results.")
		return body, nil
	}

	return structs.Response{}, nil
}

func main() {
	// Lambda function invocation
	lambda.Start(HandleLambdaEvent)
}
