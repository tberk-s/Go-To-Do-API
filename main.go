package main

import (
	"log"
	"os"
	"strconv"

	"github.com/tberk-s/Go-To-Do-API/internal/db"
	"github.com/tberk-s/Go-To-Do-API/internal/todo"
	"github.com/tberk-s/Go-To-Do-API/internal/transport"
)

func main() {
	// Load environment variables or set defaults for DB configuration
	dbUser := getEnv("DB_USER", "postgres")
	dbPassword := getEnv("DB_PASSWORD", "password")
	dbHost := getEnv("DB_HOST", "localhost")
	dbName := getEnv("DB_NAME", "todo_app")
	dbPort, err := strconv.Atoi(getEnv("DB_PORT", "5432"))
	if err != nil {
		log.Fatalf("Invalid DB_PORT value: %v", err)
	}

	database, err := db.New(dbUser, dbPassword, dbHost, dbName, dbPort)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer database.Close()

	todoService := todo.New(database)

	server := transport.New(todoService)

	log.Println("Starting server on port 8080...")
	if err := server.Serve(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

// getEnv ...
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
