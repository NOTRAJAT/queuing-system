package main

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

// Define a context for Redis operations
var ctx = context.Background()


func main() {
	// Initialize a Redis client
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // No password by default
		DB:       0,                // Default DB
	})

	// Test connection: Ping
	err := rdb.Ping(ctx).Err()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
	fmt.Println("Connected to Redis successfully!")

	// Set a key
	err = rdb.Set(ctx, "mykey", "Hello, Redis!", 0).Err()
	if err != nil {
		log.Fatalf("Could not set key: %v", err)
	}

	// Get the key
	val, err := rdb.Get(ctx, "mykey").Result()
	if err != nil {
		log.Fatalf("Could not get key: %v", err)
	}
	fmt.Printf("mykey: %s\n", val)

	// Example: Working with Hashes
	err = rdb.HSet(ctx, "myhash", "field1", "value1", "field2", "value2").Err()
	if err != nil {
		log.Fatalf("Could not set hash fields: %v", err)
	}

	// Retrieve hash fields
	fields, err := rdb.HKeys(ctx, "myhash").Result()
	if err != nil {
		log.Fatalf("Could not get hash fields: %v", err)
	}


	fmt.Printf("Hash Fields: %v\n", fields)

	for _,field:= range fields{
		value,err:=rdb.HGet(ctx,"myhash",field).Result()
		if err != nil {
			log.Fatalf("Could not get hash fields: %v", err)
		}
		fmt.Println(field," ",value)

	}
}
