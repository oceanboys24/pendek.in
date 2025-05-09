package handlers

import (
	"github.com/gofiber/fiber/v2"
	"pendek.in/models"
	"pendek.in/services"
)




func ShortHandlerGetId(c *fiber.Ctx) error  {
	id := c.Params("id")

	isExist, err := services.GetShortById(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message" : "Can't Find Data" + err.Error(), 
		})
	}

	return c.Redirect(isExist.OriginUrl)
}

func ShortHandlerCreate(c *fiber.Ctx) error  {
	var bodyRequest models.ShortModelPayload

	if err := c.BodyParser(&bodyRequest); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message" : "Cannot Parse Body" + err.Error(),
		})
	}
	
	result , err := services.CreateShortService(bodyRequest)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Message " : "Error Create Short Link" + err.Error(),
		})
	}

	resultFinal := "http://localhost:3001/" + result.ShortUrl

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message" : "Success Create Short Link",
		"short_link" : resultFinal,
	})

}