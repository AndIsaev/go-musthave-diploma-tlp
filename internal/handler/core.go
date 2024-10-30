package handler

import (
	"encoding/json"
	"github.com/AndIsaev/go-musthave-diploma-tlp/internal/service"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
)

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
