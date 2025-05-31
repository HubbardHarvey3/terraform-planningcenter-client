package core

import (
	"net/http"
	"os"
	"strings"
	"testing"
)

const URL = "https://api.planningcenteronline.com/people/v2/people"

func TestPCClient_DoRequest(t *testing.T) {
	// Set environment variables for testing
	app_id := os.Getenv("PC_APP_ID")
	secret_token := os.Getenv("PC_SECRET_TOKEN")

	// Initialize your PC_Client with the mock server URL
	client := NewPCClient(app_id, secret_token, "")

	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	responseBody, err := client.DoRequest(req)
	if err != nil {
		t.Fatalf("Request failed: %v", err)
	}

	// Verify the response
	if strings.Contains(string(responseBody), "404") {
		t.Errorf("Status Code is 404, not 200")
	}
}
