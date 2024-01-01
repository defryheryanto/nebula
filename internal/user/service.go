package user

import "context"

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) GetByUsername(ctx context.Context, username string) (*User, error) {
	result, err := s.repository.First(ctx, &Filter{
		Username: username,
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}
