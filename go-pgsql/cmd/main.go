// the main file to ingest the data

package main 


import (

	"context"
	"log"
	
	"time"
	"go-pgsql/pgsql/db/connection"
	"go-pgsql/config"
	"go-pgsql/pgsql/db/crud"

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
	if !crud.InsertLoginData(ctx, "login", pool, login) {
		log.Println("Insert into login failed")
	}

	// Test message data
	msg := crud.MessageData{
		ChatID:       "abc123",
		SenderName:   "Abhi",
		ReceiverName: "Anny",
		Message:      "Hello from Go!",
		Timestamp:    time.Now(),
		Read:         false,
	}
	if !crud.InsertMessageData(ctx, "message", pool, msg) {
		log.Println("Insert into message failed")
	}

	// Test delete message data
	if !crud.DeleteMessageData(ctx, "message", pool, "abc123") {
		log.Println("Delete message failed")
	}

	// Test delete login data
	if !crud.DeleteLoginData( ctx, "login" , pool, "abc123") {
		log.Println("Delete login data failed")
	}


	log.Println("Data operation done successfully")
}