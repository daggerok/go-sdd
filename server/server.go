package server

import (
	"net/http"

	"github.com/daggerok/go-sdd/handlers"
)

func NewServer(addr string) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/signup", handlers.SignupHandler)
	mux.HandleFunc("/signin", handlers.SigninHandler)

	// Apply middleware to all routes except /, /signin, /signup
	protectedMux := http.NewServeMux()
	protectedMux.Handle("/", handlers.AuthMiddleware(mux.ServeHTTP))

	return &http.Server{
		Addr:    addr,
		Handler: protectedMux,
	}
}
