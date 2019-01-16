package module

import (
	"errors"
	"testing"
)

func TestNewModule(t *testing.T) {
	// Try to create invalid modules
	if _, err := NewModule(""); err == nil {
		t.Fatal("expected nil error with empty URL scheme")
	}
	if _, err := NewModule("DUMMY"); err == nil {
		t.Fatal("expected nil error with invalid URL scheme")
	}

	// TODO: Aren't these unnecessary, considering we should test each
	//       module independently, in their own test files?
	// Try to create an HTTP module
	httpModule, err := NewModule("http")
	if err != nil {
		t.Fatal(err)
	}
	if httpModule == nil {
		t.Fatal(errors.New("expected httpModule to not be nil"))
	}

	// Try to create an HTTPS module
	httpsModule, err := NewModule("https")
	if err != nil {
		t.Fatal(err)
	}
	if httpsModule == nil {
		t.Fatal(errors.New("expected httpsModule to not be nil"))
	}

	// Try to create an RTSP module
	rtspModule, err := NewModule("rtsp")
	if err != nil {
		t.Fatal(err)
	}
	if rtspModule == nil {
		t.Fatal(errors.New("expected rtspModule to not be nil"))
	}
}
