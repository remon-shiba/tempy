package controller

import (
	"MondTemplate/app/model"

	"github.com/gofiber/fiber/v2"
)

func HealthCheck(c *fiber.Ctx) error {
	return c.JSON(model.ResponseBody{
		Status:  100,
		Message: "API is running",
	})
}
