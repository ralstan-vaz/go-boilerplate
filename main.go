package main

import (
	"github.com/ralstan-vaz/go-boilerplate/initiate"
)

func main() {
	// Initialize the app
	err := initiate.Initialize()
	if err != nil {
		panic(err)
	}
}
