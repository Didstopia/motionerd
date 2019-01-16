package module

// RTSPModule is a connection module that uses RTSP
type RTSPModule struct {
	Dummy bool
}

// This checks at compile time that we're implementing the Module interface
var _ Module = (*RTSPModule)(nil)

// Open the connection
func (m *RTSPModule) Open() error {
	return nil
}

// Close the connection
func (m *RTSPModule) Close() error {
	return nil
}

// TODO: Implement actual logic for connecting to cameras over RTSP
