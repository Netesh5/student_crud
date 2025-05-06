package response

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Error struct {
	Error  string `json:"error"`
	Status string `json:"status"`
}

const (
	StatusOk    = "ok"
	StatusError = "error"
)

func ToJson(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		return err
	}
	return nil
}

func ErrorResponse(err error) Error {
	return Error{
		Error:  err.Error(),
		Status: StatusError,
	}
}

func ValidationErrorResponse(err validator.ValidationErrors) Error {

	var validationErrors []string

	for _, fieldError := range err {
		switch fieldError.Tag() {
		case "required":
			validationErrors = append(validationErrors, fieldError.Field()+" is required")
		default:
			validationErrors = append(validationErrors, fieldError.Field()+" is invalid")
		}
	}
	return Error{
		Error:  strings.Join(validationErrors, ", "),
		Status: StatusError,
	}
}
