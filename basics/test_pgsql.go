// the main file to ingest the data
package main


import (

	"context"
	"log"
	"fmt"
	
	"time"
	"github.com/abhishekprakash256/go-pgsql-helper-kit/pgsql/db/connection"
	"github.com/abhishekprakash256/go-pgsql-helper-kit/config"
	"github.com/abhishekprakash256/go-pgsql-helper-kit/pgsql/db/crud"

)


func main() {
	ctx := context.Background()

	// Create the connection pool
	pool, err := connection.ConnectPgSql(
		config.DefaultConfig.Host,
		config.DefaultConfig.User,
		config.DefaultConfig.Password,
		config.DefaultConfig.DBName,
		config.DefaultConfig.Port,
	)
	
	// The connection failed
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}
	
	defer pool.Close() // Ensures pool is closed when program exits

	// Create the database schema
	if err := crud.CreateSchema(ctx, pool, config.LoginTableSQL, config.MessageTableSQL); err != nil {

		log.Fatal("Schema creation failed:", err)
	}

	// Test login data
	login := crud.LoginData{
		ChatID:  "abc123",
		UserOne: "Abhi",
		UserTwo: "Anny",
	}
	
	// login the data
	if !crud.InsertLoginData(ctx, "login", pool, login) {
		log.Println("Insert into login failed")
	}

	// Test message data
	msg := crud.MessageData{
		ChatID:       "abc123",
		SenderName:   "Abhi",
		ReceiverName: "Anny",
		Message:      "Hello There!",
		Timestamp:    time.Now(),
		Read:         false,
	}

	// Insert the message data
	if !crud.InsertMessageData(ctx, "message", pool, msg) {
		log.Println("Insert into message failed")
	}

	// Step 5: Retrieve login data
	retrievedLogin, err := crud.GetLoginData(ctx, "login", pool, "abc123")
	if err != nil {
		log.Println("Login not found:", err)
	} else {
		fmt.Printf("Login for chat %s: %s & %s\n", retrievedLogin.ChatID, retrievedLogin.UserOne, retrievedLogin.UserTwo)
	}

	// Step 6: Retrieve message data
	messages := crud.GetMessageData(ctx, "message", pool, "abc123", "Abhi")

	// print the message data
	fmt.Printf("Messages: %+v\n", messages)

	// to print the message from the database
	for _, m := range messages {
		fmt.Printf("Message from %s to %s: %s\n", m.SenderName, m.ReceiverName, m.Message , m.Timestamp , m.Read)
	}

	// Test delete message data
	if !crud.DeleteMessageData(ctx, "message", pool, "abc123") {
		log.Println("Delete message failed")
	}
	
	// Test delete login data
	if !crud.DeleteLoginData( ctx, "login" , pool, "abc123") {
		log.Println("Delete login data failed")
	}

	//Done with the operation
	log.Println("Data operation done successfully")
}