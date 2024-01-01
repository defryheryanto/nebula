package handlederror_test

import (
	"net/http"
	"testing"

	handlederror "github.com/defryheryanto/nebula/internal/errors"
	"github.com/stretchr/testify/assert"
)

func TestErrors(t *testing.T) {
	t.Parallel()
	t.Run("Internal Server Error", func(t *testing.T) {
		err := handlederror.InternalServerError("Detail error")

		assert.Equal(t, handlederror.InternalServerErrorCode, err.Code)
		assert.Equal(t, http.StatusInternalServerError, err.HttpStatus)
		assert.Equal(t, handlederror.DefaultErrorMessage, err.Message)
		assert.Equal(t, "Detail error", err.Detail)
	})

	t.Run("Validation Error", func(t *testing.T) {
		err := handlederror.ValidationError("Detail error")

		assert.Equal(t, handlederror.ValidationErrorCode, err.Code)
		assert.Equal(t, http.StatusBadRequest, err.HttpStatus)
		assert.Equal(t, handlederror.DefaultErrorMessage, err.Message)
		assert.Equal(t, "Detail error", err.Detail)
	})

	t.Run("Unauthorized Error", func(t *testing.T) {
		err := handlederror.UnauthorizedError("Detail error")

		assert.Equal(t, handlederror.UnauthorizedErrorCode, err.Code)
		assert.Equal(t, http.StatusUnauthorized, err.HttpStatus)
		assert.Equal(t, handlederror.DefaultErrorMessage, err.Message)
		assert.Equal(t, "Detail error", err.Detail)
	})

	t.Run("Not Found Error", func(t *testing.T) {
		err := handlederror.NotFoundError("Detail error")

		assert.Equal(t, handlederror.NotFoundErrorCode, err.Code)
		assert.Equal(t, http.StatusNotFound, err.HttpStatus)
		assert.Equal(t, handlederror.DefaultErrorMessage, err.Message)
		assert.Equal(t, "Detail error", err.Detail)
	})
}
