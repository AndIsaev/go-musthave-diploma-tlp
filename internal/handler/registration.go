package handler

import (
	"encoding/json"
	"github.com/AndIsaev/go-musthave-diploma-tlp/internal/model"
	"github.com/AndIsaev/go-musthave-diploma-tlp/internal/service"
	"log"
	"net/http"
)

func (h *Handler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		defer r.Body.Close()

		var candidate model.User

		if err := json.NewDecoder(r.Body).Decode(&candidate); err != nil {
			h.writeJSONResponseError(w, err, http.StatusBadRequest)
			return
		}

		if err := h.Validator.Struct(candidate); err != nil {
			h.writeJSONValidatorResponse(w, err)
			return
		}

		ctx := r.Context()
		user, err := h.UserService.Register(ctx, &service.RegisterParams{Login: candidate.Login, Password: candidate.Password})
		if err != nil {
			h.writeJSONResponseError(w, err, http.StatusConflict)
			return
		}

		response, err := json.Marshal(user)
		if err != nil {
			log.Printf("can't serialize: %v", user)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}
