package controllers

import (
	"go-app-model/internal/services"

	"github.com/gofiber/fiber/v2"
)

type TestController struct {
	TestService services.TestService `di.inject:"TestService"`
}

func (c *TestController) AddRoutes(app *fiber.App) {
	routes := app.Group("/test")

	routes.Get("/:max", c.GetRandomNumber)
}

func (c *TestController) GetRandomNumber(ctx *fiber.Ctx) error {
	context := ctx.Context()
	maxNumber := ctx.Params("max")

	randomNumber, err := c.TestService.GetNumber(context, maxNumber)
	if err != nil {
		return err
	}

	return ctx.Status(200).JSON(randomNumber)
}
