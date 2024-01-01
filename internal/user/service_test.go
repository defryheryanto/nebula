package user_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/defryheryanto/nebula/internal/user"
	usermock "github.com/defryheryanto/nebula/internal/user/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type component struct {
	ctx        context.Context
	mockedErr  error
	repository *usermock.MockRepository
	service    user.Service
}

func setupService(t *testing.T) *component {
	g := gomock.NewController(t)

	repository := usermock.NewMockRepository(g)
	service := user.NewService(repository)

	return &component{
		ctx:        context.TODO(),
		mockedErr:  fmt.Errorf("mocked"),
		repository: repository,
		service:    service,
	}
}

func TestUserService_GetByUsername(t *testing.T) {
	t.Parallel()

	t.Run("error getting user", func(t *testing.T) {
		s := setupService(t)
		s.repository.EXPECT().First(gomock.Any(), gomock.Any()).Return(nil, s.mockedErr)

		result, err := s.service.GetByUsername(s.ctx, "username")
		assert.Nil(t, result)
		assert.Equal(t, s.mockedErr, err)
	})

	t.Run("success", func(t *testing.T) {
		s := setupService(t)
		expected := &user.User{
			ID:       1,
			Username: "username",
		}
		s.repository.EXPECT().First(gomock.Any(), gomock.Any()).Return(expected, nil)

		result, err := s.service.GetByUsername(s.ctx, "username")
		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	})
}
