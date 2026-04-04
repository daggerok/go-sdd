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
	})
}

func TestSignupHandler(t *testing.T) {
	t.Run("should create user successfully", func(t *testing.T) {
		// Given: A new user with valid data
		body := map[string]string{"email": "new@example.com", "password": "password123"}
		jsonBody, _ := json.Marshal(body)
		req := httptest.NewRequest("POST", "/signup", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		handlers.SignupHandler(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)

		var response map[string]string
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Equal(t, "User created", response["message"])
	})

	t.Run("should return error for short password", func(t *testing.T) {
		// Given: A user with invalid password (too short)
		body := map[string]string{"email": "test@example.com", "password": "123"}
		jsonBody, _ := json.Marshal(body)
		req := httptest.NewRequest("POST", "/signup", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		handlers.SignupHandler(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)

		var response map[string]string
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Contains(t, response["error"], "Password")
	})
}