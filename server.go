// main.go

package main

import (
	"books/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {

	// Set the router as the default one provided by Gin
	router = gin.Default()

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/*")

	config, err := utils.LoadConfig(".")

	if err != nil {
		fmt.Println(err)
	}

	// Initialize the routes
	initializeRoutes()

	// Start serving the application
	router.Run(config.SRVPort)

}
