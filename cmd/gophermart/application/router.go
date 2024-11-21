package application

import (
	"net/http"

	mid "github.com/AndIsaev/go-musthave-diploma-tlp/internal/handler/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// initRouter - init router of app
func (a *App) initRouter() {
	r := a.Router

	r.Use(middleware.Logger, middleware.StripSlashes, middleware.CleanPath)
	r.Use(mid.JSONMiddleware)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	r.Route("/api/user", func(r chi.Router) {
		r.Post("/register", a.Handler.Register()) // POST
		r.Post("/login", a.Handler.Login())       // POST

		r.Group(func(r chi.Router) {
			r.Use(mid.JwtAuthMiddleware)
			r.Post("/orders", a.Handler.SetOrder())
			r.Get("/orders", a.Handler.ListUserOrders())
			r.Get("/balance", a.Handler.CheckBalance())
			r.Get("/withdrawals", a.Handler.GetWithdrawals())
			r.Post("/balance/withdraw", a.Handler.Withdraw())
		})
	})
}
