package auth

import "context"

//go:generate mockgen -source interface.go -package authmock -destination mock/mock.go

type Service interface {
	AuthenticateUser(ctx context.Context, username string, password string) (string, error)
	ValidateToken(ctx context.Context, token string) (*Session, error)
}
