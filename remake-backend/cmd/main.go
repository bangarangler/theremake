package main

import (
	"fmt"
	"os"

	"github.com/bangarangler/theremake/remake-backend/dotenvConfig"
	"github.com/bangarangler/theremake/remake-backend/notionAPI"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	// "github.com/gofiber/fiber/v2/middleware/cors"
	// "github.com/gofiber/fiber/v2/middleware/cors"
	// "github.com/gofiber/fiber/v2/middleware/cors"
	// "github.com/gofiber/fiber/middleware/cors"
	// "github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(middleware.Logger())
	// app.Use(middleware.cors.New(cors.New(cors.Config{
	// 	AllowOrigins: "http://localhost:5000",
	// 	AllowHeaders: "Origin, Content-Type, Accept",
	// })))

	notionAPI.InitNotionRoutes(app)

	// port := 5000
	err := app.Listen(dotenvConfig.PORT)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stdout, "🚀 Server ready at http://localhost%s\n", dotenvConfig.PORT)
}
