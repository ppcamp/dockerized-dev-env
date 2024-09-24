package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

func main() {
	// Create a context
	ctx := context.Background()

	// Create a Redis client
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379", // Redis server address
		Password: "",           // No password set
		DB:       0,            // Use default DB
	})

	// Ping the Redis server
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
	fmt.Println(pong) // Should print "PONG"

	// Set a key
	err = rdb.Set(ctx, "key", "ValueStored", 0).Err()
	if err != nil {
		log.Fatalf("Could not set key: %v", err)
	}

	// Get the key
	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		log.Fatalf("Could not get key: %v", err)
	}
	fmt.Println("key:", val) // Should print "key: ValueStored"
}
