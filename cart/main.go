package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/kibzrael/raelelectronics/cart/data"
	"github.com/kibzrael/raelelectronics/cart/routes"
	c "github.com/kibzrael/raelelectronics/common/api/cart"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

type CartServer struct {
	c.UnimplementedCartServiceServer
	db *sqlx.DB
}

func (s *CartServer) AddToWishlist(ctx context.Context, req *c.AddToWishlistRequest) (*c.AddToWishlistResponse, error) {
	ctx = context.WithValue(ctx, data.DB_CONTEXT, s.db)
	response, err := routes.AddToWishlistHandler(ctx, req)
	return response, err
}

func (s *CartServer) RemoveFromWishlist(ctx context.Context, req *c.RemoveFromWishlistRequest) (*c.RemoveFromWishlistResponse, error) {
	ctx = context.WithValue(ctx, data.DB_CONTEXT, s.db)
	response, err := routes.RemoveFromWishlistHandler(ctx, req)
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
	c.RegisterCartServiceServer(server, &CartServer{db: db})

	fmt.Println("Cart is listening on port", os.Getenv("PORT"))
	if err = server.Serve(lis); err != nil {
		log.Fatal(err.Error())
	}
}
