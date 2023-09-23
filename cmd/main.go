package main

import (
	"go-app-model/internal/app"
	"log"
	"os"

	"github.com/goioc/di"
)

func main() {
	instance, err := di.GetInstanceSafe("app")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)

		return
	}

	err = instance.(*app.App).Start()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
