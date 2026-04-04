package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/daggerok/go-sdd/handlers"

	"github.com/stretchr/testify/assert"
)

func TestHomeHandler(t *testing.T) {
	t.Run("should return OK with HTML content", func(t *testing.T) {
		// Given: A request to the home page
		req := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()

		// When: Handling the request
		handlers.HomeHandler(w, req)

		// Then: Should return OK with HTML content
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "text/html", w.Header().Get("Content-Type"))
		assert.Contains(t, w.Body.String(), "Welcome")
	})
}

func TestSignupHandler(t *testing.T) {
	t.Run("should create user successfully", func(t *testing.T) {
		// Given: A new user with valid data
		w := httptest.NewRecorder()
		body := map[string]string{"email": "new@example.com", "password": "password123"}
		jsonBody, _ := json.Marshal(body)
		req := httptest.NewRequest("POST", "/signup", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")

		// When: Posting signup request
		handlers.SignupHandler(w, req)

		// Then: Should create user
		assert.Equal(t, http.StatusCreated, w.Code)

		// And: Response should confirm user creation
		var response map[string]string
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Equal(t, "User created", response["message"])
	})

	t.Run("should not create user with short password", func(t *testing.T) {
		// Given: A user with invalid password (too short)
		w := httptest.NewRecorder()
		body := map[string]string{"email": "test@example.com", "password": "123"}
		jsonBody, _ := json.Marshal(body)
		req := httptest.NewRequest("POST", "/signup", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")

		// When: Posting signup request
		handlers.SignupHandler(w, req)

		// Then: Should return bad request
		assert.Equal(t, http.StatusBadRequest, w.Code)

		// And: Error message should mention password
		var response map[string]string
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Contains(t, response["error"], "Password")
	})

	t.Run("should not create user if email already exists", func(t *testing.T) {
		// Given: A user with invalid password (too short)
		w := httptest.NewRecorder()
		body := map[string]string{"email": "test@go-sdd.example.com", "password": "123456789"}
		jsonBody, _ := json.Marshal(body)
		req := httptest.NewRequest("POST", "/signup", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")

		// When: Posting signup request
		handlers.SignupHandler(w, req)

		// Then: Should return bad request
		assert.Equal(t, http.StatusBadRequest, w.Code)

		// And: Error message should mention password
		var response map[string]string
		json.Unmarshal(w.Body.Bytes(), &response)
		t.Logf("response: %v", response)
		assert.Equal(t, response["error"], "Email already exists")
	})
}
