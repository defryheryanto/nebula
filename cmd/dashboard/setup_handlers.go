package main

import (
	authview "github.com/defryheryanto/nebula/internal/auth/http/view"
	logsapi "github.com/defryheryanto/nebula/internal/logs/http/api"
)

type handlers struct {
	AuthHandler   *authview.Handler
	LogAPIHandler *logsapi.Handler
}

func setupHandler(s *services) *handlers {
	return &handlers{
		AuthHandler:   authview.NewHandler(s.AuthService),
		LogAPIHandler: logsapi.NewHandler(s.LogService),
	}
}
