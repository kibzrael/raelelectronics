package routes

import (
	"context"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/kibzrael/raelelectronics/catalog/data"
	"github.com/kibzrael/raelelectronics/catalog/utils"
	c "github.com/kibzrael/raelelectronics/common/api/catalog"
)

type ComponentFunction func(wg *sync.WaitGroup, db *sqlx.DB, l *data.Laptop)

func DetailsHandler(ctx context.Context, req *c.DetailsRequest) (res *c.DetailsResponse, err error) {
	start := time.Now()

	db := ctx.Value(data.DB_CONTEXT).(*sqlx.DB)
	wg := sync.WaitGroup{}

	var laptop data.Laptop

	err = db.Get(&laptop, "SELECT * FROM laptops WHERE uid=$1", req.Uid)
	if err != nil {
		return
	}

	wg.Add(1)
	go func() {
		rows, _ := db.Queryx("SELECT url FROM images WHERE laptop=$1", laptop.Uid)
		for rows.Next() {
			var url string
			rows.Scan(&url)
			laptop.Images = append(laptop.Images, url)
		}
		wg.Done()
	}()

	components := []ComponentFunction{data.FetchChassis, data.FetchCpus, data.FetchGpus, data.FetchDisplays, data.FetchMemorys, data.FetchStorages, data.FetchWirelessCards, data.FetchMotherboards, data.FetchBatteries}

	for _, c := range components {
		wg.Add(1)
		go c(&wg, db, &laptop)
	}

	wg.Wait()

	res = &c.DetailsResponse{}
	response := utils.ParseLaptop(laptop)
	res.Laptop = &response
	fetchRelated(db, res)

	log.Println("Time taken:", time.Since(start))
	return
}

func fetchRelated(db *sqlx.DB, res *c.DetailsResponse) error {
	cpus := []string{}
	for _, cpu := range res.Laptop.CPU {
		cpus = append(cpus, fmt.Sprintf("'%s %s'", cpu.Architecture, cpu.Model))
	}

	sql := fmt.Sprintf(`
		SELECT * FROM
			(SELECT
				DISTINCT ON (laptops.uid)  laptops.uid,
				laptops.name, laptops.thumbnail, laptops.price_min, laptops.brand,
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
		WHERE l.brand = '%s' OR l.cpu IN (%s)
		ORDER BY l.featured DESC
		LIMIT 4
	`, res.Laptop.Brand, strings.Join(cpus, ","))

	rows, _ := db.Queryx(sql)
	for rows.Next() {
		related := data.LaptopCard{}
		if err := rows.StructScan(&related); err != nil {
			return err
		}

		response := utils.ParseLaptopCard(related)
		res.Related = append(res.Related, &response)
	}
	return nil
}
