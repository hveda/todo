package routes

import (
	"github.com/gofiber/fiber/v2"
)

//Hello
func Hello(c *fiber.Ctx) error {
	return c.SendString("fiber")
}
