package handlers

import (
	"encoding/json"
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

type SignupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupResponse struct {
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req SignupRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(SignupResponse{Error: "Invalid JSON"})
		return
	}

	if req.Email == "" || req.Password == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(SignupResponse{Error: "Email and password required"})
		return
	}

	if len(req.Password) < 6 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(SignupResponse{Error: "Password must be at least 6 characters"})
		return
	}

	// In a real application, you would check if email already exists
	// For now, we'll just return success
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(SignupResponse{Message: "User created"})
}
