package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/codico/boilerplate/db"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)


func main() {
	logger := log.Default()
	conn := connectDB(logger)
	defer conn.Close()
	db := db.New(conn)
	app := newApp(db, logger)
	server := http.Server{
		Addr:         fmt.Sprintf(":%s", getPort()),
		Handler:      app.Router(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	logger.Printf("Server running on port %s", getPort())
	err := server.ListenAndServe()
	if err != nil {
		logger.Fatalf(err.Error())
	}
}
