package main

import (
	room_mgr "food-votes/api/roommanager"
	search "food-votes/api/search"
	structs "food-votes/structs"

	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
)

func HandleLambdaEvent(event structs.Event) (structs.Room, error) {
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
		body := search.Request(event.QueryStringParameters.Query)

		options := []structs.Option{}

		// Loop through Results and add them to Options
		for _, result := range body.Results {
			options = append(options, structs.Option{
				Name:  result.Name,
				Votes: 0,
			})
		}

		// Create a Room with the Options
		room := room_mgr.Create(options)
		fmt.Println("Returning a Response with ", len(room.Options), " results.")
		return room, nil
	case "/join":
		ID := event.QueryStringParameters.Query
		room := room_mgr.Join(ID)
		fmt.Println("Returning a Response with ", len(room.Options), " results.")
		return room, nil
	case "/vote":
		// Unimplemented
		fmt.Println("Unimplemented. Will add a vote to a location in a room id.")
		return structs.Room{}, nil
	default:
		fmt.Printf("%s.\n", path)
		return structs.Room{}, nil
	}

}

func main() {
	// Lambda function invocation
	lambda.Start(HandleLambdaEvent)
}
