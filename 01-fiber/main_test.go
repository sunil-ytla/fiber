package main

import (
	"net/http/httptest"
	"testing"
	"github.com/gofiber/fiber/v2"
)

func TestCheckToken(t *testing.T) {
	app := InitApp()

	// request without token
	req := httptest.NewRequest("GET", "/test-token", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Failed to create test request: %v", err)
	}
	if resp.StatusCode != fiber.StatusUnauthorized {
		t.Errorf("Expected status code %d, got %d", fiber.StatusUnauthorized, resp.StatusCode)
	}

	// request with token
	req = httptest.NewRequest("GET", "/test-token", nil)
	req.Header.Set("Authorization", "Bearer mock-token")
	resp, err = app.Test(req)
	if err != nil {
		t.Fatalf("Failed to create test request: %v", err)
	}
	if resp.StatusCode != fiber.StatusOK {
		t.Errorf("Expected status code %d, got %d", fiber.StatusUnauthorized, resp.StatusCode)
	}

}