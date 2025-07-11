// the main file to ingest the data

package main 


import (

	"context"
	"log"
	"go-pgsql/pgsql/db/connection"

)


func main() {
	// the main function to get the 

	host := "localhost"
	userName := "abhi"
	password := "mysecretpassword"
	dbName := "test_db"
	port := 5432

	conn, err := connection.ConnectPgSql(host, userName, password, dbName, port)

	if err != nil {
	log.Fatal(err)
	}
	defer conn.Close(context.Background())

}



