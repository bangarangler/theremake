package notionAPI

import (
	"github.com/gofiber/fiber"
)

func InitNotionRoutes() {
	notionRoutes := fiber.Router.Group("/notion")

	notionRoutes.Get("/db", func(ctx *fiber.Ctx) {
		ctx.Send("testing /db route")
	})
}
