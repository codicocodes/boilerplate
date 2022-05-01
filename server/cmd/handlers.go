package main

import (
	"fmt"
	"net/http"
)

func (app *App) Ping(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "pong\n")
}
