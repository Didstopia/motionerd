package client

import "testing"

func TestClient(t *testing.T) {
	// Create a new Client with an invalid connection URL
	invalidClient, err := NewClient("")
	if invalidClient != nil {
		t.Fatal("expected nil client with empty connection URL")
	}
	if err == nil {
		t.Fatal("expected nil error with empty connection URL")
	}

	// Create a new Client
	client, err := NewClient("DUMMY")
	if err != nil {
		t.Fatal(err)
	}

	// Open the connection
	if err := client.Open(); err != nil {
		t.Fatal(err)
	}

	// Close the connection
	if err := client.Close(); err != nil {
		t.Fatal(err)
	}
}
