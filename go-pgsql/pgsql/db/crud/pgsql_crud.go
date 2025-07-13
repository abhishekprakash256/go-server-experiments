// to make the crud operation on the data 

/*
sample data 
| message_id | chat_id | sender_name | receiver_name | message | timestamp          | read |
|------------|---------|-------------|----------------|---------|---------------------|------|
| ...        | abc123  | "Abhi"      | "Anny"         | "Hello" | 2025-07-06 15:00:00 | TRUE |

*/

package crud

import (
	"context"
	"fmt"
	"time"


	"github.com/jackc/pgx/v5/pgxpool"
)


type LoginData struct {
	ChatID string
	UserOne string
	UserTwo string
}


type MessageData struct {
	ChatID       string
	SenderName   string
	ReceiverName string
	Message      string
	Timestamp    time.Time
	Read         bool
}




// InsertLoginData inserts a row into the login table.
func InsertLoginData(ctx context.Context, tableName string, pgconnector *pgxpool.Pool, data LoginData) bool {
	query := fmt.Sprintf(`
		INSERT INTO %s (chat_id, users_1, users_2)
		VALUES ($1, $2, $3)
		ON CONFLICT (chat_id) DO NOTHING
	`, tableName)

	_, err := pgconnector.Exec(ctx, query, data.ChatID, data.UserOne, data.UserTwo)
	if err != nil {
		fmt.Println("Insert into login failed:", err)
		return false
	}

	fmt.Println("Login inserted (or already exists).")
	return true
}



func InsertMessageData(ctx context.Context, tableName string, pgconnector *pgxpool.Pool, data MessageData) bool {

	// insert the data into the message table

	query := fmt.Sprintf(`
		INSERT INTO %s (chat_id, sender_name, receiver_name, message, timestamp, read)
		VALUES ($1, $2, $3, $4, $5, $6)
	`, tableName)

	_, err := pgconnector.Exec(
		ctx,
		query,
		data.ChatID,
		data.SenderName,
		data.ReceiverName,
		data.Message,
		data.Timestamp,
		data.Read,
	)

	if err != nil {
		fmt.Println("Insert failed:", err)
		return false
	}

	fmt.Println("Message inserted")
	return true
}


func DeleteLoginData(ctx context.Context, tableName string, pgconnector *pgxpool.Pool, chatID string) bool {

	// delete the message per id

	query := fmt.Sprintf(`DELETE FROM %s WHERE chat_id = $1`, tableName)

	_, err := pgconnector.Exec(ctx, query, chatID)
	
	if err != nil {
		fmt.Println("Delete failed:", err)
		return false
	}

	fmt.Println("Login data deleted for chat_id:", chatID)
	return true
}


func DeleteMessageData(ctx context.Context, tableName string, pgconnector *pgxpool.Pool, chatID string) bool {

	// delete the message per id

	query := fmt.Sprintf(`DELETE FROM %s WHERE chat_id = $1`, tableName)

	_, err := pgconnector.Exec(ctx, query, chatID)

	if err != nil {
		fmt.Println("Delete failed:", err)
		return false
	}

	fmt.Println("Messages deleted for chat_id:", chatID)
	return true
}

