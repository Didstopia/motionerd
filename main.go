package main

import (
	"os"

	"github.com/Didstopia/motionerd/client"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// Create a new client
	client, err := client.NewClient(os.Getenv("CONNECTION_URL"))
	if err != nil {
		panic(err)
	}

	// Open the connection
	if err := client.Open(); err != nil {
		panic(err)
	}

	// Close the connection
	if err := client.Close(); err != nil {
		panic(err)
	}
}
