// main.go
package main

import (
	"golang_learning/config"
	"golang_learning/ioc"
	"log"

	"github.com/gin-gonic/gin"
)

var (
	server *gin.Engine
)

func init() {
	// Initialize the database connection
	config.ConnectDatabase()

	// Initialize IoC
	ioc.Initialize()

	// Initialize the server
	server = gin.Default()
}

func main() {
	defer config.DisconnectDatabase()

	// Register routes
	config.RegisterRoutes(server, ioc.UserController)

	log.Fatal(server.Run(":9090"))
}
