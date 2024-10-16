package templates

import (
	"embed"
	"io"
	"io/fs"
	"log"
	"strings"
	"text/template"

	"github.com/labstack/echo/v4"
)

type TemplateRegistry struct {
	templates *template.Template
}

func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	err := t.templates.ExecuteTemplate(w, name, data)
	if err != nil {
		log.Println(err)
	}
	return err
}

func NewTemplateRenderer(e *echo.Echo, resources embed.FS) {
	var paths []string
	fs.WalkDir(resources, ".", func(path string, d fs.DirEntry, err error) error {
		if strings.Contains(d.Name(), ".html") {
			paths = append(paths, path)
		}
		return nil
	})

	tmpl := template.Must(template.ParseFS(resources, paths...))

	t := newTemplate(tmpl)
	e.Renderer = t
}

func newTemplate(templates *template.Template) echo.Renderer {
	return &TemplateRegistry{
		templates: templates,
	}
}
