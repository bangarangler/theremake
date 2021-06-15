package notionAPI

import (
	"github.com/gofiber/fiber"
)

func InitNotionRoutes(grp fiber.Router) {
	notionRoutes := grp.Group("/notion")

	notionRoutes.Get("/db", func(ctx *fiber.Ctx) {
		ctx.Send("testing /db route")
	})
}
