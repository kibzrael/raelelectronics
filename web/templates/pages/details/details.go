package details

import (
	"context"
	"log"
	"net/http"

	c "github.com/kibzrael/raelelectronics/common/api/catalog"
	"github.com/kibzrael/raelelectronics/web/services"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

type Details struct {
	Laptop  *c.Laptop
	Related []*c.LaptopCard
}

func newDetailsData() Details {
	return Details{
		Laptop:  &c.Laptop{},
		Related: []*c.LaptopCard{},
	}
}

func DetailsPageHander(e echo.Context) error {
	uid := e.Param("uid")

	conn := e.Get(services.CATALOG_CONTEXT).(*grpc.ClientConn)

	catalog := c.NewCatalogServiceClient(conn)

	response, err := catalog.GetLaptopDetails(context.Background(), &c.DetailsRequest{Uid: uid})
	if err != nil {
		log.Println("failed to fetch laptop:", err)
	}

	data := newDetailsData()
	data.Laptop = response.Laptop
	data.Related = response.Related

	return e.Render(http.StatusOK, "details.html", data)
}
