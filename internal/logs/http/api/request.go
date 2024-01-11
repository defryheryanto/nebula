package api

type CreateLogRequest struct {
	ServiceName string `json:"service_name"`
	LogType     string `json:"log_type"`
	Log         any    `json:"log"`
}
