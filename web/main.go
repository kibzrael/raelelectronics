package main

import (
	"embed"
	"fmt"
	"os"

	"github.com/kibzrael/raelelectronics/web/data"
	"github.com/kibzrael/raelelectronics/web/services"
	"github.com/kibzrael/raelelectronics/web/templates"
	"github.com/kibzrael/raelelectronics/web/templates/components/wishlist"
	"github.com/kibzrael/raelelectronics/web/templates/pages"
	"github.com/kibzrael/raelelectronics/web/templates/pages/details"
	"github.com/kibzrael/raelelectronics/web/templates/pages/feed"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

//go:embed all:templates/components all:templates/pages
var resources embed.FS

func main() {
	e := echo.New()

	db := data.InitDatabase(e)
	defer db.Close()

	auth := services.InitAuthService(e)
	catalog := services.InitCatalogService(e)
	cart := services.InitCartService(e)

	defer auth.Close()
	defer catalog.Close()
	defer cart.Close()

	templates.NewTemplateRenderer(e, resources)
	e.Static("/static", "static")

	e.GET("/", pages.LandingPageHandler)
	e.GET("/feed", feed.FeedPageHander)
	e.GET("/l/:uid", details.DetailsPageHander)

	e.POST("/api/wishlist/:uid/:action", wishlist.WishlistHandler)

	fmt.Println("Web is listening on port", os.Getenv("PORT"))
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))

}
