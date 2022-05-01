package main

import (
	"log"
	"net/http"

	"github.com/codico/boilerplate/db"
	"github.com/julienschmidt/httprouter"
)

type App struct {
	log *log.Logger
	db  *db.Queries
}

func newApp(db *db.Queries, logger *log.Logger) App {
	return App{
		log: logger,
		db:  db,
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

