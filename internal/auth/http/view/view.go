package view

import (
	"net/http"

	"github.com/defryheryanto/nebula/internal/auth"
	"github.com/defryheryanto/nebula/internal/response"
)

type Handler struct {
	AuthService auth.Service
}

func NewHandler(authService auth.Service) *Handler {
	return &Handler{
		AuthService: authService,
	}
}

func (h *Handler) LoginView(w http.ResponseWriter, r *http.Request) {
	data := NewLoginTemplateData()
	response.SuccessTemplate(w, "Login", data.TemplateName(), data)
}

func (h *Handler) LoginAction(w http.ResponseWriter, r *http.Request) {
	token, err := h.AuthService.AuthenticateUser(r.Context(), r.FormValue("username"), r.FormValue("password"))
	if err != nil {
		templateData := NewLoginTemplateData().WithError(err)
		response.SuccessTemplate(w, "Login", templateData.TemplateName(), templateData)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    token,
		HttpOnly: true,
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
