package view

import (
	"net/http"

	handlederror "github.com/defryheryanto/nebula/internal/errors"
)

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

func (d *LoginTemplateData) WithError(err error) *LoginTemplateData {
	handledErr := handlederror.Extract(err)
	d.ErrorMessage = handledErr.Message
	return d
}

func (d *LoginTemplateData) TemplateName() string {
	return "/template/login.html"
}
