package view

import (
	"encoding/json"
	"net/http"

	"github.com/defryheryanto/nebula/internal/logs"
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

func (h *Handler) LogDashboardView(w http.ResponseWriter, r *http.Request) {
	type servicesPayload struct {
		Name string `json:"name"`
	}
	type payload struct {
		Services []*servicesPayload `json:"services"`
		Logs     []*logs.Log        `json:"logs"`
	}

	serviceNames, err := h.logService.GetAvailableServices(r.Context())
	if err != nil {
		response.FailedTemplate(w, err)
		return
	}
	servicesPayloads := make([]*servicesPayload, 0, len(serviceNames))
	for _, serviceName := range serviceNames {
		servicesPayloads = append(servicesPayloads, &servicesPayload{
			Name: serviceName,
		})
	}

	resultLogs, err := h.logService.List(r.Context())
	if err != nil {
		response.FailedTemplate(w, err)
		return
	}

	for _, log := range resultLogs {
		if mapLog, ok := log.Log.(map[string]any); ok {
			logString, err := json.Marshal(&mapLog)
			if err != nil {
				continue
			}
			log.Log = string(logString)
		}
	}

	response.SuccessTemplate(w, "Logs", "/template/logs/master.html", payload{
		Services: servicesPayloads,
		Logs:     resultLogs,
	})
}
