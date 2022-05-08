package main

import (
	"context"
	"net/http"
	"strings"

	userservice "github.com/codico/boilerplate/internals/UserService"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func parseTokenFromReq(r *http.Request) userservice.JwtToken {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	if len(splitToken) != 2 {
		return userservice.JwtToken("")
	}
	reqToken = splitToken[1]
	return userservice.JwtToken(reqToken)
}

type CtxKeyAuthenticated struct{}

func RequireAuth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := parseTokenFromReq(r)
		_, err := token.Validate()
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}
		ctx := context.WithValue(r.Context(), CtxKeyAuthenticated{}, true)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

type CtxKeyUser struct{}

func AddUser(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var ctx context.Context
		token := parseTokenFromReq(r)
		user, _ := token.GetUser()
		ctx = context.WithValue(r.Context(), CtxKeyUser{}, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func CORS(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", getOrigin())
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")
		next.ServeHTTP(w, r)
	})
}

func SSEMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		next.ServeHTTP(w, r)
	})
}

var GlobalMiddlewares = []Middleware{
	CORS,
	AddUser,
}

func RegisterMiddlewares(handler http.HandlerFunc, middlewares []Middleware) http.HandlerFunc {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}
