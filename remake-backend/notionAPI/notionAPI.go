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

	// dbRes := DBRes
	return ctx.Status(fiber.StatusOK).JSON(res.Body)
	// return ctx.Status(fiber.StatusOK).JSON(output)
	// return ctx.JSON(fiber.Map{data: data})
}

// func mapDBRes(dbRes interface{}) interface{} {
// 	return struct {
//
// 	}
// }
type DBRes struct {
	Object  string   `json:"object"`
	Results DBResult `json:"results"`
}

type DBResult struct {
	Object         string     `json:"object"`
	Id             string     `json:"id"`
	CreatedTime    string     `json:"created_time"`
	LastEditedTime string     `json:"last_edited_time"`
	Parent         DBParent   `json:"parent"`
	Archived       bool       `json:"archived"`
	Properties     Properties `json:"properties"`
}

type DBParent struct {
	Type       string `json:"type"`
	DatabaseId string `json:"database_id"`
}

type Properties struct {
	Date        Date        `json:"Date"`
	BlogTags    BlogTags    `json:"Blog Tags"`
	IsPublished IsPublished `json:"isPublished"`
	Author      Author      `json:"Author"`
	BlogTitle   BlogTitle   `json:"Blog Title"`
}

type Date struct {
	Id   string `json:"id"`
	Type string `json:"type"`
	Date IDate  `json:"date"`
}

type IDate struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type BlogTags struct {
	Id          string        `json:"id"`
	Type        string        `json:"type"`
	MultiSelect []MultiSelect `json:"multi_select"`
}

type MultiSelect struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type IsPublished struct {
	Id     string `json:"id"`
	Type   string `json:"type"`
	Select Select `json:"select"`
}

type Select struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type Author struct {
	Id     string   `json:"id"`
	Type   string   `json:"type"`
	People []People `json:"people"`
}

type People struct {
	Object    string `json:"object"`
	Id        string `json:"id"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url"`
	Type      string `json:"type"`
	Person    Person `json:"person"`
}

type Person struct {
	Email string `json:"email"`
}

type BlogTitle struct {
	Id    string  `json:"id"`
	Type  string  `json:"type"`
	Title []Title `json:"title"`
}

type Title struct {
	Type        string      `json:"type"`
	Text        Text        `json:"text"`
	Annotations Annotations `json:"annotations"`
	PlainText   string      `json:"plain_text"`
	href        string      `json:"href"`
}

type Text struct {
	Content string `json:"content"`
	Link    string `json:"link"`
}

type Annotations struct {
	Bold          bool   `json:"bold"`
	Italic        bool   `json:"italic"`
	Strikethrough bool   `json:"strikethrough"`
	Underline     bool   `json:"underline"`
	Code          bool   `json:"code"`
	Color         string `json:"color"`
}
