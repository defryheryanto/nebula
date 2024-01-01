package jwt

import (
	"context"
	"time"

	handlederror "github.com/defryheryanto/nebula/internal/errors"
	"github.com/golang-jwt/jwt"
)

type CustomClaims[T any] struct {
	Payload T
	jwt.StandardClaims
}

type JWTService[T any] struct {
	signMethod jwt.SigningMethod
	secret     string
}

func NewJWTService[T any](signMethod jwt.SigningMethod, secret string) *JWTService[T] {
	return &JWTService[T]{signMethod, secret}
}

func (s *JWTService[T]) GenerateToken(ctx context.Context, payload T, expiryTime *time.Duration) (string, error) {
	standardClaims := jwt.StandardClaims{}
	if expiryTime != nil {
		standardClaims.ExpiresAt = time.Now().Add(*expiryTime).Unix()
	}

	customClaims := &CustomClaims[T]{
		Payload:        payload,
		StandardClaims: standardClaims,
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)
	tokenStr, err := jwtToken.SignedString([]byte(s.secret))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func (s *JWTService[T]) Validate(ctx context.Context, token string) (T, error) {
	customClaims := CustomClaims[T]{}
	jwtToken, err := jwt.ParseWithClaims(token, &customClaims, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.secret), nil
	})
	var payload T
	if err != nil {
		return payload, err
	}
	if !jwtToken.Valid {
		return payload, handlederror.ErrTokenInvalid
	}

	return customClaims.Payload, nil
}
