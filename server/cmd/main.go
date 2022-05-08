package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/codico/boilerplate/db"
	notifierservice "github.com/codico/boilerplate/internals/NotifierService"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func main() {
	logger := log.Default()
	conn := connectDB(logger)
	defer conn.Close()
	db := db.New(conn)
	notifier := notifierservice.New()

	go func() {
		for {
			notifier.BroadcastAll(time.Now().String())
			time.Sleep(time.Second * 5)
		}
	}()

	app := newApp(db, logger, notifier)
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
