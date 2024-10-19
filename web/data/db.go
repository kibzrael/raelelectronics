package data

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

const DB_HOST = "db"

var DB_USER = os.Getenv("POSTGRES_USER")
var DB_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
var DB_NAME = os.Getenv("POSTGRES_DB")

const DB_CONTEXT string = "db"

func InitDatabase(e *echo.Echo) *sqlx.DB {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable", DB_USER, DB_PASSWORD, DB_HOST, DB_NAME))
	if err != nil {
		log.Fatal(err.Error())
	}
	e.Use(DbMiddleware(db))
	return db
}
