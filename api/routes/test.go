package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Test(c *fiber.Ctx) error {
	fmt.Println("shorten")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"res": "good"})
}
