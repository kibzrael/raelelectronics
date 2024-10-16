package routes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/kibzrael/raelelectronics/catalog/data"
	"github.com/kibzrael/raelelectronics/catalog/utils"
)

type LaptopCard struct {
	Uid   string `json:"uid"`
	Brand string `json:"brand"`
	Name  string `json:"name"`
	// Launched date YYYY-MM-DD
	Launched  string
	Thumbnail string `json:"thumbnail"`
	// Price in dollars
	PriceMin float64 `db:"price_min"`
	// Price in dollars
	PriceMax float64 `db:"price_max"`
	Colors   string  `json:"colors"`
	Size     float64 `json:"size"`
	Cpu      string  `json:"cpu"`
	Cores    int64   `json:"cores"`
	// Megahertz (MHz)
	BaseSpeed float64 `json:"base_speed" db:"base_speed"`
	// Gigabytes (GB)
	Memory int64 `json:"memory"`
	// Gigabytes (GB)
	Storage int64 `json:"storage"`
	// Hours
	BatteryLife float64 `db:"battery_life" json:"battery_life"`
	Featured    float64
}

func FeedHandler(ctx context.Context, res http.ResponseWriter, req *http.Request) {
	db := ctx.Value(data.DB_CONTEXT).(*sqlx.DB)

	sort := utils.SortQuery(req)
	filters := utils.FilterQuery(req)

	laptops := []LaptopCard{}

	sql := fmt.Sprintf(`
		SELECT * FROM
			(SELECT
				DISTINCT ON (laptops.uid)  laptops.uid,
				laptops.brand, laptops.name, laptops.thumbnail, laptops.price_min, laptops.price_max, laptops.launched,
				displays.size, chassis.colors, concat_ws(' ', cpus.architecture, cpus.model) as cpu, cpus.cores,
				cpus.base_speed, memorys.size as memory, storages.capacity as storage, batteries.life as battery_life,
				((memorys.size + (cpus.cores * cpus.base_speed)) / laptops.price_min) as featured
			FROM laptops
				LEFT JOIN chassis on laptops.uid = chassis.laptop
				LEFT JOIN displays on laptops.uid = displays.laptop
				LEFT JOIN cpus on laptops.uid = cpus.laptop
				LEFT JOIN memorys on laptops.uid = memorys.laptop
				LEFT JOIN storages on laptops.uid = storages.laptop
				LEFT JOIN batteries on laptops.uid = batteries.laptop
			%s
			ORDER BY laptops.uid, memorys.size ASC, storages.capacity ASC) l
		ORDER BY %s
		LIMIT 20 OFFSET 0
	`, filters, sort)

	rows, err := db.Queryx(sql)
	if err != nil {
		utils.ApiPanic(&res, &err)
		return
	}

	for rows.Next() {
		laptop := LaptopCard{}
		if err := rows.StructScan(&laptop); err != nil {
			utils.ApiPanic(&res, &err)
			return
		}
		laptops = append(laptops, laptop)
	}

	response := map[string]interface{}{"message": "Feed fetched successfully", "laptops": laptops}
	utils.JsonResponse(&res, response)
}
