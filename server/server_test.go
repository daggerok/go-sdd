package server_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/daggerok/go-sdd/server"
	"github.com/stretchr/testify/assert"
)

func TestE2EServer(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	assert.NoError(t, err)

	srv := server.NewServer(listener.Addr().String())
	go func() {
		_ = srv.Serve(listener)
	}()

	t.Cleanup(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_ = srv.Shutdown(ctx)
	})

	baseURL := "http://" + listener.Addr().String()

	t.Run("should show home page", func(t *testing.T) {
		// When: Getting home page
		resp, err := http.Get(baseURL + "/")

		// Then: Should return home page successfully
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, "text/html", resp.Header.Get("Content-Type"))
		_ = resp.Body.Close()
	})

	t.Run("should signup new user", func(t *testing.T) {
		// When: Posting signup request
		payload := map[string]string{"email": "e2e@example.com", "password": "password123"}
		body, _ := json.Marshal(payload)
		resp, err := http.Post(baseURL+"/signup", "application/json", bytes.NewReader(body))

		// Then: Should create user successfully
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, resp.StatusCode)
		_ = resp.Body.Close()
	})
}
