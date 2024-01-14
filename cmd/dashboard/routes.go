package main

import (
	"fmt"
	"net/http"

	"github.com/defryheryanto/nebula/config"
	"github.com/defryheryanto/nebula/internal/auth"
	"github.com/go-chi/chi/v5"
)

func buildRoutes(h *handlers, authService auth.Service) http.Handler {
	root := chi.NewRouter()

	root.Handle("/static/assets/*", http.StripPrefix("/static/assets/", http.FileServer(http.Dir(fmt.Sprintf("%s/assets", config.WebFolderPath)))))

	root.Group(func(r chi.Router) {
		r.Use(auth.RedirectAuthorizedMiddleware(authService, "/dashboard/logs"))
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "/dashboard/logs", http.StatusSeeOther)
		})
		r.Get("/login", h.AuthHandler.LoginView)
		r.Post("/login/action", h.AuthHandler.LoginAction)
	})

	root.Group(func(r chi.Router) {
		r.Post("/api/logs", h.LogAPIHandler.CreateLog)
	})

	root.Route("/dashboard", func(r chi.Router) {
		r.Use(auth.CheckTokenMiddleware(authService))
		r.Get("/logs", h.LogViewhandler.LogDashboardView)
	})

	return root
}
