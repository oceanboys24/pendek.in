package handlers

import (
	"os"

	"github.com/go-playground/validator/v10"
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
	validate := validator.New()

	if err := c.BodyParser(&bodyRequest); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message" : "Cannot Parse Body" + err.Error(),
		})
	}

	err := validate.Var(bodyRequest.Url, "required,url")
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"message" : "Error Validate" + err.Error(),
		})
	}
	
	result , err := services.CreateShortService(bodyRequest)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Message " : "Error Create Short Link" + err.Error(),
		})
	}
	domain := os.Getenv("DOMAIN")

	resultFinal := domain + result.ShortUrl

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message" : "Success Create Short Link",
		"short_link" : resultFinal,
	})

}