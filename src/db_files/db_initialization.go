package db_files

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

func InitializeDatabase() (*sql.DB, error) {
	connStr := "user=youruser password=yourpassword dbname=tasksdb host=db port=5432 sslmode=disable"
	// Retry connection logic
	var db *sql.DB
	var err error
	for i := 0; i < 5; i++ { 
		db, err = sql.Open("postgres", connStr)
		if err != nil {
			return nil, err
		}

		// Check if the connection is successful
		err = db.Ping()
		if err == nil {
			fmt.Println("Connected to the database")
			return db, nil
		}

		// Wait before retrying
		fmt.Println("Waiting for database to be ready...")
		time.Sleep(2 * time.Second)
	}

	return nil, fmt.Errorf("failed to connect to the database after retries: %v", err)
}

// Create the table if it doesn't exist
func CreateTable(db *sql.DB) error {

	// Define the SQL statement to create the table
	createTableSQL := `
	CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
	CREATE TABLE IF NOT EXISTS tasks (
		id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),  -- UUID as primary key
		task TEXT NOT NULL,
		scheduled_at TIMESTAMP NOT NULL,
		picked_at TIMESTAMP,  
		started_at TIMESTAMP, -- when the worker started executing the task
		completed_at TIMESTAMP, -- when the task was completed (success case)
		failed_at TIMESTAMP -- when the task failed (failure case)
    )`

	// Execute the query
	_, err := db.Exec(createTableSQL)
	if err != nil {
		return err
	}

	fmt.Println("Table created or already exists")
	return nil
}
