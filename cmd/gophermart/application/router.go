package application

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
)

// initRouter - init router of app
func (a *App) initRouter() {
	r := a.Router
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

}
