package module

// Module defines the connection type and logic
type Module struct {
	ConnectionType string
}

// TODO: Allow easy creation of specific module types, like HTTPModule
// TODO: Potentially parse the connection string and determine the module automatically
// TODO: Provide reusable functions for connecting, disconnecting and receiving data (raw bytes?)
