package module

// HTTPModule is a connection module that uses HTTP(S)
type HTTPModule struct {
	HTTPS bool
}

// This checks at compile time that we're implementing the Module interface
var _ Module = (*HTTPModule)(nil)

// Open the connection
func (m *HTTPModule) Open() error {
	return nil
}

// Close the connection
func (m *HTTPModule) Close() error {
	return nil
}

// TODO: Implement actual logic for connecting to cameras over HTTP(S)
