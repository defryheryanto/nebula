package main

import (
	authview "github.com/defryheryanto/nebula/internal/auth/http/view"
)

type handlers struct {
	AuthHandler *authview.Handler
}

func setupHandler(s *services) *handlers {
	return &handlers{
		AuthHandler: authview.NewHandler(s.AuthService),
	}
}
