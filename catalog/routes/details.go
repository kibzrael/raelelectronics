package routes

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/kibzrael/raelelectronics/catalog/data"
	"github.com/kibzrael/raelelectronics/catalog/utils"
	c "github.com/kibzrael/raelelectronics/common/api/catalog"
)

type ComponentFunction func(wg *sync.WaitGroup, db *sqlx.DB, l *data.Laptop)

func DetailsHandler(ctx context.Context, req *c.DetailsRequest) (res *c.Laptop, err error) {
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

	response := utils.ParseLaptop(laptop)
	res = &response
	log.Println("Time taken:", time.Since(start))
	return
}
