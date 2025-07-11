// the main file to ingest the data

package main 


import (

	"context"
	"log"
	"go-pgsql/pgsql/db/connection"
	"go-pgsql/config"

)


func main() {

	// the main function to eastablish the connection

	conn, err := connection.ConnectPgSql(
		config.DefaultConfig.Host,
		config.DefaultConfig.User,
		config.DefaultConfig.Password,
		config.DefaultConfig.DBName,
		config.DefaultConfig.Port,
	)

	if err != nil {
	log.Fatal(err)
	}
	defer conn.Close(context.Background())

}



