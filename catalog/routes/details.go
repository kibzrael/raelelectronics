package routes

import (
	"context"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/kibzrael/raelelectronics/catalog/data"
	"github.com/kibzrael/raelelectronics/catalog/utils"
)

type ComponentFunction func(wg *sync.WaitGroup, db *sqlx.DB, l *data.Laptop)

func DetailsHandler(ctx context.Context, res http.ResponseWriter, req *http.Request) {
	start := time.Now()

	db := ctx.Value(data.DB_CONTEXT).(*sqlx.DB)
	wg := sync.WaitGroup{}

	var laptop data.Laptop

	err := db.Get(&laptop, "SELECT * FROM laptops WHERE uid=$1", req.PathValue("uid"))
	if err != nil {
		utils.ApiPanic(&res, &err)
		return
	}

	go func() {
		rows, _ := db.Queryx("SELECT url FROM images WHERE laptop=$1", laptop.Uid)
		for rows.Next() {
			var url string
			rows.Scan(&url)
			laptop.Images = append(laptop.Images, url)
		}
	}()

	components := []ComponentFunction{data.FetchChassis, data.FetchCpus, data.FetchGpus, data.FetchDisplays, data.FetchMemorys, data.FetchStorages, data.FetchWirelessCards, data.FetchMotherboards}

	for _, c := range components {
		wg.Add(1)
		go c(&wg, db, &laptop)
	}

	wg.Wait()

	response := map[string]interface{}{"message": "Laptop fetched successfully", "laptop": laptop}
	utils.JsonResponse(&res, response)
	log.Println("Time taken:", time.Since(start))
}
