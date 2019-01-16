package client

import "testing"

func TestClient(t *testing.T) {
	// Try creating invalid clients
	if _, err := NewClient(""); err == nil {
		t.Fatal("expected nil error with empty connection URL")
	}
	if _, err := NewClient("DUMMY"); err == nil {
		t.Fatal("expected nil error with empty connection URL")
	}
	if _, err := NewClient("WWW.DUMMY.COM"); err == nil {
		t.Fatal("expected nil error with empty connection URL")
	}

	// Try creating a valid client
	client, err := NewClient("http://localhost:3000/videostream.cgi")
	if err != nil {
		t.Fatal(err)
	}

	// Try opening the client connection
	if err := client.Open(); err != nil {
		t.Fatal(err)
	}

	// Try closing the client connection
	if err := client.Close(); err != nil {
		t.Fatal(err)
	}
}
