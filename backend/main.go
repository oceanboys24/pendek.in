package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"pendek.in/config"
	"pendek.in/routes"
)




func main() {
	// Load Env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Cannot Load ENV")
	}
	// Connect Database
	config.ConnectDatabaseNeon()

	// Init Fiber 
	app := fiber.New()

	// Short Link Generator Endpoint
	routes.ShortRoutes(app)

	// Run Services
	app.Listen("0.0.0.0:3001")

}