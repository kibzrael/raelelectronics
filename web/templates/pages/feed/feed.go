package feed

import (
	"context"
	"log"
	"net/http"
	"net/url"

	c "github.com/kibzrael/raelelectronics/common/api/catalog"
	"github.com/kibzrael/raelelectronics/web/services"
	"github.com/kibzrael/raelelectronics/web/utils"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

type Feed struct {
	Laptops []*c.LaptopCard

	Colors       []utils.Color
	DisplayTypes []string
	DisplaySizes []uint64
	GPUBrands    []string
	CPUs         []string
	MemorySizes  []uint64
	StorageSizes []uint64
	BrandFilters []utils.BrandFilter
}

func newFeedData() Feed {
	return Feed{
		Laptops:      []*c.LaptopCard{},
		Colors:       utils.Colors,
		DisplayTypes: utils.DisplayTypes,
		DisplaySizes: utils.DisplaySizes,
		GPUBrands:    utils.GPUBrands,
		CPUs:         utils.CPUs,
		MemorySizes:  utils.MemorySizes,
		StorageSizes: utils.StorageSizes,
		BrandFilters: utils.BrandFilters,
	}
}

func FeedPageHander(e echo.Context) error {
	conn := e.Get(services.CATALOG_CONTEXT).(*grpc.ClientConn)

	catalog := c.NewCatalogServiceClient(conn)

	query := e.QueryParams()

	response, err := catalog.LaptopFeed(context.Background(), &c.FeedRequest{
		Sort:    feedSort(query),
		Filters: feedFilters(query),
	})
	if err != nil {
		log.Println("failed to fetch laptop feed:", err)
	}

	data := newFeedData()
	data.Laptops = response.Laptops

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
