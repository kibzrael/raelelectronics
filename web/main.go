package main

import (
	"embed"
	"fmt"
	"net/http"
	"os"

	"github.com/kibzrael/raelelectronics/web/templates"
	"github.com/labstack/echo/v4"
)

//go:embed all:templates/components all:templates/pages
var resources embed.FS

func main() {
	e := echo.New()

	templates.NewTemplateRenderer(e, resources)
	e.Static("/static", "static")

	e.GET("/", func(e echo.Context) error {
		return e.Render(http.StatusOK, "index.html", nil)
	})

	fmt.Println("Web is listening on port", os.Getenv("PORT"))
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))

}
