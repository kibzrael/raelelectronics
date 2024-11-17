package services

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const CART_HOST = "cart"
const CART_PORT = "8003"

const CART_CONTEXT = "cart"

func InitCartService(e *echo.Echo) *grpc.ClientConn {
	cart, err := grpc.NewClient(fmt.Sprintf("%s:%s", CART_HOST, CART_PORT), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("failed to connect to cart service:", err)
	}
	cart.Connect()
	e.Use(ServiceMiddleware(cart, CART_CONTEXT))
	return cart
}
