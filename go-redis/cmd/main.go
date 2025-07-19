/*
The main file to 
*/

package main 


import (

	"context"
	"fmt"
	"time"
	"log"

	"go-redis/config" 
	"go-redis/redis/db/connection"
	"go-redis/redis/db/crud"

)

func main() {
	ctx := context.Background()

	// Making the connection
	client, err := connection.ConnectRedis(config.DefaultConfig.Host, config.DefaultConfig.Port)
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	defer client.Close()

	// Prepare test session data
	sessionKey := "session:abc123:Abhi"
	session := config.SessionData{
		ChatID:      "abc123",
		Sender:        "Abhi",
		Reciever: 		"Anny",
		LastSeen:    time.Now(),
		WSConnected: 1,
		Notify:      0,
	}

	// Store session data
	ok := crud.StoreSessionData(ctx, client, sessionKey, session)
	if !ok {
		log.Println("Failed to store session data")
	} else {
		log.Println("Session data stored successfully")
	}

	// Retrieve session data
	retrieved, err := crud.GetSessionData(ctx, client, sessionKey)
	if err != nil {
		log.Println("Failed to fetch session data:", err)
	} else {
		fmt.Printf("Fetched session data: %+v\n", retrieved)
	}

	// Delete session data
	deleted := crud.DeleteSessionData(ctx, client, sessionKey)
	if !deleted {
		log.Println("Failed to delete session data")
	} else {
		log.Println("Session data deleted successfully")
	}


	// Prepare test session data
	sessionKeyTwo := "session:abc123:Abhi"
	sessiontwo := config.SessionData{
		ChatID:      "abc123",
		Sender:        "Anny",
		Reciever: 		"Abhi",
		LastSeen:    time.Now(),
		WSConnected: 1,
		Notify:      0,
	}

	// Store session data
	ok = crud.StoreSessionData(ctx, client, sessionKeyTwo, sessiontwo)
	if !ok {
		log.Println("Failed to store session data")
	} else {
		log.Println("Session data stored successfully")
	}

	// Retrieve session data
	retrieved, err = crud.GetSessionData(ctx, client, sessionKeyTwo)
	if err != nil {
		log.Println("Failed to fetch session data:", err)
	} else {
		fmt.Printf("Fetched session data: %+v\n", retrieved)
	}

	// Delete session data
	deleted = crud.DeleteSessionData(ctx, client, sessionKeyTwo)
	if !deleted {
		log.Println("Failed to delete session data")
	} else {
		log.Println("Session data deleted successfully")
	}
}
