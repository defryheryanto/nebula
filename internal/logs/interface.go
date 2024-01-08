package logs

import "context"

//go:generate mockgen -source interface.go -package logsmock -destination mock/mock.go

type Repository interface {
	Insert(ctx context.Context, service string, data any) error
	Find(ctx context.Context, filter *Filter) ([]*Log, error)
	AvailableServices(ctx context.Context) ([]string, error)
}

type Service interface {
	Push(ctx context.Context, service string, data any) error
	List(ctx context.Context, filter *Filter) ([]*Log, error)
	GetAvailableServices(ctx context.Context) ([]string, error)
}
