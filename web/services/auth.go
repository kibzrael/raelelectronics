package services

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	a "github.com/kibzrael/raelelectronics/common/api/auth"
	"github.com/labstack/echo/v4"
	"github.com/patrickmn/go-cache"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const AUTH_HOST = "auth"
const AUTH_PORT = "8001"

const AUTH_CONTEXT = "auth"

const USER_CONTEXT = "user"

var IGNORE_PATHS = []string{"/static"}

func InitAuthService(e *echo.Echo) *grpc.ClientConn {
	auth, err := grpc.NewClient(fmt.Sprintf("%s:%s", AUTH_HOST, AUTH_PORT), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("failed to connect to auth service:", err)
	}
	auth.Connect()

	c := cache.New(15*time.Minute, 15*time.Minute)

	e.Use(ServiceMiddleware(auth, AUTH_CONTEXT))
	e.Use(AuthMiddleware(auth, c))
	return auth
}

func AuthMiddleware(auth *grpc.ClientConn, c *cache.Cache) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(e echo.Context) error {
			for _, prefix := range IGNORE_PATHS {
				if strings.HasPrefix(e.Path(), prefix) {
					return next(e)
				}
			}

			conn := e.Get(AUTH_CONTEXT).(*grpc.ClientConn)
			auth := a.NewAuthServiceClient(conn)

			createToken := func() {
				// Create anonymous user token
				ip := e.RealIP()
				userAgent := e.Request().UserAgent()
				if res, err := auth.CreateToken(context.Background(), &a.CreateTokenRequest{Ip: ip, UserAgent: userAgent}); err != nil {
					log.Println("Error creating token", err)
				} else {
					e.Set(USER_CONTEXT, res.User)
					e.SetCookie(&http.Cookie{Name: "token", Value: res.Token})
					c.Set(res.Token, res.User, cache.DefaultExpiration)
				}
			}

			if cookie, err := e.Cookie("token"); err != nil {
				createToken()
			} else {
				// Check if token is in the cache
				user, found := c.Get(cookie.Value)
				if !found {
					res, err := auth.ValidateToken(context.Background(), &a.ValidateTokenRequest{Token: cookie.Value})
					if err != nil || !res.Valid {
						createToken()
						log.Println("Error validating token", err)
					} else if res.Valid {
						e.Set(USER_CONTEXT, res.User)
						c.Set(cookie.Value, res.User, cache.DefaultExpiration)
					}
				} else {
					e.Set(USER_CONTEXT, user.(string))
				}
			}
			return next(e)
		}
	}
}
