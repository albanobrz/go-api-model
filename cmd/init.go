package main

import (
	"log"
	"os"
	"reflect"

	"go-app-model/internal/app"
	"go-app-model/internal/http"

	"github.com/gofiber/fiber/v2"
	"github.com/goioc/di"
	_ "github.com/joho/godotenv/autoload"
	"github.com/nleeper/goment"
)

func init() {
	configureFiber()
	registerBeans()

	err := di.InitializeContainer()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func configureFiber() {
	fiber.SetParserDecoder(fiber.ParserConfig{
		IgnoreUnknownKeys: true,
		ParserType: []fiber.ParserType{
			{
				Customtype: goment.Goment{},
				Converter: func(value string) reflect.Value {
					if value == "" {
						return reflect.Value{}
					}

					if v, err := goment.New(value, "YYYY-MM-DDTHH:mm:ssZ"); err == nil {
						return reflect.ValueOf(*v)
					}

					if v, err := goment.New(value, "YYYY-MM-DD"); err == nil {
						return reflect.ValueOf(*v)
					}

					return reflect.Value{}
				},
			},
		},
		ZeroEmpty: true,
	})
}

func registerBeans() {
	// Router
	di.RegisterBean("Routers", reflect.TypeOf((*http.DefaultRouter)(nil)))

	// Controllers

	// Application
	di.RegisterBean("app", reflect.TypeOf((*app.App)(nil)))
}
