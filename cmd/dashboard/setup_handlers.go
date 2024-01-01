package main

import (
	authview "github.com/defryheryanto/nebula/internal/auth/http/view"
)

type handlers struct {
	AuthHandler *authview.Handler
}

func setupHandler() *handlers {
	return &handlers{
		AuthHandler: authview.NewHandler(),
	}
}
