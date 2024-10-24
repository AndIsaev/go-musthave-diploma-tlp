package application

import (
	"errors"
	c "github.com/AndIsaev/go-musthave-diploma-tlp/cmd/gophermart/configuration"
	"github.com/AndIsaev/go-musthave-diploma-tlp/internal/storage"
	"github.com/AndIsaev/go-musthave-diploma-tlp/internal/storage/postgres"

	"context"
	"log"
	"net/http"
)

type App struct {
	Name   string
	Server *http.Server
	Config *c.Config
	DBConn storage.Storage
}

func NewApp() *App {
	app := &App{Name: "Gophermart"}
	app.Config = c.NewConfig()
	return app
}

// StartApp - start app
func (a *App) StartApp() error {
	conn, err := postgres.NewPostgresStorage(a.Config.DB)
	if err != nil {
		return err
	}
	a.DBConn = conn

	log.Printf("start app - %v", a.Name)
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

// Shutdown - close active connections
func (a *App) Shutdown() {
	ctx, cancel := context.WithCancel(context.Background())

	if err := a.DBConn.Close(ctx); err != nil {
		log.Println(errors.Unwrap(err))
	}

	defer cancel()
}
