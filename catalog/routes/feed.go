package routes

import (
	"context"
	"fmt"
	"strings"

	c "github.com/kibzrael/raelelectronics/common/api/catalog"

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

func FeedHandler(ctx context.Context, req *c.FeedRequest) (res []*c.LaptopCard, err error) {
	db := ctx.Value(data.DB_CONTEXT).(*sqlx.DB)

	sort := utils.SortQuery(req)
	filters := utils.FilterQuery(req)

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
		return
	}

	for rows.Next() {
		laptop := LaptopCard{}
		if err := rows.StructScan(&laptop); err != nil {
			return res, err
		}

		response := &c.LaptopCard{
			Uid:         laptop.Uid,
			Name:        laptop.Name,
			Brand:       laptop.Brand,
			Thumbnail:   laptop.Thumbnail,
			Launched:    laptop.Thumbnail,
			PriceMin:    laptop.PriceMin,
			PriceMax:    laptop.PriceMax,
			Colors:      strings.Split(laptop.Colors, ","),
			Size:        laptop.Size,
			Cpu:         laptop.Cpu,
			Cores:       laptop.Cores,
			BaseSpeed:   laptop.BaseSpeed,
			Memory:      laptop.Memory,
			Storage:     laptop.Storage,
			BatteryLife: laptop.BatteryLife,
		}
		res = append(res, response)
	}

	return
}
