package main

import (
	"fmt"

	"github.com/bangarangler/theremake/remake-backend/dotenvConfig"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())

	PORT := dotenvConfig.Port
	err := app.Listen("localhost:" + PORT)
	if err != nil {
		fmt.Println("err", err)
		panic("Problem Starting App")
	}
	fmt.Printf("🌙 Server Running on %s\n", PORT)

}
