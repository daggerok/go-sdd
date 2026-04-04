package main

import (
	"fmt"
	"net/http"

	"github.com/daggerok/go-sdd/handlers"
)

func newServer(addr string) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/signup", handlers.SignupHandler)

	return &http.Server{
		Addr:    addr,
		Handler: mux,
	}
}

func main() {
	port := 8081
	srv := newServer(fmt.Sprintf(":%d", port))

	fmt.Println("Server running on port", port)
	fmt.Printf("Visit http://localhost:%d to see the home page\n", port)

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Println("Server failed to start:", err)
	}
}
