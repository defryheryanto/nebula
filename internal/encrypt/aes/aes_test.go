package aes_test

import (
	"context"
	"testing"

	"github.com/defryheryanto/nebula/internal/encrypt/aes"
	"github.com/stretchr/testify/assert"
)

type components struct {
	ctx       context.Context
	encryptor *aes.AESEncryptor
}

func setupTest() *components {
	encryptor, _ := aes.NewAESEncryptor(("secret_need_to_be_32_characters!"))

	return &components{
		ctx:       context.TODO(),
		encryptor: encryptor,
	}
}

func TestAESEncryptor_New(t *testing.T) {
	t.Parallel()

	t.Run("Invalid Secret", func(t *testing.T) {
		_, err := aes.NewAESEncryptor("not_32_characters!")
		assert.Equal(t, aes.ErrInvalidSecret, err)
	})

	t.Run("Valid Secret", func(t *testing.T) {
		service, err := aes.NewAESEncryptor(("secret_need_to_be_32_characters!"))
		assert.NoError(t, err)
		assert.NotNil(t, service)
	})
}

func TestAESEncryptor_Encrypt(t *testing.T) {
	t.Parallel()
	s := setupTest()

	encrypted, err := s.encryptor.Encrypt(context.TODO(), "to be encrypted")
	assert.NoError(t, err)
	assert.NotEmpty(t, encrypted)
}

func TestAESEncryptor_Decrypt(t *testing.T) {
	t.Parallel()

	t.Run("Not Match", func(t *testing.T) {
		s := setupTest()

		isValid, err := s.encryptor.Check(context.TODO(), "not string", "NT2fasr2aWinx8SNxvIgLWYSKE1Ro1x7GQ4htnOSJ0VuIZ03PTC08db9Vw==")
		assert.NoError(t, err)
		assert.False(t, isValid)
	})

	t.Run("Match", func(t *testing.T) {
		s := setupTest()

		isValid, err := s.encryptor.Check(context.TODO(), "to be encrypted", "NT2fasr2aWinx8SNxvIgLWYSKE1Ro1x7GQ4htnOSJ0VuIZ03PTC08db9Vw==")
		assert.NoError(t, err)
		assert.True(t, isValid)
	})
}
