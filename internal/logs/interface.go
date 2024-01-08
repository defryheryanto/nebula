package logs

import "context"

//go:generate mockgen -source interface.go -package logsmock -destination mock/mock.go

type Repository interface {
	Insert(ctx context.Context, service string, data any) error
	Find(ctx context.Context) ([]*Log, error)
}

type Service interface {
	Push(ctx context.Context, service string, data any) error
	List(ctx context.Context) ([]*Log, error)
}
