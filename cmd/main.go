package main

import (
	"fmt"
	"github.com/akashkrao99/go-sample-http/router"
)

func main() {
	fmt.Println("hello")
	router := router.GetRouter()

	// Start the server on port 8080
	router.Run(":8080")
}
