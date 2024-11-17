package wishlist

import (
	"context"
	"errors"
	"log"
	"net/http"

	c "github.com/kibzrael/raelelectronics/common/api/cart"
	"github.com/kibzrael/raelelectronics/web/services"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

func WishlistHandler(e echo.Context) (err error) {
	laptop := e.Param("uid")
	action := e.Param("action")

	status := http.StatusOK
	defer func() {
		if err != nil {
			log.Print("err", err)
			e.Render(status, "wishlist_button", map[string]any{"Laptop": laptop, "Wishlist": action != "true"})
		} else {
			e.Render(status, "wishlist_button", map[string]any{"Laptop": laptop, "Wishlist": action == "true"})
		}
	}()

	conn := e.Get(services.CART_CONTEXT).(*grpc.ClientConn)
	cart := c.NewCartServiceClient(conn)

	user, _ := e.Get(services.USER_CONTEXT).(string)

	if action == "true" {
		if res, err := cart.AddToWishlist(context.Background(), &c.AddToWishlistRequest{Uid: laptop, User: user}); err != nil {
			status = http.StatusInternalServerError
			return err
		} else if !res.Success {
			status = http.StatusInternalServerError
			return errors.New(res.Message)
		}
	} else {
		if res, err := cart.RemoveFromWishlist(context.Background(), &c.RemoveFromWishlistRequest{Uid: laptop, User: user}); err != nil {
			status = http.StatusInternalServerError
			return err
		} else if !res.Success {
			status = http.StatusInternalServerError
			return errors.New(res.Message)
		}
	}

	return nil
}
