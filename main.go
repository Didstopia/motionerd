package main

import "fmt"

// Version is set automatically by GoReleaser
var version = "development"

func main() {
	fmt.Println("Version:", version)
}
