package main

import (
	"dev.azure.com/lucasfinetti/Finetti/_git/Gopportunities/config"
	"dev.azure.com/lucasfinetti/Finetti/_git/Gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// Initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	// Initialize Router
	router.Initialize()
}
