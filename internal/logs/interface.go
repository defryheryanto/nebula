package logs

import "context"

//go:generate mockgen -source interface.go -package logsmock -destination mock/mock.go

type Repository interface {
	Insert(ctx context.Context, data any) error
}

type Service interface {
	Push(ctx context.Context, data any) error
}
