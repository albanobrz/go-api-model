package app

import (
	"fmt"
	"go-app-model/internal/http"
	"html"
	"os"
	"strings"
)

type App struct {
	router http.Router `di.inject:"Routers"`
}

func (app *App) bootstrapDb() error {
	connUri := strings.TrimRight(html.UnescapeString(os.Getenv("DB_URI")), "\r\n")
	fmt.Println(connUri)
	fmt.Println("test1")
	// connect to db here

	return nil
}

func (app *App) Start() error {
	err := app.bootstrapDb()
	if err != nil {
		return err
	}

	fmt.Println("test")
	return app.router.Start()
}
