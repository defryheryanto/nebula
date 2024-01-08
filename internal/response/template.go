package response

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/defryheryanto/nebula/config"
)

func FailedTemplate(w http.ResponseWriter, err error) {
	t, _ := template.ParseFiles(fmt.Sprintf("%s/static/error.html", config.WebFolderPath))
	t.Execute(w, nil)
}

func SuccessTemplate(w http.ResponseWriter, title string, templateName string, data any) {
	folderPath := config.WebFolderPath
	t, err := template.ParseFiles(fmt.Sprintf("%s%s", folderPath, templateName))
	if err != nil {
		FailedTemplate(w, err)
		return
	}

	type path struct {
		Assets   string
		Static   string
		Template string
	}
	type payload struct {
		Title string
		Path  path
		Data  any
	}

	p := &payload{
		Title: title,
		Path: path{
			Assets:   fmt.Sprintf("%s/assets", folderPath),
			Static:   fmt.Sprintf("%s/static", folderPath),
			Template: fmt.Sprintf("%s/template", folderPath),
		},
		Data: data,
	}
	t.Execute(w, p)
}
