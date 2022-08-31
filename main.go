package main

import (
	"devto-google-fit/controllers"
	"log"
)

func main() {
	log.Println("Server started. Press Ctrl-C to stop server")
	controllers.RunAllRoutes()
}

