package token

import (
	"context"
	"time"
)

//go:generate mockgen -source token.go -package tokenmock -destination mock/mock.go

type Tokener[T any] interface {
	GenerateToken(ctx context.Context, payload T, expiryTime *time.Duration) (string, error)
	Validate(ctx context.Context, token string) (T, error)
}
