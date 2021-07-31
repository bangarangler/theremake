package contact

import (
	"net/http/httptest"
	"testing"

	_ "github.com/bangarangler/theremake/remake-backend/dotenvConfig"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

// go test -run -v Test_Handler
func Test_Ping(t *testing.T) {
	app := fiber.New()

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})

	resp, err := app.Test(httptest.NewRequest("GET", "/ping", nil))

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, 200, resp.StatusCode, "Status code")
}
