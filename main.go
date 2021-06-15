package main

import (
	"fmt"
	"os"

	"github.com/bangarangler/theremake/remake-backend/dotenvConfig"
	"github.com/bangarangler/theremake/remake-backend/notionAPI"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	// app.Use(cors.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "https://gofiber.io, https://gofiber.net",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	notionAPI.InitNotionRoutes(app)

	// port := 5000
	err := app.Listen(dotenvConfig.Port)
	if err != nil {
		fmt.Println("err", err)
		panic(err)
	}
	fmt.Fprintf(os.Stdout, "🚀 Server ready at http://localhost%s\n", dotenvConfig.Port)
}
