package main

import (
	"log"
	"src/routing"
	"src/db_files"
)



func main() {
	database, err := db_files.InitializeDatabase()
	if err != nil {
		log.Fatal(err)
	}

	// Create the table if it doesn't exist
	err = db_files.CreateTable(database)
	if err != nil {
		log.Fatal(err)
	}
	router := routing.SetupRouter(database)
	router.Run(":8080")
}