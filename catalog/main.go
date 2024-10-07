package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/kibzrael/raelelectronics/catalog/data"
	"github.com/kibzrael/raelelectronics/catalog/routes"
	"github.com/kibzrael/raelelectronics/catalog/utils"
	_ "github.com/lib/pq"
)

func main() {

	db, err := sqlx.Connect("postgres", fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable", data.DB_USER, data.DB_PASSWORD, data.DB_HOST, data.DB_NAME))
	if err != nil {
		log.Fatal(err.Error())
	}

	db.MustExec(data.Schema)

	ctx := context.WithValue(context.Background(), data.DB_CONTEXT, db)

	// data.SeedLaptops(ctx)

	catalogRouter := http.ServeMux{}
	catalogRouter.Handle("GET /feed", utils.CatalogHandler(ctx, routes.FeedHandler))
	catalogRouter.Handle("POST /seed", utils.CatalogHandler(ctx, data.SeedLaptops))
	catalogRouter.Handle("GET /{uid}", utils.CatalogHandler(ctx, routes.DetailsHandler))

	server := http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: &catalogRouter,
	}
	fmt.Println("Catalog is listening on port", os.Getenv("PORT"))
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
