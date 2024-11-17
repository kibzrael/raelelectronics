package routes

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/kibzrael/raelelectronics/cart/data"
	c "github.com/kibzrael/raelelectronics/common/api/cart"
)

func AddToWishlistHandler(ctx context.Context, req *c.AddToWishlistRequest) (*c.AddToWishlistResponse, error) {
	db := ctx.Value(data.DB_CONTEXT).(*sqlx.DB)

	item := data.WishlistItem{
		Uid:    uuid.New().String(),
		Laptop: req.Uid,
		User:   req.User,
	}
	_, err := db.NamedExec("INSERT INTO wishlist (uid, laptop, \"user\") VALUES(:uid, :laptop, :user)", &item)
	if err != nil {
		return &c.AddToWishlistResponse{Success: false, Message: err.Error()}, err
	}

	return &c.AddToWishlistResponse{Success: true, Message: "Item added to wishlist"}, nil
}

func RemoveFromWishlistHandler(ctx context.Context, req *c.RemoveFromWishlistRequest) (*c.RemoveFromWishlistResponse, error) {
	db := ctx.Value(data.DB_CONTEXT).(*sqlx.DB)

	_, err := db.Exec("DELETE FROM wishlist WHERE laptop = $1 AND \"user\" = $2", req.Uid, req.User)
	if err != nil {
		return &c.RemoveFromWishlistResponse{Success: false, Message: err.Error()}, err
	}
	return &c.RemoveFromWishlistResponse{Success: true, Message: "Item removed from wishlist"}, nil
}
