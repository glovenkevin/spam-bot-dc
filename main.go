package main

import (
	"discord-spam-bot/application"
	"log"
)

func main() {
	app, err := application.Setup()
	if err != nil {
		log.Fatal(err)
	}

	err = app.Run()
	if err != nil {
		log.Fatal(err)
	}
}
