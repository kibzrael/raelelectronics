package services

import (
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

func ServiceMiddleware(conn *grpc.ClientConn, key string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(key, conn)
			return next(c)
		}
	}
}
