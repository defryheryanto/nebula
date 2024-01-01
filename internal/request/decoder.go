package request

import (
	"encoding/json"
	"io"
	"net/http"

	handlederror "github.com/defryheryanto/nebula/internal/errors"
)

func DecodeBody(r *http.Request, payload any) error {
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		if err == io.EOF {
			return handlederror.ErrEmptyRequestBody
		}
		return handlederror.ErrInvalidRequestBody
	}

	return nil
}
