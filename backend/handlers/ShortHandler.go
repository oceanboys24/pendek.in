package handlers

import "github.com/gofiber/fiber/v2"




func ShortHandlerGet(c *fiber.Ctx) error  {
	return c.SendString("Testing Get")
}