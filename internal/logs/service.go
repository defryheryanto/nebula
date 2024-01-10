package logs

import (
	"context"
	"encoding/json"
)

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) Push(ctx context.Context, service string, data any) error {
	if data == nil {
		return nil
	}

	dataBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = s.repository.Insert(ctx, service, string(dataBytes))
	if err != nil {
		return err
	}

	return nil
}

func (s *service) List(ctx context.Context, filter *Filter) ([]*Log, error) {
	result, err := s.repository.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *service) GetAvailableServices(ctx context.Context) ([]string, error) {
	result, err := s.repository.AvailableServices(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}
