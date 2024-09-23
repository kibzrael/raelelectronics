package utils

import (
	"context"
	"net/http"
)

type handler = func(ctx context.Context, res http.ResponseWriter, req *http.Request)

func AuthHandler(ctx context.Context, next handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		next(ctx, res, req)
	})
}
