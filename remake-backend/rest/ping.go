package rest

import (
	"github.com/gofiber/fiber/v2"
)

func PingRoute(route fiber.Router) {
	route.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "pong",
		})
	})
}
