package handler

import (
	"encoding/json"
	"errors"
	"github.com/AndIsaev/go-musthave-diploma-tlp/internal/exception"
	"github.com/AndIsaev/go-musthave-diploma-tlp/internal/model"
	"log"
	"net/http"
)

func (h *Handler) GetBalance() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		defer r.Body.Close()

		key := ContextKey("login")

		login, ok := r.Context().Value(key).(model.UserLogin)
		if !ok {
			h.writeJSONResponseError(w, errors.New("could not find login in context"), http.StatusInternalServerError)
			return
		}

		orders, err := h.UserService.GetUserOrders(r.Context(), &login)
		if err != nil {
			h.writeJSONResponseError(w, exception.ErrInternalServer, http.StatusInternalServerError)
			return
		}
		if orders == nil {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		response, err := json.Marshal(orders)
		if err != nil {
			log.Printf("can't serialize: %v", orders)
			h.writeJSONResponseError(w, exception.ErrInternalServer, http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}
