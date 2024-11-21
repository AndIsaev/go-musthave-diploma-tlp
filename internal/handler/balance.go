package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/AndIsaev/go-musthave-diploma-tlp/internal/exception"
	"github.com/AndIsaev/go-musthave-diploma-tlp/internal/model"
)

func (h *Handler) CheckBalance() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		key := ContextKey("login")

		login, ok := r.Context().Value(key).(model.UserLogin)
		if !ok {
			h.writeJSONResponseError(w, errors.New("could not find login in context"), http.StatusInternalServerError)
			return
		}

		balance, err := h.UserService.GetUserBalance(r.Context(), &login)
		if err != nil {
			h.writeJSONResponseError(w, exception.ErrInternalServer, http.StatusInternalServerError)
			return
		}
		response, err := json.Marshal(balance)
		if err != nil {
			log.Printf("can't serialize: %v", balance)
			h.writeJSONResponseError(w, exception.ErrInternalServer, http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func (h *Handler) Withdraw() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		withdraw := &model.Withdraw{}

		defer r.Body.Close()

		key := ContextKey("login")

		if err := json.NewDecoder(r.Body).Decode(&withdraw); err != nil {
			h.writeJSONResponseError(w, err, http.StatusBadRequest)
			return
		}
		val, err := strconv.Atoi(withdraw.Order)
		if err != nil {
			h.writeJSONResponseError(w, exception.ErrInternalServer, http.StatusInternalServerError)
			return
		}
		if !isLuhnValid(val) {
			h.writeJSONResponseError(w, exception.ErrInvalidOrderNumber, http.StatusUnprocessableEntity)
			return
		}

		login, ok := r.Context().Value(key).(model.UserLogin)
		if !ok {
			h.writeJSONResponseError(w, errors.New("could not find login in context"), http.StatusInternalServerError)
			return
		}

		newWithdraw, err := h.UserService.DeductPoints(r.Context(), withdraw, &login)
		if errors.Is(err, exception.ErrNotEnoughBonuses) {
			h.writeJSONResponseError(w, err, http.StatusPaymentRequired)
			return
		} else if err != nil {
			h.writeJSONResponseError(w, exception.ErrInternalServer, http.StatusInternalServerError)
			return
		}

		response, err := json.Marshal(newWithdraw)
		if err != nil {
			log.Printf("can't serialize: %v", newWithdraw)
			h.writeJSONResponseError(w, exception.ErrInternalServer, http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func (h *Handler) GetWithdrawals() http.HandlerFunc {
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

		withdrawals, err := h.UserService.GetUserWithdrawals(r.Context(), &login)

		if err != nil {
			response, _ := json.Marshal(Response{Message: "internal server error"})
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(response)
			return
		}
		if withdrawals == nil {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		response, err := json.Marshal(withdrawals)
		if err != nil {
			log.Printf("can't serialize: %v", withdrawals)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}
