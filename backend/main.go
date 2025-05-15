package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"pendek.in/config"
	"pendek.in/routes"
)




func main() {
	// Connect Database
	config.ConnectDatabaseNeon()

	// Init Fiber 
	app := fiber.New()

	// Use Cors
	app.Use(cors.New())

	
	// Short Link Generator Endpoint
	routes.ShortRoutes(app)

	// Run Services
	app.Listen("0.0.0.0:3001")

}