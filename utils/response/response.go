package response

import "github.com/gofiber/fiber/v2"

func SuccessResponse(c *fiber.Ctx, statusCode int, message string, data interface{}) error {
	response := fiber.Map{
		"data":    data,
		"message": message,
	}

	return c.Status(statusCode).JSON(response)
}

func ErrorResponse(c *fiber.Ctx, statusCode int, message string) error {
	response := fiber.Map{
		"message": message,
	}
	return c.Status(statusCode).JSON(response)
}
