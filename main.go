package main

import (
	"fmt"
	"net/http"

	"go-sdd/handlers"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	fmt.Println("Server running on port 8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
