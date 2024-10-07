package data

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"slices"
	"strings"
	"sync"

	"github.com/jmoiron/sqlx"
	"github.com/kibzrael/raelelectronics/catalog/utils"
)

func SeedLaptops(ctx context.Context, res http.ResponseWriter, req *http.Request) {
	log.Println("Seeding Laptops")
	wg := sync.WaitGroup{}
	ctx = context.WithValue(ctx, WG_CONTEXT, &wg)

	for _, b := range utils.BRANDS {
		wg.Add(1)
		go ListLaptops(ctx, b)
	}

	wg.Wait()
	fmt.Fprintln(res, "Success")
}

func ListLaptops(ctx context.Context, brand string) {
	url := "https://noteb.com/api/webservice.php"
	payload := strings.NewReader(fmt.Sprintf(`apikey=%s&method=list_models&param[model_id]=&param[model_name]=%s`, API_KEY, brand))

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var data map[string]map[string]map[string]interface{}
	json.Unmarshal(body, &data)

	laptops := data["result"]
	var laptopIds []uint
	for _, l := range laptops {
		details := l["model_info"].([]interface{})
		for _, d := range details {
			info := d.(map[string]interface{})
			laptopIds = append(laptopIds, uint(info["id"].(float64)))
		}
	}

	var wg = ctx.Value(WG_CONTEXT).(*sync.WaitGroup)
	wgDetail := sync.WaitGroup{}
	ctx = context.WithValue(ctx, WG_DETAIL_CONTEXT, &wgDetail)

	db := ctx.Value(DB_CONTEXT).(*sqlx.DB)
	rows, err := db.Queryx("SELECT noteb_id FROM laptops WHERE brand = $1", brand)
	if err != nil {
		fmt.Println(err)
		return
	}
	for rows.Next() {
		var value uint
		if err = rows.Scan(&value); err == nil {
			i := slices.Index(laptopIds, value)
			if i != -1 {
				laptopIds = append(laptopIds[:i], laptopIds[i+1:]...)
			}
		}
	}

	for _, id := range laptopIds {
		wgDetail.Add(1)
		go LaptopDetails(ctx, id)
	}

	wgDetail.Wait()
	wg.Done()
}

func LaptopDetails(ctx context.Context, id uint) {
	url := "https://noteb.com/api/webservice.php"
	payload := strings.NewReader(fmt.Sprintf(`apikey=%s&method=get_model_info_all&param[model_id]=%v`, API_KEY, id))

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var laptopInfo map[string]map[string]map[string]interface{}
	json.Unmarshal(body, &laptopInfo)

	result := laptopInfo["result"]["0"]

	if len(result) > 0 {
		laptop := Laptop{}
		laptop.Serialize(result)
		laptop.Save(ctx)
	}

	var wgDetail = ctx.Value(WG_DETAIL_CONTEXT).(*sync.WaitGroup)
	wgDetail.Done()
}
