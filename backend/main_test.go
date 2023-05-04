package main

import (
	"food-votes/structs"
	"testing"
)

func TestMain(t *testing.T) {
	body, _ := HandleLambdaEvent(structs.Event{
		RawPath: "/getLocal",
		Headers: struct {
			XForwardedFor string `json:"x-forwarded-for"`
			Origin        string `json:"origin"`
		}{"abc", "def"},
		QueryStringParameters: struct {
			Zipcode string `json:"zipcode"`
		}{"20148"},
	})

	if len(body.Results) == 0 {
		t.Errorf("Expected at least 1 result, got %d.", len(body.Results))
	}
}
