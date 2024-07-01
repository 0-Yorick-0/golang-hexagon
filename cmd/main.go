package main

import (
	"golang-hexagon/internal/app/infrastructure/config"
	"log"
)

func main() {
	cfg, err := config.NewAppConfig("./configs/config.json")
	if err != nil {
		log.Fatalln("Unable to load config file", err)
	}

	application := app.NewApp(cfg)
	application.Run()
}
