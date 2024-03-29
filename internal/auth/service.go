package auth

import (
	"context"
	"time"

	"github.com/defryheryanto/nebula/internal/encrypt"
	handlederror "github.com/defryheryanto/nebula/internal/errors"
	"github.com/defryheryanto/nebula/internal/token"
	"github.com/defryheryanto/nebula/internal/user"
	"github.com/google/uuid"
)

type service struct {
	userService user.Service
	encryptor   encrypt.Encryptor
	tokener     token.Tokener[*Session]
}

func NewService(
	userService user.Service,
	encryptor encrypt.Encryptor,
	tokener token.Tokener[*Session],
) Service {
	return &service{
		userService: userService,
		encryptor:   encryptor,
		tokener:     tokener,
	}
}

func (s *service) AuthenticateUser(ctx context.Context, username string, password string) (string, error) {
	if username == "" || password == "" {
		return "", handlederror.ErrInvalidCredentials
	}

	targetUser, err := s.userService.GetByUsername(ctx, username)
	if err != nil {
		return "", err
	}
	if targetUser == nil {
		return "", handlederror.ErrInvalidCredentials
	}

	isPasswordValid, err := s.encryptor.Check(ctx, password, targetUser.Password)
	if err != nil {
		return "", err
	}
	if !isPasswordValid {
		return "", handlederror.ErrInvalidCredentials
	}

	session := &Session{
		SessionID: uuid.New().String(),
		UserID:    targetUser.ID,
		Username:  targetUser.Username,
	}
	expiryDuration := 24 * time.Hour
	token, err := s.tokener.GenerateToken(ctx, session, &expiryDuration)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *service) ValidateToken(ctx context.Context, token string) (*Session, error) {
	currentSession, err := s.tokener.Validate(ctx, token)
	if err != nil {
		return nil, handlederror.ErrTokenInvalid
	}

	return currentSession, nil
}
