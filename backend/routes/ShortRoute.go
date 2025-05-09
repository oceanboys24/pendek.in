package routes

import (
	"github.com/gofiber/fiber/v2"
	"pendek.in/handlers"
)




func ShortRoutes(route fiber.Router)  {
	route.Get("/", handlers.ShortHandlerGet)
}