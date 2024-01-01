package view

import (
	"net/http"

	"github.com/defryheryanto/nebula/internal/response"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) LoginView(w http.ResponseWriter, r *http.Request) {
	data := NewLoginTemplateData()
	response.SuccessTemplate(w, data.TempateName(), data)
}
