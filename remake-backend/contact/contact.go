package contact

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type MessageFromContactForm struct {
	Name    string `json:"name" validate:"required,min=2"`
	Email   string `json:"email" validate:"required,email"`
	Message string `json:"message" validate:"required,min=10"`
	Subject string `json:"subject" validate:"required,min=5"`
}

func PostMessage(c *fiber.Ctx) error {

	// var body MessageFromContactForm
	body := new(MessageFromContactForm)

	err := c.BodyParser(&body)

	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
		})
	}

	errors := ValidateStruct(*body)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
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
