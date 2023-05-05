package main

import (
	search "food-votes/api/search"
	structs "food-votes/structs"

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

	switch path := event.RawPath; path {
	case "/search":
		body := search.Request(event.QueryStringParameters.Zipcode)
		fmt.Println("Returning a Response with ", len(body.Results), " results.")
		return body, nil // TODO Return Room ID as well
	case "/join":
		// Unimplemented
		fmt.Println("Unimplemented. Will join an existing room id with existing votes for each location.")
		return structs.Response{}, nil
	case "/vote":
		// Unimplemented
		fmt.Println("Unimplemented. Will add a vote to a location in a room id.")
		return structs.Response{}, nil
	default:
		fmt.Printf("%s.\n", path)
		return structs.Response{}, nil
	}

}

func main() {
	// Lambda function invocation
	lambda.Start(HandleLambdaEvent)
}
