// to make the connetor of the pgsql 

package connection

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)




func ConnectPgSql(host, userName, password, dbName string, port int) (*pgx.Conn, error) {
	// Format the connection string properly
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", userName, password, host, port, dbName)

	// Connect to PostgreSQL
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatal(" Failed to connect:", err)
	}

	//defer conn.Close(context.Background())

	fmt.Println("Connected to PostgreSQL with pgx!")

	return conn, nil
}