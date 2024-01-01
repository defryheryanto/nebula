package main

import (
	"github.com/defryheryanto/nebula/config"
	"github.com/defryheryanto/nebula/internal/auth"
	"github.com/defryheryanto/nebula/internal/encrypt"
	"github.com/defryheryanto/nebula/internal/encrypt/aes"
	"github.com/defryheryanto/nebula/internal/token"
	jwtservice "github.com/defryheryanto/nebula/internal/token/jwt"
	"github.com/defryheryanto/nebula/internal/user"
	"github.com/golang-jwt/jwt"
)

type services struct {
	UserService user.Service
	AuthService auth.Service
}

func setupServices(r *repositories) *services {
	encryptor, err := aes.NewAESEncryptor(config.EncryptorSecret)
	if err != nil {
		panic(err)
	}
	tokener := jwtservice.NewJWTService[*auth.Session](jwt.SigningMethodES256, config.JWTSecret)

	userService := setupUserService(r.UserRepository)
	authService := setupAuthService(userService, encryptor, tokener)

	return &services{
		UserService: userService,
		AuthService: authService,
	}
}

func setupUserService(userRepo user.Repository) user.Service {
	return user.NewService(userRepo)
}

func setupAuthService(
	userService user.Service,
	encryptor encrypt.Encryptor,
	tokener token.Tokener[*auth.Session],
) auth.Service {
	return auth.NewService(
		userService,
		encryptor,
		tokener,
	)
}
