package application

import (
	c "github.com/AndIsaev/go-musthave-diploma-tlp/cmd/gophermart/configuration"
	"log"
	"net/http"
)

type App struct {
	Name   string
	config *c.Config
	Server *http.Server
}

func NewApp() *App {
	app := &App{Name: "Gophermart"}
	app.config = c.NewConfig()
	return app
}

func (a *App) StartApp() error {
	return a.startHTTPServer()
}

// initHTTPServer - init http server
func (a *App) initHTTPServer() {
	server := &http.Server{}
	server.Addr = a.config.Address
	a.Server = server
}

// startHTTPServer - start http server
func (a *App) startHTTPServer() error {
	a.initHTTPServer()
	log.Printf("start server on: %s\n", a.config.Address)
	return a.Server.ListenAndServe()
}
