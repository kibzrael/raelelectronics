package main

import (
	"embed"
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/kibzrael/raelelectronics/web/data"
	"github.com/kibzrael/raelelectronics/web/templates"
	"github.com/kibzrael/raelelectronics/web/templates/pages"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

//go:embed all:templates/components all:templates/pages
var resources embed.FS

func main() {
	e := echo.New()

	db, err := sqlx.Connect("postgres", fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable", data.DB_USER, data.DB_PASSWORD, data.DB_HOST, data.DB_NAME))
	if err != nil {
		log.Fatal(err.Error())
	}
	e.Use(data.DbMiddleware(db))

	templates.NewTemplateRenderer(e, resources)
	e.Static("/static", "static")

	e.GET("/", pages.LandingPageHandler)

	fmt.Println("Web is listening on port", os.Getenv("PORT"))
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))

}
