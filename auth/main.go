package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/kibzrael/raelelectronics/auth/data"
	"github.com/kibzrael/raelelectronics/auth/routes"
	c "github.com/kibzrael/raelelectronics/common/api/auth"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

type AuthServer struct {
	c.UnimplementedAuthServiceServer
	db *sqlx.DB
}

func (s *AuthServer) CreateToken(ctx context.Context, req *c.CreateTokenRequest) (*c.CreateTokenResponse, error) {
	ctx = context.WithValue(ctx, data.DB_CONTEXT, s.db)
	response, err := routes.CreateTokenHandler(ctx, req)
	return response, err
}

func (s *AuthServer) ValidateToken(ctx context.Context, req *c.ValidateTokenRequest) (*c.ValidateTokenResponse, error) {
	ctx = context.WithValue(ctx, data.DB_CONTEXT, s.db)
	response, err := routes.ValidateTokenHandler(ctx, req)
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
	c.RegisterAuthServiceServer(server, &AuthServer{db: db})

	fmt.Println("Auth is listening on port", os.Getenv("PORT"))
	if err = server.Serve(lis); err != nil {
		log.Fatal(err.Error())
	}

	// authRouter.Handle("POST /login", utils.AuthHandler(ctx, routes.LoginHandler))
	// authRouter.Handle("POST /register", utils.AuthHandler(ctx, routes.RegisterHandler))

}
