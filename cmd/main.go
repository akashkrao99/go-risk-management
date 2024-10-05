package main

import (
	"github.com/akashkrao99/go-sample-http/config"
	"github.com/akashkrao99/go-sample-http/router"
)

func main() {
	config.InitializeConfig()

	router := router.GetRouter()
	
	router.Run(config.GetConfig().HttpServerConfig.Port)
}
