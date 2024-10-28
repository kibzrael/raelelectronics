package routes

import (
	"context"
	"fmt"
	"math"

	c "github.com/kibzrael/raelelectronics/common/api/catalog"

	"github.com/jmoiron/sqlx"
	"github.com/kibzrael/raelelectronics/catalog/data"
	"github.com/kibzrael/raelelectronics/catalog/utils"
)

func FeedHandler(ctx context.Context, req *c.FeedRequest) (res []*c.LaptopCard, pages int64, err error) {
	db := ctx.Value(data.DB_CONTEXT).(*sqlx.DB)

	sort := utils.SortQuery(req)
	filters := utils.FilterQuery(req)

	page := req.Page

	sql := fmt.Sprintf(`
		SELECT *, count(*) OVER() AS count FROM
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
		LIMIT 20 OFFSET %v
	`, filters, sort, (page-1)*20)

	rows, err := db.Queryx(sql)
	if err != nil {
		return
	}

	for rows.Next() {
		laptop := data.LaptopCard{}
		if err := rows.StructScan(&laptop); err != nil {
			return res, 1, err
		}

		response := utils.ParseLaptopCard(laptop)
		res = append(res, &response)
		if pages == 0 {
			pages = int64(math.Ceil(float64(laptop.Count) / 20.0))
		}
	}

	return
}
