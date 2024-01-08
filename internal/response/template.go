package response

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/defryheryanto/nebula/config"
)

type TemplateOptions struct {
	Title            string
	PreviousPageLink string
	NextPageLink     string
}

func FailedTemplate(w http.ResponseWriter, err error) {
	t, _ := template.ParseFiles(fmt.Sprintf("%s/static/error.html", config.WebFolderPath))
	t.Execute(w, nil)
}

func SuccessTemplate(w http.ResponseWriter, templateName string, data any, templateOption *TemplateOptions) {
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
	type pagination struct {
		PreviousPage string
		NextPage     string
	}
	type payload struct {
		Title      string
		Pagination pagination
		Path       path
		Data       any
	}

	title := "Nebula Dashboard"
	paginationOption := pagination{
		PreviousPage: "#",
		NextPage:     "#",
	}

	if templateOption != nil {
		title = templateOption.Title
		paginationOption.PreviousPage = templateOption.PreviousPageLink
		paginationOption.NextPage = templateOption.NextPageLink
	}

	p := &payload{
		Title:      title,
		Pagination: paginationOption,
		Path: path{
			Assets:   fmt.Sprintf("%s/assets", folderPath),
			Static:   fmt.Sprintf("%s/static", folderPath),
			Template: fmt.Sprintf("%s/template", folderPath),
		},
		Data: data,
	}
	t.Execute(w, p)
}
