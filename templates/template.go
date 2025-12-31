package templates

import (
	"embed"
	"html/template"
)

//go:embed */*.html
var templatesDir embed.FS

func GetTemplatesFS() (*template.Template, error) {
	tmpl, err := template.ParseFS(templatesDir, "*/*.html")
	if err != nil {
		return nil, err
	}
	return tmpl, nil
}
