/*
The main file to 
*/

package main 


import (

	"log"
	"go-redis/config" 
	"go-redis/redis/db/connection"

)

func main() {

	//making the connection

	client, err := connection.ConnectRedis(config.DefaultConfig.Host, config.DefaultConfig.Port)

	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	defer client.Close()

}