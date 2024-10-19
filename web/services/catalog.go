package services

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const CATALOG_HOST = "catalog"
const CATALOG_PORT = "8002"

const CATALOG_CONTEXT = "catalog"

func InitCatalogService(e *echo.Echo) *grpc.ClientConn {
	catalog, err := grpc.NewClient(fmt.Sprintf("%s:%s", CATALOG_HOST, CATALOG_PORT), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("failed to connect to catalog service:", err)
	}
	catalog.Connect()
	e.Use(ServiceMiddleware(catalog, CATALOG_CONTEXT))
	return catalog
}
