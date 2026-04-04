package main_test

import (
	"fmt"
	"os"
	"testing"
)

// E2E or integration tests for the main app can go here
// For example, testing server startup or full flows

func TestMain(m *testing.M) {
	exitCode := m.Run()
	if exitCode != 0 {
		fmt.Println("Tests failed with exit code:", exitCode)
	}
	os.Exit(exitCode)
}
