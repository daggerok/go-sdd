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
	t.Run("should access home page", func(t *testing.T) {
		// Given: A request to the home page
		req := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()

		// When: Handling the request
		handlers.HomeHandler(w, req)

		// Then: Should return OK with HTML content
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "text/html", w.Header().Get("Content-Type"))
		assert.Contains(t, w.Body.String(), "Welcome", "Sign In")
	})
}

func TestSignupHandler(t *testing.T) {
	t.Run("should not accept non-POST requests", func(t *testing.T) {
		// Given: A GET request to signup endpoint
		req := httptest.NewRequest("GET", "/signup", nil)
		w := httptest.NewRecorder()

		// When: Handling the request
		handlers.SignupHandler(w, req)

		// Then: Should return Method Not Allowed
		assert.Equal(t, http.StatusMethodNotAllowed, w.Code)
	})

	t.Run("should not accept invalid JSON", func(t *testing.T) {
		// Given: A POST request with invalid JSON
		req := httptest.NewRequest("POST", "/signup", bytes.NewBufferString("invalid json"))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		// When: Handling the request
		handlers.SignupHandler(w, req)

		// Then: Should return Bad Request with error message
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
		var resp handlers.SignupResponse
		err := json.NewDecoder(w.Body).Decode(&resp)
		assert.NoError(t, err)
		assert.Equal(t, "Invalid JSON", resp.Error)
	})

	t.Run("should not signup if missing email", func(t *testing.T) {
		// Given: A POST request with empty email
		reqBody := handlers.SignupRequest{Email: "", Password: "password123"}
		bodyBytes, _ := json.Marshal(reqBody)
		req := httptest.NewRequest("POST", "/signup", bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		// When: Handling the request
		handlers.SignupHandler(w, req)

		// Then: Should return Bad Request with error message
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
		var resp handlers.SignupResponse
		err := json.NewDecoder(w.Body).Decode(&resp)
		assert.NoError(t, err)
		assert.Equal(t, "Email and password required", resp.Error)
	})

	t.Run("should not signup if missing password", func(t *testing.T) {
		// Given: A POST request with empty password
		reqBody := handlers.SignupRequest{Email: "test@example.com", Password: ""}
		bodyBytes, _ := json.Marshal(reqBody)
		req := httptest.NewRequest("POST", "/signup", bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		// When: Handling the request
		handlers.SignupHandler(w, req)

		// Then: Should return Bad Request with error message
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
		var resp handlers.SignupResponse
		err := json.NewDecoder(w.Body).Decode(&resp)
		assert.NoError(t, err)
		assert.Equal(t, "Email and password required", resp.Error)
	})

	t.Run("should not signup if password is too short", func(t *testing.T) {
		// Given: A POST request with short password
		reqBody := handlers.SignupRequest{Email: "test@example.com", Password: "12345"}
		bodyBytes, _ := json.Marshal(reqBody)
		req := httptest.NewRequest("POST", "/signup", bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		// When: Handling the request
		handlers.SignupHandler(w, req)

		// Then: Should return Bad Request with error message
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
		var resp handlers.SignupResponse
		err := json.NewDecoder(w.Body).Decode(&resp)
		assert.NoError(t, err)
		assert.Equal(t, "Password must be at least 6 characters", resp.Error)
	})

	t.Run("should accept valid signup request", func(t *testing.T) {
		// Given: A POST request with valid credentials
		reqBody := handlers.SignupRequest{Email: "test@example.com", Password: "password123"}
		bodyBytes, _ := json.Marshal(reqBody)
		req := httptest.NewRequest("POST", "/signup", bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		// When: Handling the request
		handlers.SignupHandler(w, req)

		// Then: Should return Created with success message
		assert.Equal(t, http.StatusCreated, w.Code)
		assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
		var resp handlers.SignupResponse
		err := json.NewDecoder(w.Body).Decode(&resp)
		assert.NoError(t, err)
		assert.Equal(t, "User created", resp.Message)
	})
}
