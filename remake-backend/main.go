package main

import (
	"fmt"

	"github.com/bangarangler/theremake/remake-backend/dotenvConfig"
	"github.com/bangarangler/theremake/remake-backend/rest"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func ping(ctx *fiber.Ctx) {
	ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "ping 🌑",
	})
}

func setupRoutes(app *fiber.App) {
	rest.PingRoute(app.Group("/ping"))
	rest.PostMessageRoute(app.Group("/api"))

}

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000, https://jonathandain.dev",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	setupRoutes(app)

	PORT := dotenvConfig.Port
	// err := app.Listen("localhost:" + PORT)
	err := app.Listen("jonathandain.dev:" + PORT)
	if err != nil {
		fmt.Println("err", err)
		panic(err)
	}
	fmt.Printf("🌙 Server Running on %s\n", PORT)

}
