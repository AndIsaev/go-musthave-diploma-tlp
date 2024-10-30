package application

import (
	"github.com/go-chi/chi/v5"
	"net/http"

	mid "github.com/AndIsaev/go-musthave-diploma-tlp/internal/handler/middleware"
	// "github.com/AndIsaev/go-musthave-diploma-tlp/internal/service"

	"github.com/go-chi/chi/v5/middleware"
)

// initRouter - init router of app
func (a *App) initRouter() {
	r := a.Router

	r.Use(middleware.Logger, middleware.StripSlashes, middleware.CleanPath)
	r.Use(mid.JsonMiddleware)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	r.Route("/api/user", func(r chi.Router) {
		r.Post("/register", a.Handler.Register()) // POST

	})

}
