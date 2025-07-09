package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	// Connect to Redis
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		Protocol: 2,
	})

	// Context
	ctx := context.Background()

	// Session data
	key1 := "session:abc123:Abhi"

	key2 := "session:abc123:Anny"

	data1 := map[string]interface{}{
		"chat_id":      "abc123",
		"user":         "Abhi",
		"last_seen":    time.Now().Format(time.RFC3339), // or use hardcoded string
		"ws_connected": 1,
		"notify":       0,
	}

	data2 := map[string]interface{}{
		"chat_id":      "abc123",
		"user":         "Anny",
		"last_seen":    time.Now().Format(time.RFC3339), // or use hardcoded string
		"ws_connected": 0,
		"notify":       1,
	}


	// Store hash in Redis
	err := client.HSet(ctx, key1, data1).Err()
	if err != nil {
		panic(err)
	}


	err2 := client.HSet(ctx, key2, data2).Err()
	if err != nil {
		panic(err2)
	}


	// Retrieve and print hash to verify
	result1, err1 := client.HGetAll(ctx, key1).Result()
	if err1 != nil {
		panic(err1)
	}

	fmt.Printf("Stored Hash at key %s:\n", key1)
	for field1, val1 := range result1 {
		fmt.Printf("  %s → %s\n", field1, val1)
	}



	// Retrieve and print hash to verify
	result2, err2 := client.HGetAll(ctx, key2).Result()
	if err2 != nil {
		panic(err2)
	}

	fmt.Printf("Stored Hash at key %s:\n", key2)
	for field2, val2 := range result2 {
		fmt.Printf("  %s → %s\n", field2, val2)
	}
}
