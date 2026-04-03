package main_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"go-sdd/handlers"
)

func TestHomeHandler(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)
	handlers.HomeHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
	if w.Header().Get("Content-Type") != "text/html" {
		t.Errorf("Expected text/html")
	}
}
