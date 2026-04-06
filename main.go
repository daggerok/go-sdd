package main

import (
	"fmt"
	"net/http"

	"github.com/daggerok/go-sdd/server"
)

func main() {
	port := 8081
	srv := server.NewServer(fmt.Sprintf(":%d", port))

	fmt.Println("Server running on port", port)
	fmt.Printf("Visit http://localhost:%d to see the home page\n", port)

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Println("Server failed to start:", err)
	}
}
