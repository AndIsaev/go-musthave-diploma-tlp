package main

import (
	"log"

	"github.com/AndIsaev/go-musthave-diploma-tlp/cmd/gophermart/application"
)

func main() {
	app := application.NewApp()
	err := app.StartApp()
	if err != nil {
		log.Fatalf("close process with error: %s\n", err.Error())
	}
}
