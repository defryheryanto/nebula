package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetString(t *testing.T) {
	t.Run("Environment not found", func(t *testing.T) {
		res := getString("key", "defaultValue")
		assert.Equal(t, "defaultValue", res)
	})

	t.Run("Environment found", func(t *testing.T) {
		os.Setenv("key", "notDefaultValue")
		res := getString("key", "defaultValue")
		assert.Equal(t, "notDefaultValue", res)
	})
}
