package api

type CreateLogRequest struct {
	ServiceName string `json:"service_name"`
	Log         any    `json:"log"`
}
