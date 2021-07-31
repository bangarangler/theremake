package contact

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type MessageFromContactForm struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
	Subject string `json:"subject"`
}

func PostMessage(c *fiber.Ctx) error {

	var body MessageFromContactForm

	fmt.Println(body)
	fmt.Println(c.Request())
	err := c.BodyParser(&body)

	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
		})
	}

	newMsg := &MessageFromContactForm{
		Name:    body.Name,
		Email:   body.Email,
		Message: body.Message,
		Subject: body.Subject,
	}

	fmt.Println(newMsg)

	err = SendEmailFromContact(*newMsg)
	if err != nil {
		fmt.Println("err Sending Email...", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Issue Sending Email... Please try again; Or reach out on Slack ; )",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "We Sent Your Message!",
	})
}
