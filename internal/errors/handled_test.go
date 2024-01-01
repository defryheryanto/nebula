package handlederror_test

import (
	"fmt"
	"net/http"
	"testing"

	handlederror "github.com/defryheryanto/nebula/internal/errors"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestHandledError(t *testing.T) {
	t.Parallel()
	t.Run("Should be an error type", func(t *testing.T) {
		var err error = handlederror.HandledError{
			HttpStatus: http.StatusBadRequest,
			Code:       handlederror.ValidationErrorCode,
			Message:    "Test Error Mocked",
			Detail:     "Error details",
		}

		assert.Equal(t, err.Error(), "Error details")
	})

	t.Run("With message", func(t *testing.T) {
		err := handlederror.HandledError{
			HttpStatus: http.StatusBadRequest,
			Code:       handlederror.ValidationErrorCode,
			Detail:     "Error details",
		}.WithMessage("Test bro")

		assert.Equal(t, "Test bro", err.Message)
	})
}

func TestExtract(t *testing.T) {
	t.Parallel()
	t.Run("Error given is not handled error", func(t *testing.T) {
		err := fmt.Errorf("mocked")
		err = errors.Wrap(err, "inner layer")
		err = errors.Wrap(err, "outer layer")

		handledErr := handlederror.Extract(err)

		assert.Equal(t, http.StatusInternalServerError, handledErr.HttpStatus)
		assert.Equal(t, handlederror.InternalServerErrorCode, handledErr.Code)
		assert.Equal(t, err.Error(), handledErr.Detail)
		assert.Equal(t, handlederror.DefaultErrorMessage, handledErr.Message)
	})

	t.Run("Error given is a handled error", func(t *testing.T) {
		var err error = handlederror.ValidationError("gatau lagi nih")
		err = errors.Wrap(err, "inner layer")
		err = errors.Wrap(err, "outer layer")

		handledErr := handlederror.Extract(err)

		assert.Equal(t, http.StatusBadRequest, handledErr.HttpStatus)
		assert.Equal(t, handlederror.ValidationErrorCode, handledErr.Code)
		assert.Equal(t, "gatau lagi nih", handledErr.Detail)
		assert.Equal(t, handlederror.DefaultErrorMessage, handledErr.Message)
	})
}
