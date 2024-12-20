package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"

	"github.com/AndIsaev/go-musthave-diploma-tlp/internal/service"
)

type ContextKey string

type Handler struct {
	UserService service.Service
	Validator   *validator.Validate
}

type Response struct {
	Message string `json:"message"`
}

type ResponseMulti struct {
	Message map[string]string `json:"message"`
}

func (h *Handler) writeJSONValidatorResponse(w http.ResponseWriter, err error) {
	errors := make(map[string]string)
	for _, err := range err.(validator.ValidationErrors) {
		errors[strings.ToLower(err.Field())] = err.Error()
	}
	w.WriteHeader(http.StatusBadRequest)
	data, _ := json.Marshal(ResponseMulti{Message: errors})
	w.Write(data)
}

func (h *Handler) writeJSONResponseError(w http.ResponseWriter, err error, statusCode int) {
	response, _ := json.Marshal(Response{Message: err.Error()})
	w.WriteHeader(statusCode)
	w.Write(response)
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
