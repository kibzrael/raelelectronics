package routes

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/kibzrael/raelelectronics/auth/data"
	c "github.com/kibzrael/raelelectronics/common/api/auth"
)

var JwtKey string = os.Getenv("JWT_PRIVATE_KEY")

func CreateTokenHandler(ctx context.Context, req *c.CreateTokenRequest) (*c.CreateTokenResponse, error) {
	db := ctx.Value(data.DB_CONTEXT).(*sqlx.DB)

	user := jwt.MapClaims{
		"uid":        uuid.New().String(),
		"ip":         req.Ip,
		"user_agent": req.UserAgent,
		"logged_at":  time.Now(),
	}
	_, err := db.NamedExec("INSERT INTO users (uid, ip, user_agent, logged_at) VALUES (:uid, :ip, :user_agent, :logged_at)", user)
	if err != nil {
		return nil, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, user)
	result, err := token.SignedString([]byte(JwtKey))
	if err != nil {
		return nil, err
	}

	return &c.CreateTokenResponse{
		User:  user["uid"].(string),
		Token: result,
	}, nil
}

func ValidateTokenHandler(ctx context.Context, req *c.ValidateTokenRequest) (*c.ValidateTokenResponse, error) {
	token, err := jwt.Parse(req.Token, func(token *jwt.Token) (interface{}, error) {
		return []byte(JwtKey), nil
	})
	if err != nil {
		return &c.ValidateTokenResponse{Valid: false}, err
	} else if !token.Valid {
		return &c.ValidateTokenResponse{Valid: false}, errors.New("invalid token")
	}

	claims, _ := token.Claims.(jwt.MapClaims)
	return &c.ValidateTokenResponse{Valid: true, User: claims["uid"].(string)}, nil
}
