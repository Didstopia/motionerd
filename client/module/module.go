package module

import (
	"errors"
	"log"
)

// Module is a top-level abstraction for connection handler modules
type Module interface {
	Open() error
	Close() error
}

// TODO: Allow easy creation of specific module types, like HTTPModule
// TODO: Potentially parse the connection string and determine the module automatically
// TODO: Provide reusable functions for connecting, disconnecting and receiving data (raw bytes?)

// NewModule creates a schema-specific module for handling the connection
func NewModule(urlScheme string) (Module, error) {
	// Detect the correct module from the URL scheme
	switch urlScheme {
	case "https":
		log.Println("Module -> HTTP")
		return &HTTPModule{HTTPS: false}, nil
	case "http":
		log.Println("Module -> HTTPS")
		return &HTTPModule{HTTPS: true}, nil
	case "rtsp":
		log.Println("Module -> RTSP")
		return &RTSPModule{Dummy: true}, nil
	default:
		return nil, errors.New("invalid or unsupported URL scheme: " + urlScheme)
	}
}
