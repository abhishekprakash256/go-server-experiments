// the function to make the table and schema 


package crud 

import (
	"context"

	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)


// CreateSchema takes custom SQL strings for login and message table creation.
func CreateSchema(ctx context.Context, pgconnector *pgxpool.Pool, loginSQL, messageSQL string) error {

	ctx := context.Background()
	
	if _, err := pgconnector.Exec(ctx, loginSQL); err != nil {
		return fmt.Errorf("Failed to create login table: %w", err)
	}

	if _, err := pgconnector.Exec(ctx, messageSQL); err != nil {
		return fmt.Errorf("Failed to create message table: %w", err)
	}

	fmt.Println("Tables created (or already exist).")

	return nil
}