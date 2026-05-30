package handler

import (
	"html/template"
	"io"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

type TemplateRenderer struct {
	templates map[string]*template.Template
}

func NewTemplateRenderer(pattern string) (*TemplateRenderer, error) {
	layoutPath := filepath.Join("templates", "layout.html")
	pages, err := filepath.Glob(filepath.Join("templates", "*.html"))
	if err != nil {
		return nil, err
	}

	templates := make(map[string]*template.Template)
	for _, page := range pages {
		if filepath.Base(page) == "layout.html" {
			continue
		}
		name := filepath.Base(page)
		tmpl, err := template.ParseFiles(layoutPath, page)
		if err != nil {
			return nil, err
		}
		templates[name] = tmpl
	}

	return &TemplateRenderer{templates: templates}, nil
}

func (r *TemplateRenderer) Render(w io.Writer, name string, data interface{}, _ echo.Context) error {
	tmpl, ok := r.templates[name]
	if !ok {
		return echo.ErrNotFound
	}
	return tmpl.ExecuteTemplate(w, "layout", data)
}
