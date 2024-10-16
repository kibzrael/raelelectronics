package pages

import (
	"log"
	"net/http"
	"sync"

	"github.com/jmoiron/sqlx"
	"github.com/kibzrael/raelelectronics/web/data"
	"github.com/labstack/echo/v4"
)

type NewArrival struct {
	Uid       string
	Name      string
	Thumbnail string
	Price     float64 `db:"price_min"`
}

type Featured struct {
	Uid       string
	Name      string
	Thumbnail string
	Price     float64 `db:"price_min"`
	Colors    string
	Size      float64
	Cpu       string
	Memory    int64
	Storage   int64
	Featured  float64
}

type LandingPage struct {
	NewArrivals []NewArrival
	Featured    []Featured
}

func newLandingPageData() LandingPage {
	return LandingPage{
		NewArrivals: []NewArrival{},
		Featured:    []Featured{},
	}
}

func LandingPageHandler(e echo.Context) error {
	db := e.Get(data.DB_CONTEXT).(*sqlx.DB)
	wg := sync.WaitGroup{}

	data := newLandingPageData()

	wg.Add(2)
	go fetchNewArrivals(db, &data, &wg)
	go fetchFeatured(db, &data, &wg)

	wg.Wait()
	return e.Render(http.StatusOK, "index.html", data)
}

func fetchNewArrivals(db *sqlx.DB, data *LandingPage, wg *sync.WaitGroup) {
	rows, _ := db.Queryx("SELECT laptops.uid, laptops.name, laptops.thumbnail, laptops.price_min FROM laptops ORDER BY laptops.launched DESC LIMIT 10")
	for rows.Next() {
		row := NewArrival{}
		if err := rows.StructScan(&row); err != nil {
			continue
		}
		data.NewArrivals = append(data.NewArrivals, row)
	}
	wg.Done()
}

func fetchFeatured(db *sqlx.DB, data *LandingPage, wg *sync.WaitGroup) {
	rows, _ := db.Queryx(`
		SELECT * FROM
			(SELECT
				DISTINCT ON (laptops.uid)  laptops.uid,
				laptops.name, laptops.thumbnail, laptops.price_min,
				displays.size, concat_ws(' ', cpus.architecture, cpus.model) as cpu,
				chassis.colors, memorys.size as memory, storages.capacity as storage,
				((memorys.size + (cpus.cores * cpus.base_speed)) / laptops.price_min) as featured
			FROM laptops
				LEFT JOIN chassis on laptops.uid = chassis.laptop
				LEFT JOIN displays on laptops.uid = displays.laptop
				LEFT JOIN cpus on laptops.uid = cpus.laptop
				LEFT JOIN memorys on laptops.uid = memorys.laptop
				LEFT JOIN storages on laptops.uid = storages.laptop
			ORDER BY laptops.uid, memorys.size ASC, storages.capacity ASC) l
		ORDER BY l.featured DESC
		LIMIT 4
	`)
	for rows.Next() {
		row := Featured{}
		if err := rows.StructScan(&row); err != nil {
			log.Println(err)
			continue
		}
		data.Featured = append(data.Featured, row)
	}
	wg.Done()
}
