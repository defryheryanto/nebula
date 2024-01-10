package view

import (
	"encoding/json"
	"log/slog"
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

func (h *Handler) LogDashboardView(w http.ResponseWriter, r *http.Request) {
	serviceName := r.URL.Query().Get("service")
	search := r.URL.Query().Get("search")
	page, pageSize, _ := request.GetPagination(r, 1, 20)

	type servicesPayload struct {
		Name string `json:"name"`
	}
	type logPayload struct {
		Timestamp string `json:"timestamp"`
		Log       any    `json:"log"`
	}
	type payload struct {
		Services []*servicesPayload `json:"services"`
		Logs     []*logPayload      `json:"logs"`
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

	if serviceName == "" && len(serviceNames) > 0 {
		serviceName = serviceNames[0]
	}
	resultLogs, err := h.logService.List(r.Context(), &logs.Filter{
		Page:        page,
		PageSize:    pageSize,
		ServiceName: serviceName,
		Search:      search,
	})
	if err != nil {
		response.FailedTemplate(w, err)
		return
	}

	logsData := make([]*logPayload, 0, len(resultLogs))
	for _, log := range resultLogs {
		logData := &logPayload{
			Timestamp: log.Timestamp.Local().Format("02 Jan 2006 15:04:05 MST"),
			Log:       log.Log,
		}
		if mapLog, ok := log.Log.(map[string]any); ok {
			logString, err := json.Marshal(&mapLog)
			if err != nil {
				continue
			}
			logData.Log = string(logString)
		}
		logsData = append(logsData, logData)
	}

	prevPageLink, err := request.GetPreviousPageLink(*r.URL)
	if err != nil {
		slog.Error("error LogsView.LogDashboardView.GetPreviousPageLink", "error", err)
	}

	nextPageLink, err := request.GetNextPageLink(*r.URL)
	if err != nil {
		slog.Error("error LogsView.LogDashboardView.GetNextPageLink", "error", err)
	}

	response.SuccessTemplate(w, "/template/logs/master.html", payload{
		Services: servicesPayloads,
		Logs:     logsData,
	}, &response.TemplateOptions{
		Title:            "Logs",
		PreviousPageLink: prevPageLink,
		NextPageLink:     nextPageLink,
	})
}
