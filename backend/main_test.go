package main

import (
	"food-votes/structs"
	"testing"
)

func TestE2E(t *testing.T) {
	// ! TODO - Split E2E into 3 Tests
	var sharedRoomID string

	// Search and Create Room
	createdRoom, _ := HandleLambdaEvent(structs.Event{
		RawPath: "/search",
		QueryStringParameters: struct {
			Query string `json:"query"`
		}{"20001"},
	})

	// Check if body.ID is a UUID
	if len(createdRoom.ID) != 8 {
		t.Errorf("Expected a partial UUID (8 characters), got %s.", createdRoom.ID)
	} else {
		sharedRoomID = createdRoom.ID
	}

	if len(createdRoom.Options) == 0 {
		t.Errorf("Expected at least 1 result, got %d.", len(createdRoom.Options))
	}

	// Join and Get Existing Room
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

	// Vote and Validate Existing Room
	knownRestaurant := joinedRoom.Options[0].Name

	votedRoom, _ := HandleLambdaEvent(structs.Event{
		RawPath: "/vote",
		QueryStringParameters: struct {
			Query string `json:"query"`
		}{sharedRoomID + ":" + knownRestaurant},
	})

	if votedRoom.ID != sharedRoomID {
		t.Errorf("Expected %s, got %s.", sharedRoomID, votedRoom.ID)
	}

	if len(votedRoom.Options) == 0 {
		t.Errorf("Expected at least 1 result, got %d.", len(votedRoom.Options))
	}

	// Should only be 1 for knownRestaurant
	for _, option := range votedRoom.Options {
		if option.Name == knownRestaurant {
			if option.Votes != 1 {
				t.Errorf("Expected %s to have 1 vote, got %d.", knownRestaurant, option.Votes)
			}
		} else {
			if option.Votes != 0 {
				t.Errorf("Expected %s to have 0 votes, got %d.", option.Name, option.Votes)
			}
		}
	}
}
