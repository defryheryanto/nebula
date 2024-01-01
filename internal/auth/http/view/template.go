package view

import "net/http"

type LoginTemplateData struct {
	Action       string
	ActionMethod string
	ErrorMessage string
}

func NewLoginTemplateData() *LoginTemplateData {
	return &LoginTemplateData{
		Action:       "/login/action",
		ActionMethod: http.MethodPost,
	}
}

func (d *LoginTemplateData) TempateName() string {
	return "/template/login.html"
}
