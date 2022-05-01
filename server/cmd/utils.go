package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

// Connect to PostgreSQL with DATABASE_URL from env variables
func connectDB(logger *log.Logger) *sql.DB {
  db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
  if err != nil {
		logger.Fatalf("unable to connect to the database: %s", err.Error())
  }
	err = db.Ping()
  if err != nil {
		logger.Fatalf("unable to connect to the database: %s", err.Error())
  }
	logger.Println("Successfully connected to DB.")
  return db
}

func getPort() string {
	port := os.Getenv("PORT");
	if port == "" {
		return DEFAULT_PORT
	}
	return port
}

func getOrigin() string {
	clientUri := os.Getenv("ALLOWED_ORIGIN");
	if clientUri == "" {
		fmt.Println("ALLOWED_ORIGIN not configured defaulting to *")
		return "*"
	}
	return clientUri
}
