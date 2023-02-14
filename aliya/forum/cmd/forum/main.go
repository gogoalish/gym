package main

import (
	"log"

	"project/config"
	"project/internal/app"
)

func main() {
	if err := config.ReadConfig(); err != nil {
		log.Fatalln(err)
	}

	if err := app.Run(); err != nil {
		log.Fatalln(err)
	}
}
