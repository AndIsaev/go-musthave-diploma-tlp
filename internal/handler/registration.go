package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/AndIsaev/go-musthave-diploma-tlp/internal/model"
)

func (h *Handler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		var params model.AuthParams

		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			h.writeJSONResponseError(w, err, http.StatusBadRequest)
			return
		}

		if err := h.Validator.Struct(params); err != nil {
			h.writeJSONValidatorResponse(w, err)
			return
		}

		ctx := r.Context()
		user, err := h.UserService.Register(ctx, &params)
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

		w.Header().Set("Authorization", user.Token)
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}
