package handlers

import (
	"fmt"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	isAuthenticated := false // Placeholder
	// In a real application, you would check for a session token or JWT
	// For simplicity, let's assume authenticated if not on signin/signup and not redirected
	if r.Context().Value("authenticated") != nil && r.Context().Value("authenticated").(bool) {
		isAuthenticated = true
	}

	var authLink string
	if isAuthenticated {
		authLink = `<p><a href="/logout">Sign Out</a></p>`
	} else {
		authLink = `<p><a href="/signin">Sign In</a></p>`
	}

	fmt.Fprintf(w, "<html><body><h1>Welcome</h1>%s</body></html>", authLink)
}
