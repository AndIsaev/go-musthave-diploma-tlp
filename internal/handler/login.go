package handler

import (
	"encoding/json"
	"github.com/AndIsaev/go-musthave-diploma-tlp/internal/model"
	"log"
	"net/http"
)

func (h *Handler) Login() http.HandlerFunc {
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
		existsUser, err := h.UserService.Login(ctx, &params)
		if err != nil {
			h.writeJSONResponseError(w, err, http.StatusUnauthorized)
			return
		}

		response, err := json.Marshal(existsUser)
		if err != nil {
			log.Printf("can't serialize: %v", params)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Authorization", existsUser.Token)
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}
