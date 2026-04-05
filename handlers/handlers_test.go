package handlers_test

import (
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
