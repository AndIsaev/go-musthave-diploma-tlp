package handler

import (
	"encoding/json"
	"errors"
	"github.com/AndIsaev/go-musthave-diploma-tlp/internal/exception"
	"github.com/AndIsaev/go-musthave-diploma-tlp/internal/model"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) SetOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		defer r.Body.Close()

		params := model.UserOrder{}
		var number int

		if err := json.NewDecoder(r.Body).Decode(&number); err != nil {
			h.writeJSONResponseError(w, err, http.StatusBadRequest)
			return
		}

		if !isLuhnValid(number) {
			h.writeJSONResponseError(w, exception.ErrInvalidOrderNumber, http.StatusUnprocessableEntity)
			return
		}

		params.Number = strconv.Itoa(number)

		if err := h.Validator.Struct(params); err != nil {
			h.writeJSONValidatorResponse(w, err)
			return
		}
		key := ContextKey("login")

		login, ok := r.Context().Value(key).(model.UserLogin)
		if !ok {
			response, _ := json.Marshal(Response{Message: "could not find login in context"})
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(response)
			return
		}

		params.UserLogin = login

		val, err := h.UserService.SetOrder(r.Context(), &params)
		if errors.Is(err, exception.ErrOrderAlreadyExistsAnotherUser) {
			h.writeJSONResponseError(w, err, http.StatusConflict)
			return
		}

		if errors.Is(err, exception.ErrOrderAlreadyExists) {
			h.writeJSONResponseError(w, err, http.StatusOK)
			return
		}

		response, err := json.Marshal(val)
		if err != nil {
			log.Printf("can't serialize: %v", params)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusAccepted)
		w.Write(response)
	}
}

func (h *Handler) ListUserOrders() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		defer r.Body.Close()

		key := ContextKey("login")

		login, ok := r.Context().Value(key).(model.UserLogin)
		if !ok {
			response, _ := json.Marshal(Response{Message: "could not find login in context"})
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(response)
			return
		}

		orders, err := h.UserService.GetUserOrders(r.Context(), &login)
		if err != nil {
			response, _ := json.Marshal(Response{Message: "internal server error"})
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(response)
			return
		}
		if orders == nil {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		response, err := json.Marshal(orders)
		if err != nil {
			log.Printf("can't serialize: %v", orders)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}
