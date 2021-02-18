package main

import (
	"context"
	"math/rand"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	// Create a new Redis Client
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",  // We connect to host redis, thats what the hostname of the redis service is set to in the docker-compose
		Password: "superSecret", // The password IF set in the redis Config file
		DB:       0,
	})
	// Ping the Redis server and check if any errors occurred
	// err := redisClient.Ping(context.Background()).Err()
	// if err != nil {
	// 	// Sleep for 3 seconds and wait for Redis to initialize
	// 	time.Sleep(3 * time.Second)
	// 	err := redisClient.Ping(context.Background()).Err()
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
	// Generate a new background context that  we will use
	ctx := context.Background()
	// Loop and randomly generate users on a random timer
	for {
		// Publish a generated user to the new_users channel
		err := redisClient.Publish(ctx, "new_users", GenerateRandomUser()).Err()
		if err != nil {
			panic(err)
		}
		// Sleep random time
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(4)
		time.Sleep(time.Duration(n) * time.Second)
	}
}
