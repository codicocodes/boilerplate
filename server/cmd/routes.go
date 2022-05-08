package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Method string

const (
	GET    Method = http.MethodGet
	PUT           = http.MethodPut
	POST          = http.MethodPost
	DELETE        = http.MethodDelete
)

type Version string

const (
	v1 Version = "v1"
)

type Route struct {
	Path        string
	Handler     http.HandlerFunc
	Method      Method
	Version     Version
	Middlewares []Middleware
}

func registerRoutes(r *httprouter.Router, routes []Route) {
	for _, route := range routes {
		path := fmt.Sprintf("/%s%s", route.Version, route.Path)
		handler := RegisterMiddlewares(
			RegisterMiddlewares(route.Handler, GlobalMiddlewares),
			route.Middlewares,
		)
		r.HandlerFunc(string(route.Method), path, handler)
	}
}

func (app *App) getRoutes() []Route {
	return []Route{
		{
			Path:        "/ping",
			Method:      http.MethodGet,
			Handler:     app.Ping,
			Version:     v1,
			Middlewares: []Middleware{},
		},
		{
			Path:        "/register",
			Method:      http.MethodPost,
			Handler:     app.Register,
			Version:     v1,
			Middlewares: []Middleware{},
		},
		{
			Path:        "/login",
			Method:      http.MethodPost,
			Handler:     app.Login,
			Version:     v1,
			Middlewares: []Middleware{},
		},
		{
			Path:        "/notifier",
			Method:      http.MethodGet,
			Handler:     app.NotifierConnect,
			Version:     v1,
			Middlewares: []Middleware{SSEMiddleware},
		},
	}
}
