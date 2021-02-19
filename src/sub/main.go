package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "redis:6379", // We connect to host redis, thats what the hostname of the redis service is set to in the docker-compose
		Password: "secret",     // The password IF set in the redis Config file
		DB:       0,
	})
	ctx := context.Background()
	// Subscribe to the Topic given
	topic := redisClient.Subscribe(ctx, "new_users")
	// Get the Channel to use
	channel := topic.Channel()
	// Iterate any messages sent on the channel
	for msg := range channel {
		u := &User{}
		// Unmarshal the data into the user
		err := u.UnmarshalBinary([]byte(msg.Payload))
		if err != nil {
			panic(err)
		}
		fmt.Println(u)
	}
}
