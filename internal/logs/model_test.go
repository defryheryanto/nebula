package logs_test

import (
	"testing"

	"github.com/defryheryanto/nebula/internal/logs"
	"github.com/stretchr/testify/assert"
)

func TestLogTypeFromString(t *testing.T) {
	t.Parallel()

	t.Run("Unknown", func(t *testing.T) {
		res := logs.LogTypeFromString("unsure")
		assert.Empty(t, res)
	})

	t.Run("Http Log Type", func(t *testing.T) {
		res := logs.LogTypeFromString("http-log")
		assert.Equal(t, logs.HttpLogType, res)
	})

	t.Run("Std Type", func(t *testing.T) {
		res := logs.LogTypeFromString("std-log")
		assert.Equal(t, logs.StdLogType, res)
	})
}
