package server

import (
	"net/http"

	"github.com/daggerok/go-sdd/handlers"
)

func NewServer(addr string) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/signup", handlers.SignupHandler)

	return &http.Server{
		Addr:    addr,
		Handler: mux,
	}
}
