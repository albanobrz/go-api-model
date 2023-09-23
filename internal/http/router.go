package http

import (
	"github.com/bytedance/sonic"
	"github.com/gobuffalo/validate"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"go-app-model/internal/http/controllers"
)

type Router interface {
	Start() error
}

type DefaultRouter struct {
	controllers []controllers.Controller `di.inject:"Controllers"`
}

func (r *DefaultRouter) Start() error {
	app := fiber.New(fiber.Config{
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			statusCode := fiber.StatusInternalServerError
			var body interface{}

			body = &fiber.Map{"message": err.Error()}

			if fiberErr, ok := err.(*fiber.Error); ok {
				statusCode = fiberErr.Code
				body = &fiber.Map{"message": err.Error()}
			}

			if validationErrs, ok := err.(*validate.Errors); ok {
				statusCode = fiber.StatusBadRequest
				body = validationErrs
			}

			if _, ok := err.(*fiber.SyntaxError); ok {
				statusCode = fiber.StatusUnprocessableEntity
				body = &fiber.Map{"message": err.Error()}
			}

			return ctx.Status(statusCode).JSON(body)
		},
	})

	app.Use(cors.New())

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	// add controllers
	for _, controller := range r.controllers {
		controller.AddRoutes(app)
	}

	err := app.Listen(":3003")

	return err
}
