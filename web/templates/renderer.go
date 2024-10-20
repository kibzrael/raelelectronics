package templates

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"log"
	"strings"
	"text/template"

	c "github.com/kibzrael/raelelectronics/common/api/catalog"
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

	tmpl := template.Must(template.New("").Funcs(template.FuncMap{
		"hasFilter":   hasFilter,
		"indexFilter": indexFilter,
	}).ParseFS(resources, paths...))

	t := newTemplate(tmpl)
	e.Renderer = t
}

func newTemplate(templates *template.Template) echo.Renderer {
	return &TemplateRegistry{
		templates: templates,
	}
}

func hasFilter(filters []*c.FeedFilter, key string, val interface{}) bool {
	for _, f := range filters {
		for _, filter := range strings.Split(f.Val, ",") {
			if f.Key == key && filter == fmt.Sprintf("%v", val) {
				return true
			}
		}

	}
	return false
}

func indexFilter(filters []*c.FeedFilter, key string) string {
	for _, filter := range filters {
		if filter.Key == key {
			return filter.Val
		}
	}
	return ""
}
