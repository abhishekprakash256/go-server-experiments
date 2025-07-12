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
		log.Fatal(err)
	}

	
	// Create the database schema
	err = crud.CreateSchema(ctx, pool, config.LoginTableSQL, config.MessageTableSQL)
	if err != nil {
		log.Fatal("Schema creation failed:", err)
	}

	//test login data

	login := crud.LoginData{
	ChatID:  "abc123",
	UserOne: "Abhi",
	UserTwo: "Anny",
	}


	ok := crud.InsertLoginData(ctx, "login", pool, login)

	if !ok {
		log.Println("Insert into login failed")
	}

	// test message 

	msg := crud.MessageData{
	ChatID:       "abc123",
	SenderName:   "Abhi",
	ReceiverName: "Anny",
	Message:      "Hello from Go!",
	Timestamp:    time.Now(),
	Read:         false,
	}

	// insert the message 
	ok = crud.InsertMessage(ctx, "message", pool, msg)
	if !ok {
		log.Println("Insert failed")
	}



	defer pool.Close() // <-- closes only when app shuts down
}
