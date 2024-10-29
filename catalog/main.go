package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	c "github.com/kibzrael/raelelectronics/common/api/catalog"
	"google.golang.org/grpc"

	"github.com/jmoiron/sqlx"
	"github.com/kibzrael/raelelectronics/catalog/data"
	"github.com/kibzrael/raelelectronics/catalog/routes"
	_ "github.com/lib/pq"
)

type CatalogServer struct {
	c.UnimplementedCatalogServiceServer
	db *sqlx.DB
}

func (s *CatalogServer) LaptopFeed(ctx context.Context, req *c.FeedRequest) (*c.FeedResponse, error) {
	ctx = context.WithValue(ctx, data.DB_CONTEXT, s.db)
	laptops, pages, err := routes.FeedHandler(ctx, req)
	response := c.FeedResponse{
		Laptops: laptops,
		Current: req.Page,
		Pages:   pages,
	}
	return &response, err
}

func (s *CatalogServer) GetLaptopDetails(ctx context.Context, req *c.DetailsRequest) (*c.DetailsResponse, error) {
	ctx = context.WithValue(ctx, data.DB_CONTEXT, s.db)
	response, err := routes.DetailsHandler(ctx, req)
	return response, err
}

func main() {

	db, err := sqlx.Connect("postgres", fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable", data.DB_USER, data.DB_PASSWORD, data.DB_HOST, data.DB_NAME))
	if err != nil {
		log.Fatal(err.Error())
	}

	db.MustExec(data.Schema)

	lis, err := net.Listen("tcp", ":"+os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err.Error())
	}

	server := grpc.NewServer()
	c.RegisterCatalogServiceServer(server, &CatalogServer{db: db})

	fmt.Println("Catalog is listening on port", os.Getenv("PORT"))
	if err = server.Serve(lis); err != nil {
		log.Fatal(err.Error())
	}
}
