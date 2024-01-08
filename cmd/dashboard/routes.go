package main

import (
	"fmt"
	"net/http"

	"github.com/defryheryanto/nebula/config"
	"github.com/go-chi/chi/v5"
)

func buildRoutes(h *handlers) http.Handler {
	root := chi.NewRouter()

	root.Handle("/static/assets/*", http.StripPrefix("/static/assets/", http.FileServer(http.Dir(fmt.Sprintf("%s/assets", config.WebFolderPath)))))

	root.Get("/login", h.AuthHandler.LoginView)
	root.Post("/login/action", h.AuthHandler.LoginAction)

	root.Group(func(r chi.Router) {
		r.Post("/api/logs", h.LogAPIHandler.CreateLog)
	})

	root.Route("/dashboard", func(r chi.Router) {
		r.Get("/logs", h.LogViewhandler.LogDashboardView)
	})

	return root
}
