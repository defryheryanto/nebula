package user

import "context"

//go:generate mockgen -source interface.go -package usermock -destination mock/mock.go

type (
	Repository interface {
		First(ctx context.Context, filter *Filter) (*User, error)
	}
	Service interface {
		GetByUsername(ctx context.Context, username string) (*User, error)
	}
)
