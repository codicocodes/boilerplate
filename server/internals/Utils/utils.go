package utils

import (
	"database/sql"
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

