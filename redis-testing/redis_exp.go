// to test the redis functionlaity in go 


/*
using the redis docs from here -- 

https://redis.io/docs/latest/develop/clients/go/

*/

package main



import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)



func makeClient(host string) *redis.Client {
		client := redis.NewClient(&redis.Options{
		Addr:      host + ":6379", // Use the passed host,
		Password: "", // No password
		DB:       0,  // Default DB
		Protocol: 2,  // RESP2
	})


	return client

}




func main() {
	// Create Redis client using the factory function
	client := makeClient("localhost")

	// Create context
	ctx := context.Background()

	// Set key "foo" with value "bar"
	err := client.Set(ctx, "foo", "bar", 0).Err()
	if err != nil {
		panic(err)
	}

	// Get key "foo"
	val, err := client.Get(ctx, "foo").Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("foo:", val)

	hashFields := []string{
    "model", "Deimos",
    "brand", "Ergonom",
    "type", "Enduro bikes",
    "price", "4972",
	}

	res1, err := client.HSet(ctx, "bike:1", hashFields).Result()

	if err != nil {
		panic(err)
	}

	fmt.Println(res1) // >>> 4

	res2, err := client.HGet(ctx, "bike:1", "model").Result()

	if err != nil {
		panic(err)
	}

	fmt.Println(res2) // >>> Deimos

	res3, err := client.HGet(ctx, "bike:1", "price").Result()

	if err != nil {
		panic(err)
	}

	fmt.Println(res3) // >>> 4972

	res4, err := client.HGetAll(ctx, "bike:1").Result()

	if err != nil {
		panic(err)
	}

	fmt.Println(res4)

}
