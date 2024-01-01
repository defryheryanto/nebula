package api

import (
	"net/http"

	"github.com/defryheryanto/nebula/internal/logs"
	"github.com/defryheryanto/nebula/internal/request"
	"github.com/defryheryanto/nebula/internal/response"
)

type Handler struct {
	logService logs.Service
}

func NewHandler(logService logs.Service) *Handler {
	return &Handler{
		logService: logService,
	}
}

func (h *Handler) CreateLog(w http.ResponseWriter, r *http.Request) {
	body := &CreateLogRequest{}
	err := request.DecodeBody(r, &body)
	if err != nil {
		response.FailedJSON(w, r, err)
		return
	}

	err = h.logService.Push(r.Context(), body.Log)
	if err != nil {
		response.FailedJSON(w, r, err)
		return
	}

	response.SuccessJSON[any](w, http.StatusOK, nil)
}
