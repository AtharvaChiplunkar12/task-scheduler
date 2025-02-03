package db_files

import (
	"src/models"
	"database/sql"
	"log"
)

func AddRequestDB(db *sql.DB, request models.UserRequest) (string, error) {
	// Define the SQL statement to insert the request
	insertSQL := `
		INSERT INTO tasks (task, scheduled_at) 
		VALUES ($1, $2) RETURNING id;
	`

	var id string
	// Execute the query and get the generated ID
	err := db.QueryRow(insertSQL, request.Task, request.ScheduledAt).Scan(&id)
	if err != nil {
		// Handle error and log it
		log.Printf("Error inserting request into database: %v", err)
		return "", err
	}

	// Return the generated ID
	return id, nil
}

func GetRequestDB(db *sql.DB, id string) (*models.Task, error) {
	// Define the SQL query to fetch the request
	query := `
		SELECT id, task, scheduled_at, picked_at, started_at, completed_at, failed_at
		FROM tasks 
		WHERE id = $1;
	`

	// Create a variable to hold the result
	var request models.Task

	// Execute the query
	err := db.QueryRow(query, id).Scan(
		&request.Id,
		&request.Task,
		&request.ScheduledAt,
		&request.PickedAt,
		&request.StartedAt,
		&request.CompletedAt,
		&request.FailedAt,
	)
	if err != nil {
		log.Printf("Error retrieving request from database: %v", err)
		return nil, err
	}

	// Return the retrieved request
	return &request, nil
}