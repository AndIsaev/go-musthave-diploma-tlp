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
			h.writeJSONResponseError(w, exception.InvalidOrderNumber, http.StatusUnprocessableEntity)
			return
		}

		params.Number = number

		if err := h.Validator.Struct(params); err != nil {
			h.writeJSONValidatorResponse(w, err)
			return
		}
		key := model.ContextKey("login")

		login, ok := r.Context().Value(key).(model.UserLogin)
		if !ok {
			response, _ := json.Marshal(Response{Message: "could not find login in context"})
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(response)
			return
		}

		params.UserLogin = login

		val, err := h.UserService.SetOrder(r.Context(), &params)
		if errors.Is(err, exception.OrderAlreadyExistsAnotherUser) {
			h.writeJSONResponseError(w, err, http.StatusConflict)
			return
		}

		if errors.Is(err, exception.OrderAlreadyExists) {
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

func isLuhnValid(orderNumber int) bool {
	orderStr := strconv.Itoa(orderNumber)

	var sum int
	double := false

	for i := len(orderStr) - 1; i >= 0; i-- {
		digit, err := strconv.Atoi(string(orderStr[i]))
		if err != nil {
			return false
		}

		if double {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
		double = !double
	}

	return sum%10 == 0
}
