package client

import (
	"errors"
	"log"
	"net/url"

	"github.com/Didstopia/motionerd/client/module"
)

// Client is a top-level abstraction for the connection logic
type Client struct {
	Module        module.Module
	ConnectionURL *url.URL
}

// NewClient creates a new instance of Client and returns a reference to it
func NewClient(connectionURL string) (*Client, error) {
	log.Println("Creating a Client with URL:", connectionURL)

	// Parse the connection URL
	url, err := url.Parse(connectionURL)
	if err != nil {
		return nil, err
	}

	// Validate the scheme (eg. http://)
	log.Println("Parsed connection URL scheme:", url.Scheme)
	if url.Scheme == "" {
		return nil, errors.New("missing scheme from connection URL: " + connectionURL)
	}

	// Create a new Client
	c := &Client{}

	// Create the appropriate module based on the scheme
	module, err := module.NewModule(url.Scheme)
	if err != nil {
		return nil, err
	}
	c.Module = module
	log.Println("Client module:", c.Module)

	/*// Create and validate the RTSP client
	c.RTSPClient = &rtspclient.New()
	if c.RTSPClient == nil {
		return errors.New("failed to create RTSP client")
	}*/

	// Store the connection URL
	c.ConnectionURL = url

	// Return the Client, and a nil error on success
	return c, nil
}

// Open the camera connection
func (c *Client) Open() error {
	log.Println("Opening connection with URL:", c.ConnectionURL)

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
	log.Println("Closing connection with URL:", c.ConnectionURL)

	/*// Close the RTSP connection
	c.RTSPClient.Close()*/

	// Return nil on success
	return nil
}
