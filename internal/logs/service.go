package logs

import "context"

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) Push(ctx context.Context, data any) error {
	if data == nil {
		return nil
	}

	err := s.repository.Insert(ctx, data)
	if err != nil {
		return err
	}

	return nil
}
