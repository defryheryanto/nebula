package logs_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/defryheryanto/nebula/internal/logs"
	logsmock "github.com/defryheryanto/nebula/internal/logs/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type components struct {
	ctx         context.Context
	mockedError error
	repository  *logsmock.MockRepository
	service     logs.Service
}

func setupTest(t *testing.T) *components {
	g := gomock.NewController(t)

	repository := logsmock.NewMockRepository(g)
	service := logs.NewService(repository)

	return &components{
		ctx:         context.Background(),
		mockedError: fmt.Errorf("mocked"),
		repository:  repository,
		service:     service,
	}
}

func TestService_Push(t *testing.T) {
	t.Parallel()

	t.Run("Nil data", func(t *testing.T) {
		s := setupTest(t)

		err := s.service.Push(s.ctx, "nebula-dashboard", nil)
		assert.NoError(t, err)
	})

	t.Run("Insert Failed", func(t *testing.T) {
		s := setupTest(t)
		s.repository.EXPECT().Insert(gomock.Any(), gomock.Any(), gomock.Any()).Return(s.mockedError)

		err := s.service.Push(s.ctx, "nebula-dashboard", map[string]any{
			"foo": "bar",
		})
		assert.Equal(t, s.mockedError, err)
	})

	t.Run("Success", func(t *testing.T) {
		s := setupTest(t)
		data := map[string]any{
			"foo": "bar",
		}
		s.repository.EXPECT().Insert(s.ctx, "nebula-dashboard", data).Return(nil)

		err := s.service.Push(s.ctx, "nebula-dashboard", data)
		assert.NoError(t, err)
	})
}

func TestService_List(t *testing.T) {
	t.Parallel()

	t.Run("Failed find", func(t *testing.T) {
		s := setupTest(t)
		s.repository.EXPECT().Find(gomock.Any(), gomock.Any()).Return(nil, s.mockedError)

		res, err := s.service.List(s.ctx, &logs.Filter{})
		assert.Nil(t, res)
		assert.Equal(t, s.mockedError, err)
	})

	t.Run("Success", func(t *testing.T) {
		s := setupTest(t)
		expected := []*logs.Log{
			{
				Timestamp: time.Now(),
				Log:       "rwa",
			},
			{
				Timestamp: time.Now(),
				Log: map[string]any{
					"foo": "bar",
				},
			},
		}
		filter := &logs.Filter{}
		s.repository.EXPECT().Find(s.ctx, filter).Return(expected, nil)

		res, err := s.service.List(s.ctx, filter)
		assert.NoError(t, err)
		assert.Equal(t, expected, res)
	})
}

func TestService_GetAvailableServices(t *testing.T) {
	t.Parallel()

	t.Run("Failed Get", func(t *testing.T) {
		s := setupTest(t)

		s.repository.EXPECT().AvailableServices(gomock.Any()).Return(nil, s.mockedError)

		res, err := s.service.GetAvailableServices(s.ctx)
		assert.Nil(t, res)
		assert.Equal(t, s.mockedError, err)
	})

	t.Run("Success", func(t *testing.T) {
		s := setupTest(t)
		services := []string{"api1", "api2"}
		s.repository.EXPECT().AvailableServices(gomock.Any()).Return(services, nil)

		res, err := s.service.GetAvailableServices(s.ctx)
		assert.NoError(t, err)
		assert.Equal(t, services, res)
	})
}
