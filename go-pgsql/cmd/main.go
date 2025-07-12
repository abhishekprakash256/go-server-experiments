// the main file to ingest the data

package main 


import (

	"context"
	"log"
	"go-pgsql/pgsql/db/connection"
	"go-pgsql/config"
	"go-pgsql/pgsql/db/crud"

)


func main() {
	
	//create the connection 
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
	if err := crud.CreateSchema(pool, config.LoginTableSQL, config.MessageTableSQL); err != nil {
		log.Fatal("Schema creation failed:", err)
	}

	// close the connection
	defer pool.Close(context.Background())

}