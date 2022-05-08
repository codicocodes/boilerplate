package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	userservice "github.com/codico/boilerplate/internals/UserService"
)

func SendResponse[T any](w http.ResponseWriter, data T) {
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
	SendResponse(w, string(*token))
}
