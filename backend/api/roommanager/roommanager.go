package roommanager

import (
	"encoding/json"
	"os"
	"time"

	"food-votes/structs"

	"github.com/go-redis/redis"
	"github.com/google/uuid"
)

func Create(options []structs.Option) structs.Room {

	URL := os.Getenv("REDIS_URL")
	PW := os.Getenv("REDIS_PW")

	// Connect to Redis
	client := redis.NewClient(&redis.Options{
		Addr:     URL,
		Password: PW,
		DB:       0, // use default DB
	})

	// Random ID for the room
	ID := uuid.NewString()[0:8]
	room := structs.Room{
		ID:      ID,
		Options: options,
	}

	// Convert the Room object to JSON
	jsonRoom, err := json.Marshal(room)
	if err != nil {
		panic(err)
	}

	// Set the key-value pair with a 1-hour expiration
	err = client.Set(ID, jsonRoom, 60*60*time.Second).Err()
	if err != nil {
		panic(err)
	}

	// Get the value of the key
	val, err := client.Get(ID).Bytes()
	if err != nil {
		panic(err)
	}
	// Convert the JSON back to a Room object
	var newRoom structs.Room
	err = json.Unmarshal(val, &newRoom)
	if err != nil {
		panic(err)
	}

	// Close the connection when done
	err = client.Close()
	if err != nil {
		panic(err)
	}

	return newRoom

}

func Join(ID string) structs.Room {

	URL := os.Getenv("REDIS_URL")
	PW := os.Getenv("REDIS_PW")

	// Connect to Redis
	client := redis.NewClient(&redis.Options{
		Addr:     URL,
		Password: PW,
		DB:       0, // use default DB
	})

	// Get the value of the key
	val, err := client.Get(ID).Bytes()
	if err != nil {
		panic(err)
	}
	// Convert the JSON back to a Room object
	var newRoom structs.Room
	err = json.Unmarshal(val, &newRoom)
	if err != nil {
		panic(err)
	}

	// Close the connection when done
	err = client.Close()
	if err != nil {
		panic(err)
	}

	return newRoom
}

func Update(ID string, room structs.Room) structs.Room {

	URL := os.Getenv("REDIS_URL")
	PW := os.Getenv("REDIS_PW")

	// Connect to Redis
	client := redis.NewClient(&redis.Options{
		Addr:     URL,
		Password: PW,
		DB:       0, // use default DB
	})

	// Convert the Room object to JSON
	jsonRoom, err := json.Marshal(room)
	if err != nil {
		panic(err)
	}

	// Set the key-value pair with a 1-hour expiration
	err = client.Set(ID, jsonRoom, 60*60*time.Second).Err()
	if err != nil {
		panic(err)
	}

	return room
}

func Vote(ID string, LocationName string) structs.Room {

	var room structs.Room = Join(ID)

	for i, option := range room.Options {
		if option.Name == LocationName {
			room.Options[i].Votes++
		}
	}

	room = Update(ID, room)

	return room
}
