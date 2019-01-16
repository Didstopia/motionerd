package client

import (
	"errors"

	"github.com/Didstopia/motionerd/client/module"
)

// Client is a top-level abstraction for the connection logic
type Client struct {
	Module        module.Module
	ConnectionURL string
}

// NewClient creates a new instance of Client and returns a reference to it
func NewClient(connectionURL string) (*Client, error) {
	// Validate the connection URL
	if len(connectionURL) < 1 {
		return nil, errors.New("invalid or empty connection URL: " + connectionURL)
	}

	// Create a new Client
	c := &Client{}

	/*// Create and validate the RTSP client
	c.RTSPClient = &rtspclient.New()
	if c.RTSPClient == nil {
		return errors.New("failed to create RTSP client")
	}*/

	// Store the connection string
	c.ConnectionURL = connectionURL

	// Return the Client, and a nil error on success
	return c, nil
}

// Open the camera connection
func (c *Client) Open() error {
	/*// Attempt to connect to the RTSP server
	if c.RTSPClient.DialRTSP() != true {
		return errors.new("failed to open RTSP connection with URL: " + c.RTSPURL)
	}

	// TODO: What does this actually do? How can we access the output?
	// Attempt to send a "DESCRIBE" command to the RTSP server
	if c.RTSPClient.SendRequest() != true {
		return errors.new("failed to send describe command to RTSP server at URL: " + c.RTSPURL)
	}*/

	// Return nil on success
	return nil
}

// Close the camera connection
func (c *Client) Close() error {
	/*// Close the RTSP connection
	c.RTSPClient.Close()*/

	// Return nil on success
	return nil
}
