package notionAPI

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/bangarangler/theremake/remake-backend/dotenvConfig"
	"github.com/gofiber/fiber/v2"
)

// type Handlers struct {
// }
//
// func NewHandlers() *Handlers {
// 	return &Handlers{}
// }

func InitNotionRoutes(app *fiber.App) {
	notionRoutes := app.Group("/notion")
	// handlers := NewHandlers()

	notionRoutes.Get("/db", getDB)
	// notionRoutes.Get("/test", test)
}

// func test(ctx *fiber.Ctx) error {
// 	ctx.SendString("Hello THERE")
// }

func getDB(ctx *fiber.Ctx) error {
	authorization := fmt.Sprintf("Bearer %s", dotenvConfig.NotionAPIKey)
	requestBody := strings.NewReader(`{}`)
	req, err := http.NewRequest("POST", "https://api.notion.com/v1/databases/229f3080be484e9cbd9b649c2c045ec5/query", requestBody)
	req.Header.Set("Authorization", authorization)
	req.Header.Set("Notion-Version", "2021-05-13")
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("err", err)
		log.Fatal(err)
	}

	data, _ := ioutil.ReadAll(res.Body)
	defer req.Body.Close()
	defer res.Body.Close()
	fmt.Printf("%s\n", data)
	fmt.Printf("%s\n", data[0])
	if err != nil {
		// handle err
		fmt.Println("err", err)
	}
	return ctx.Status(fiber.StatusOK).JSON(res.Body)
	// return ctx.JSON(fiber.Map{data: data})
}

// func mapDBRes(dbRes interface{}) interface{} {
// 	return struct {
//
// 	}
// }
