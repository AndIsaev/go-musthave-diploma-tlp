package application

import (
	c "github.com/AndIsaev/go-musthave-diploma-tlp/cmd/gophermart/configuration"
	// "github.com/AndIsaev/go-musthave-diploma-tlp/internal/storage"

	"log"
	"net/http"
)

type App struct {
	Name   string
	Server *http.Server
	Config *c.Config
}

func NewApp() *App {
	app := &App{Name: "Gophermart"}
	app.Config = c.NewConfig()
	return app
}

func (a *App) StartApp() error {
	return a.startHTTPServer()
}

// initHTTPServer - init http server
func (a *App) initHTTPServer() {
	server := &http.Server{}
	server.Addr = a.Config.Address
	a.Server = server
}

// startHTTPServer - start http server
func (a *App) startHTTPServer() error {
	a.initHTTPServer()
	log.Printf("start server on: %s\n", a.Config.Address)
	return a.Server.ListenAndServe()
}
