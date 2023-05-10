package main

import (
	"food-votes/structs"
	"testing"
)

func TestSearchAndJoin(t *testing.T) {
	var sharedRoomID string

	createdRoom, _ := HandleLambdaEvent(structs.Event{
		RawPath: "/search",
		QueryStringParameters: struct {
			Query string `json:"query"`
		}{"20001"},
	})

	// Check if body.ID is a UUID
	if len(createdRoom.ID) != 36 {
		t.Errorf("Expected a UUID, got %s.", createdRoom.ID)
	} else {
		sharedRoomID = createdRoom.ID
	}

	if len(createdRoom.Options) == 0 {
		t.Errorf("Expected at least 1 result, got %d.", len(createdRoom.Options))
	}

	// TODO dependent on
	joinedRoom, _ := HandleLambdaEvent(structs.Event{
		RawPath: "/join",
		QueryStringParameters: struct {
			Query string `json:"query"`
		}{sharedRoomID},
	})

	if joinedRoom.ID != sharedRoomID {
		t.Errorf("Expected %s, got %s.", sharedRoomID, joinedRoom.ID)
	}

	if len(joinedRoom.Options) == 0 {
		t.Errorf("Expected at least 1 result, got %d.", len(joinedRoom.Options))
	}
}
