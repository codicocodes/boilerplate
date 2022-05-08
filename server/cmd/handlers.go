package main

import (
	"fmt"
	"encoding/json"
	"net/http"

	userservice "github.com/codico/boilerplate/internals/UserService"
)

func SendResponse(w http.ResponseWriter, data any) {
	err := json.NewEncoder(w).Encode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (app *App) Ping(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "pong\n")
}

func (app *App) Register(w http.ResponseWriter, r *http.Request) {
	var input userservice.UserInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid user input", http.StatusBadRequest)
		return
	}
	svc := userservice.New(app.db, input)
	user, err := svc.Register()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	SendResponse(w, user)
}

func (app *App) Login(w http.ResponseWriter, r *http.Request) {
	var input userservice.UserInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid user input", http.StatusBadRequest)
		return
	}
	svc := userservice.New(app.db, input)
	token, err := svc.Login()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	SendResponse(w, token)
}

func (app *App) Refresh(w http.ResponseWriter, r *http.Request) {
	prevToken := parseTokenFromReq(r)
	newToken, err := prevToken.RefreshToken()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	SendResponse(w, newToken)
}

func (app *App) NotifierConnect(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Connection does not support streaming", http.StatusBadRequest)
		return
	}
	d := r.Context().Done()
	notifier, disconnect := app.notifier.Connect()
	for {
		select {
		case <-d:
			disconnect()
			return
		case data := <-notifier:
			fmt.Fprintf(w, "data: %v \n\n", data)
			flusher.Flush()
		}
	}
}
