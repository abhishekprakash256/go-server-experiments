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

	defer pool.Close() // <-- closes only when app shuts down
}
