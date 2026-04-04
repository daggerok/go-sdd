package main

import (
	"fmt"
	"net/http"

	"github.com/daggerok/go-sdd/handlers"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/signup", handlers.SignupHandler)

	port := 8081
	fmt.Println("Server running on port", port)
	fmt.Printf("Visit http://localhost:%d to see the home page\n", port)

	listen := fmt.Sprintf(":%d", port)
	if err := http.ListenAndServe(listen, nil); err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
