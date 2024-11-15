package main

import (
	"github.com/AndIsaev/go-musthave-diploma-tlp/cmd/gophermart/application"
	"log"
)

func main() {
	app := application.NewApp()
	err := app.StartApp()
	if err != nil {
		log.Fatalf("close process with error: %s\n", err.Error())
	}
}
