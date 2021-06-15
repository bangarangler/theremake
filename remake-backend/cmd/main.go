package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
)

func main() {
	app := fiber.New()
	app.Use(middleware.Logger())

	port := 5000
	err := app.Listen(port)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stdout, "🚀 Server ready at http://localhost%s\n", port)
}
