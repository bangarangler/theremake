package rest

import (
	"github.com/bangarangler/theremake/remake-backend/contact"
	"github.com/gofiber/fiber/v2"
)

func PostMessageRoute(route fiber.Router) {
	route.Post("/", contact.PostMessage)
}
