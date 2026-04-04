package main_test

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
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)
	handlers.HomeHandler(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "text/html", w.Header().Get("Content-Type"))
}

func TestSignupHandler(t *testing.T) {
	// Test successful signup
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
}

func TestSignupHandler_InvalidData(t *testing.T) {
	// Test invalid data (short password)
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
}
