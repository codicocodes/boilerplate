package main

import (
	"log"
	"net/http"

	"github.com/codico/boilerplate/db"
	notifierservice "github.com/codico/boilerplate/internals/NotifierService"
	"github.com/julienschmidt/httprouter"
)

type App struct {
	log      *log.Logger
	db       *db.Queries
	notifier *notifierservice.NotifierService
}

func newApp(db *db.Queries, logger *log.Logger, notifier *notifierservice.NotifierService) App {
	return App{
		log:      logger,
		db:       db,
		notifier: notifier,
	}
}

func (app *App) Router() http.Handler {
	r := httprouter.New()
	r.NotFound = http.HandlerFunc(http.NotFound)
	r.MethodNotAllowed = http.HandlerFunc(http.NotFound)
	routes := app.getRoutes()
	registerRoutes(r, routes)
	return r
}
