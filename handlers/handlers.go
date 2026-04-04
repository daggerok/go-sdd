package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

var (
	users = make(map[string]string) // email -> password
	mu    sync.RWMutex
)

func init() {
	// Pre-registered user
	users["test@go-sdd.example.com"] = "password123"
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "<html><body><h1>Welcome</h1></body></html>")
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error": "Invalid JSON"}`, http.StatusBadRequest)
		return
	}

	if req.Email == "" || req.Password == "" {
		http.Error(w, `{"error": "Email and password required"}`, http.StatusBadRequest)
		return
	}

	if len(req.Password) < 6 {
		http.Error(w, `{"error": "Password must be at least 6 characters"}`, http.StatusBadRequest)
		return
	}

	mu.Lock()
	if _, exists := users[req.Email]; exists {
		mu.Unlock()
		http.Error(w, `{"error": "Email already exists"}`, http.StatusBadRequest)
		return
	}
	users[req.Email] = req.Password
	mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, `{"message": "User created"}`)
}
