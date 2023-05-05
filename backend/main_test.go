package main

import (
	"food-votes/structs"
	"testing"
)

func TestMain(t *testing.T) {
	body, _ := HandleLambdaEvent(structs.Event{
		RawPath: "/search",
		QueryStringParameters: struct {
			Zipcode string `json:"zipcode"`
		}{"20001"},
	})

	if len(body.Results) == 0 {
		t.Errorf("Expected at least 1 result, got %d.", len(body.Results))
	}
}
