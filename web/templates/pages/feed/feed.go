package feed

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"

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

	Sort    string
	Filters []*c.FeedFilter

	Current int64
	Pages   int64
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
		Sort:         "featured",
		Filters:      []*c.FeedFilter{},
		Current:      1,
		Pages:        1,
	}
}

func FeedPageHander(e echo.Context) error {
	conn := e.Get(services.CATALOG_CONTEXT).(*grpc.ClientConn)

	catalog := c.NewCatalogServiceClient(conn)

	query := e.QueryParams()

	data := newFeedData()

	data.Filters = feedFilters(query)
	var sortQuery c.FeedSort
	sortQuery, data.Sort = feedSort(query)
	page, _ := strconv.ParseInt(query.Get("page"), 10, 64)
	if page == 0 {
		page = 1
	}

	response, err := catalog.LaptopFeed(context.Background(), &c.FeedRequest{
		Sort:    sortQuery,
		Filters: data.Filters,
		Page:    page,
	})
	if err != nil {
		log.Println("failed to fetch laptop feed:", err)
	}

	data.Laptops = response.Laptops
	data.Current = response.Current
	data.Pages = response.Pages

	return e.Render(http.StatusOK, "feed.html", data)
}

func feedSort(query url.Values) (c.FeedSort, string) {
	switch query.Get("sort") {
	case "new":
		return c.FeedSort_New, "new"
	case "price_asc":
		return c.FeedSort_PriceAsc, "price_asc"
	case "price_desc":
		return c.FeedSort_PriceDesc, "price_desc"
	default:
		return c.FeedSort_Featured, "featured"
	}
}

func feedFilters(query url.Values) (filters []*c.FeedFilter) {
	for key, val := range query {
		values := strings.Trim(strings.Join(val, ","), " ")
		if values != "" && key != "sort" && key != "page" {
			filters = append(filters, &c.FeedFilter{Key: key, Val: values})
		}
	}
	sort.Slice(filters, func(i, j int) bool {
		return i < j
	})
	return
}
