package response

import (
	"encoding/json"
	"net/http"
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
