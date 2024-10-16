package data

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func DbMiddleware(db *sqlx.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(DB_CONTEXT, db)
			return next(c)
		}
	}
}
