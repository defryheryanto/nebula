package jwt_test

import (
	"context"
	"testing"
	"time"

	jwtservice "github.com/defryheryanto/nebula/internal/token/jwt"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

type TestJWTPayload struct {
	Uid   int
	Email string
}

type components struct {
	ctx     context.Context
	service *jwtservice.JWTService[*TestJWTPayload]
}

func setupTest() *components {
	return &components{
		ctx:     context.TODO(),
		service: jwtservice.NewJWTService[*TestJWTPayload](jwt.SigningMethodHS256, "my_most_secured_secret_key"),
	}
}

func TestJwtService_GenerateToken(t *testing.T) {
	t.Parallel()
	s := setupTest()

	token, err := s.service.GenerateToken(
		s.ctx,
		&TestJWTPayload{
			Uid:   1,
			Email: "admin@piggybank.com",
		},
		nil,
	)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestJwtService_Validate(t *testing.T) {
	t.Parallel()
	t.Run("Empty Token", func(t *testing.T) {
		s := setupTest()
		_, err := s.service.Validate(s.ctx, "")
		assert.Error(t, err)
	})

	t.Run("Invalid Token", func(t *testing.T) {
		s := setupTest()
		_, err := s.service.Validate(context.TODO(), "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJQYXlsb2FkIjp7IlVpZCI6MSwiRW1haWwiOiJhZG1pbkBwaWdneWJhbmsuY29tIiwiSGFoYSI6ImhlaGUifX0.xElnCeuuFJYrsoK2efmgnTk7LWeymOEOvBhA1vLTypM")
		assert.Error(t, err)
	})

	t.Run("Success", func(t *testing.T) {
		s := setupTest()
		initialPayload := &TestJWTPayload{
			Uid:   1,
			Email: "admin@piggybank.com",
		}

		duration := 24 * time.Hour
		token, err := s.service.GenerateToken(context.TODO(), initialPayload, &duration)
		assert.NoError(t, err)
		assert.NotEmpty(t, token)

		payload, err := s.service.Validate(context.TODO(), token)
		assert.NoError(t, err)
		assert.EqualValues(t, initialPayload, payload)
	})
}
