package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"golang.org/x/crypto/bcrypt"
)

var (
	users = make(map[string]string) // email -> hashed password
	mu    sync.RWMutex
)

func init() {
	// Pre-registered user
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	users["test@go-sdd.example.com"] = string(hashedPassword)
}

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

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, `{"error": "Method not allowed"}`)
		return
	}

	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error": "Invalid JSON"}`)
		return
	}

	if req.Email == "" || req.Password == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error": "Email and password required"}`)
		return
	}

	if len(req.Password) < 6 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error": "Password must be at least 6 characters"}`)
		return
	}

	mu.Lock()
	if _, exists := users[req.Email]; exists {
		mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error": "Email already exists"}`)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "Failed to hash password"}`)
		return
	}
	users[req.Email] = string(hashedPassword)
	mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, `{"message": "User created"}`)
}

func SigninHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `<html><body><h1>Sign In</h1><form method="POST" action="/signin"><label for="email">Email:</label><br><input type="email" id="email" name="email"><br><label for="password">Password:</label><br><input type="password" id="password" name="password"><br><input type="submit" value="Submit"></form><p>Don't have an account? <a href="/signup">Sign Up</a></p></body></html>`)
		return
	}

	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, `{"error": "Method not allowed"}`)
		return
	}

	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if r.Header.Get("Content-Type") == "application/json" {
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, `{"error": "Invalid JSON"}`)
			return
		}
	} else {
		r.ParseForm()
		req.Email = r.FormValue("email")
		req.Password = r.FormValue("password")
	}

	mu.RLock()
	hashedPassword, exists := users[req.Email]
	mu.RUnlock()

	if !exists || bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password)) != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, `{"error": "Invalid credentials"}`)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	http.Redirect(w, r, "/", http.StatusFound)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// In a real application, you would invalidate the session/JWT here
	// For simplicity, we just redirect to home page
	http.Redirect(w, r, "/", http.StatusFound)
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" || r.URL.Path == "/signin" || r.URL.Path == "/signup" {
			next.ServeHTTP(w, r)
			return
		}

		// Here you would typically check for a session token or JWT
		// For simplicity, we'll just check if a user is authenticated
		// This is a placeholder and should be replaced with actual authentication logic
		// For now, let's assume no one is authenticated for paths other than /, /signin, /signup

		// For simplicity, we'll just check if a user is authenticated
		// This is a placeholder and should be replaced with actual authentication logic
		// For now, let's assume no one is authenticated for paths other than /, /signin, /signup

		// If not authenticated, redirect to signin
		http.Redirect(w, r, "/signin", http.StatusFound)
		return
	}

	ctx := context.WithValue(r.Context(), "authenticated", true)
	next.ServeHTTP(w, r.WithContext(ctx))
}
