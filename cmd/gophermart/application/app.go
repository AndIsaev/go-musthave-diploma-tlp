package application

import (
	"errors"

	c "github.com/AndIsaev/go-musthave-diploma-tlp/cmd/gophermart/configuration"
	"github.com/AndIsaev/go-musthave-diploma-tlp/internal/storage"
	"github.com/AndIsaev/go-musthave-diploma-tlp/internal/storage/postgres"
	"github.com/go-chi/chi/v5"

	"context"
	"log"
	"net/http"
)

type App struct {
	Name   string
	Server *http.Server
	Config *c.Config
	DBConn storage.Storage
	Router chi.Router
}

func NewApp() *App {
	app := &App{Name: "Gophermart"}
	app.Config = c.NewConfig()
	app.Router = chi.NewRouter()
	return app
}

// StartApp - start app
func (a *App) StartApp() error {
	conn, err := postgres.NewPostgresStorage(a.Config.DB)
	if err != nil {
		return err
	}
	a.DBConn = conn
	err = a.upMigrations()
	if err != nil {
		return err
	}

	a.initRouter()

	log.Printf("start app - %v", a.Name)
	return a.startHTTPServer()
}

// initHTTPServer - init http server
func (a *App) initHTTPServer() {
	server := &http.Server{}
	server.Addr = a.Config.Address
	server.Handler = a.Router
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

	if err := a.DBConn.System().Close(ctx); err != nil {
		log.Println(errors.Unwrap(err))
	}

	defer cancel()
}

// upMigrations - run migrations of db
func (a *App) upMigrations() error {
	if err := a.DBConn.System().RunMigrations(context.Background()); err != nil {
		return err
	}
	return nil

}
