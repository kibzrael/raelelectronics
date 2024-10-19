package feed

import (
	"context"
	"log"
	"net/http"
	"net/url"

	c "github.com/kibzrael/raelelectronics/common/api/catalog"
	"github.com/kibzrael/raelelectronics/web/services"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

func FeedPageHander(e echo.Context) error {
	conn := e.Get(services.CATALOG_CONTEXT).(*grpc.ClientConn)

	catalog := c.NewCatalogServiceClient(conn)

	query := e.QueryParams()

	data, err := catalog.LaptopFeed(context.Background(), &c.FeedRequest{
		Sort:    feedSort(query),
		Filters: feedFilters(query),
	})
	if err != nil {
		log.Println("failed to fetch laptop feed:", err)
	}

	return e.Render(http.StatusOK, "feed.html", data)
}

func feedSort(query url.Values) c.FeedSort {
	switch query.Get("sort") {
	case "new":
		return c.FeedSort_New
	case "price_asc":
		return c.FeedSort_PriceAsc
	case "price_desc":
		return c.FeedSort_PriceDesc
	default:
		return c.FeedSort_Featured
	}
}

func feedFilters(query url.Values) (filters []*c.FeedFilter) {
	for key, val := range query {
		filters = append(filters, &c.FeedFilter{Key: key, Val: val[0]})
	}
	return
}
