package main_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"go-sdd/handlers"

	"github.com/stretchr/testify/assert"
)

func TestHomeHandler(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)
	handlers.HomeHandler(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "text/html", w.Header().Get("Content-Type"))
}
