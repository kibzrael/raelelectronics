package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/kibzrael/raelelectronics/auth/data"
	"github.com/kibzrael/raelelectronics/auth/routes"
	"github.com/kibzrael/raelelectronics/auth/utils"
	_ "github.com/lib/pq"
)

func main() {
	ctx := context.Background()

	db, err := sqlx.Connect("postgres", fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable", data.DB_USER, data.DB_PASSWORD, data.DB_HOST, data.DB_NAME))
	if err != nil {
		log.Fatal(err.Error())
	}

	db.MustExec(data.Schema)
	ctx = context.WithValue(ctx, data.DB_CONTEXT, db)

	authRouter := http.ServeMux{}
	authRouter.Handle("POST /login", utils.AuthHandler(ctx, routes.LoginHandler))
	authRouter.Handle("POST /register", utils.AuthHandler(ctx, routes.RegisterHandler))

	server := http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: &authRouter,
	}
	fmt.Println("Auth is listening on port", os.Getenv("PORT"))
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
