package auth_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/defryheryanto/nebula/internal/auth"
	encryptormock "github.com/defryheryanto/nebula/internal/encrypt/mock"
	handlederror "github.com/defryheryanto/nebula/internal/errors"
	tokenmock "github.com/defryheryanto/nebula/internal/token/mock"
	"github.com/defryheryanto/nebula/internal/user"
	usermock "github.com/defryheryanto/nebula/internal/user/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type components struct {
	ctx         context.Context
	mockedErr   error
	userService *usermock.MockService
	encryptor   *encryptormock.MockEncryptor
	tokener     *tokenmock.MockTokener[*auth.Session]
	service     auth.Service
}

func setupTest(t *testing.T) *components {
	g := gomock.NewController(t)
	userService := usermock.NewMockService(g)
	encryptor := encryptormock.NewMockEncryptor(g)
	tokener := tokenmock.NewMockTokener[*auth.Session](g)
	service := auth.NewService(
		userService,
		encryptor,
		tokener,
	)

	return &components{
		ctx:         context.TODO(),
		mockedErr:   fmt.Errorf("mocked"),
		userService: userService,
		encryptor:   encryptor,
		tokener:     tokener,
		service:     service,
	}
}

func TestService_AuthenticateUser(t *testing.T) {
	t.Parallel()

	t.Run("Failed get user", func(t *testing.T) {
		s := setupTest(t)
		s.userService.EXPECT().GetByUsername(gomock.Any(), gomock.Any()).Return(nil, s.mockedErr)

		token, err := s.service.AuthenticateUser(s.ctx, "username", "password")
		assert.Empty(t, token)
		assert.Equal(t, s.mockedErr, err)
	})

	t.Run("User not found", func(t *testing.T) {
		s := setupTest(t)
		s.userService.EXPECT().GetByUsername(gomock.Any(), gomock.Any()).Return(nil, nil)

		token, err := s.service.AuthenticateUser(s.ctx, "username", "password")
		assert.Empty(t, token)
		assert.Equal(t, handlederror.ErrInvalidCredentials, err)
	})

	t.Run("Failed check password", func(t *testing.T) {
		s := setupTest(t)
		s.userService.EXPECT().GetByUsername(gomock.Any(), gomock.Any()).Return(&user.User{}, nil)
		s.encryptor.EXPECT().Check(gomock.Any(), gomock.Any(), gomock.Any()).Return(false, s.mockedErr)

		token, err := s.service.AuthenticateUser(s.ctx, "username", "password")
		assert.Empty(t, token)
		assert.Equal(t, s.mockedErr, err)
	})

	t.Run("Invalid Password", func(t *testing.T) {
		s := setupTest(t)
		s.userService.EXPECT().GetByUsername(gomock.Any(), gomock.Any()).Return(&user.User{}, nil)
		s.encryptor.EXPECT().Check(gomock.Any(), gomock.Any(), gomock.Any()).Return(false, nil)

		token, err := s.service.AuthenticateUser(s.ctx, "username", "password")
		assert.Empty(t, token)
		assert.Equal(t, handlederror.ErrInvalidCredentials, err)
	})

	t.Run("Failed generate token", func(t *testing.T) {
		s := setupTest(t)
		s.userService.EXPECT().GetByUsername(gomock.Any(), gomock.Any()).Return(&user.User{}, nil)
		s.encryptor.EXPECT().Check(gomock.Any(), gomock.Any(), gomock.Any()).Return(true, nil)
		s.tokener.EXPECT().GenerateToken(gomock.Any(), gomock.Any(), gomock.Any()).Return("", s.mockedErr)

		token, err := s.service.AuthenticateUser(s.ctx, "username", "password")
		assert.Empty(t, token)
		assert.Equal(t, s.mockedErr, err)
	})

	t.Run("Success", func(t *testing.T) {
		s := setupTest(t)
		s.userService.EXPECT().GetByUsername(gomock.Any(), gomock.Any()).Return(&user.User{}, nil)
		s.encryptor.EXPECT().Check(gomock.Any(), gomock.Any(), gomock.Any()).Return(true, nil)
		s.tokener.EXPECT().GenerateToken(gomock.Any(), gomock.Any(), gomock.Any()).Return("token", nil)

		token, err := s.service.AuthenticateUser(s.ctx, "username", "password")
		assert.NoError(t, err)
		assert.Equal(t, "token", token)
	})
}

func TestService_ValidateToken(t *testing.T) {
	t.Parallel()

	t.Run("Failed to validate", func(t *testing.T) {
		s := setupTest(t)
		s.tokener.EXPECT().Validate(gomock.Any(), gomock.Any()).Return(nil, s.mockedErr)

		sess, err := s.service.ValidateToken(s.ctx, "token")
		assert.Nil(t, sess)
		assert.Equal(t, handlederror.ErrTokenInvalid, err)
	})

	t.Run("Success", func(t *testing.T) {
		s := setupTest(t)
		expected := &auth.Session{
			SessionID: "session-uuid",
			UserID:    1,
			Username:  "username",
		}
		s.tokener.EXPECT().Validate(s.ctx, "token").Return(expected, nil)

		sess, err := s.service.ValidateToken(s.ctx, "token")
		assert.NoError(t, err)
		assert.Equal(t, expected, sess)
	})
}
