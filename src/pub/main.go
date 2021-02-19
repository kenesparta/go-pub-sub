package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	// Create a new Redis Client
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "redis:6379", // We connect to host redis, thats what the hostname of the redis service is set to in the docker-compose
		Password: "secret",     // The password IF set in the redis Config file
		DB:       0,
	})
	// Generate a new background context that  we will use
	ctx := context.Background()
	// Loop and randomly generate users on a random timer
	for {
		// Publish a generated user to the new_users channel
		user := GenerateRandomUser()
		err := redisClient.Publish(ctx, "new_users", user).Err()
		if err != nil {
			panic(err)
		}
		fmt.Println(fmt.Sprintf("Created: %s", user.Email))
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(4)
		time.Sleep(time.Duration(n) * time.Second)
	}
}
