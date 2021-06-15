package notionAPI

import (
	"fmt"
	"net/http"

	"github.com/bangarangler/theremake/remake-backend/dotenvConfig"
	"github.com/gofiber/fiber"
)

func InitNotionRoutes(app *fiber.App) {
	notionRoutes := app.Group("/notion")

	notionRoutes.Get("/db", getDB)
}

func getDB(ctx *fiber.Ctx) {
	req, err := http.NewRequest("POST", "https://api.notion.com/v1/databases/229f3080be484e9cbd9b649c2c045ec5/query", nil)
	if err != nil {
		// handle err
		fmt.Println("err", err)
	}
	// req.Header.Set("Authorization", os.ExpandEnv("Bearer $(cat /Users/jonathanpalacio/Desktop/theremake/remake-backend/curl/token.txt)"))
	authorization := fmt.Sprintf("Authorization: Bearer {dotenvConfig.NotionAPIKey}")
	fmt.Println("authorization", authorization)
	req.Header.Set("Authorization", "Bearer $(dotenvConfig.NotionAPIKey)")
	req.Header.Set("Notion-Version", "2021-05-13")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	fmt.Println("resp", resp)
	fmt.Println("NotionAPIKey", dotenvConfig.NotionAPIKey)
	if err != nil {
		// handle err
		fmt.Println("err", err)
	}
	// var prettyJSON bytes.Buffer
	// error := json.Indent(&prettyJSON, resp, "", "\t")
	// if error != nil {
	// 	fmt.Println("JSON parse error: ", error)
	// 	return
	// }
	//
	// fmt.Println("Pretty JSON", string(prettyJSON.Bytes()))
	defer resp.Body.Close()
}
