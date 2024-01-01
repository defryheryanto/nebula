package response

import (
	"encoding/json"
	"net/http"

	handlederror "github.com/defryheryanto/nebula/internal/errors"
)

type SuccessJSONResponse[T any] struct {
	Success bool `json:"success"`
	Data    T    `json:"data"`
}

func SuccessJSON[T any](w http.ResponseWriter, status int, data T) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(SuccessJSONResponse[T]{
		Success: true,
		Data:    data,
	})
}

type FailedJSONReponse struct {
	Success bool   `json:"success"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func FailedJSON(w http.ResponseWriter, r *http.Request, err error) {
	handledError := handlederror.Extract(err)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(handledError.HttpStatus)

	json.NewEncoder(w).Encode(FailedJSONReponse{
		Success: false,
		Code:    handledError.Code,
		Message: handledError.Message,
	})
}
