package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	DB_NAME        string = "mydb"
	DB_USER        string = "myuser"
	DB_PASSWORD    string = "mypass"
	DB_HOST        string = "localhost"
	DB_PORT        string = "5432"
	PGBOUNCER_PORT string = "6432"
)

func InitDatabase() {
	// Retrieve the connection string for PgBouncer
	connStr := getConnectionString()

	// Initialize a connection to the database using `database/sql`
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error initializing database connection: %v", err)
	}

	// Verify the PgBouncer connection
	if err := verifyPgBouncerConnection(); err != nil {
		log.Fatalf("PgBouncer connection verification failed: %v", err)
	}

	log.Println("Database initialized successfully and PgBouncer connection verified.")
}

func getConnectionString() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable binary_parameters=yes",
		DB_HOST,
		DB_USER,
		DB_PASSWORD,
		DB_NAME,
		PGBOUNCER_PORT, // DB_PORT, postgres port is not used anymore because of using PgBouncer
	)
}

func verifyPgBouncerConnection() error {
	// Ping the database to check the connection
	if err := db.Ping(); err != nil {
		return fmt.Errorf("unable to ping the database: %w", err)
	}
	return nil
}

func simulateConcurrentConnections(count int) {
	var wg sync.WaitGroup

	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			conn, err := db.Conn(context.Background())
			if err != nil {
				log.Printf("Connection %d failed: %v", id, err)
				return
			}
			defer conn.Close()

			// Perform a simple query
			var result string
			err = conn.QueryRowContext(context.Background(), "SELECT 'PgBouncer Test'").Scan(&result)
			if err != nil {
				log.Printf("Query on connection %d failed: %v", id, err)
				return
			}
			log.Printf("Connection %d succeeded: %s", id, result)
		}(i)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	log.Println("All concurrent connections completed.")
}

func executeQuery() {
	rows, err := db.Query("SELECT current_timestamp")
	if err != nil {
		log.Printf("Query failed: %v", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var timestamp string
		if err := rows.Scan(&timestamp); err != nil {
			log.Printf("Failed to read row: %v", err)
		} else {
			log.Printf("Current timestamp: %s", timestamp)
		}
	}
}

func Cleanup() {
	if db != nil {
		if err := db.Close(); err != nil {
			log.Printf("Error closing database connection: %v", err)
		} else {
			log.Println("Database connection closed.")
		}
	}
}

func main() {
	log.Println("Starting PgBouncer demonstration...")

	InitDatabase()
	defer Cleanup()

	log.Println("Executing query to verify functionality...")
	executeQuery()

	log.Println("Simulating multiple concurrent connections...")
	simulateConcurrentConnections(10)

	log.Println("PgBouncer demonstration completed.")
}
