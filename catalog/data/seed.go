package data

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"slices"
	"strconv"
	"strings"
	"sync"

	"github.com/jmoiron/sqlx"
)

var BRANDS []string = []string{"Lenovo", "HP", "Dell", "Apple", "Samsung", "Asus", "MSI", "Acer", "Microsoft", "Alienware", "Huawei", "LG", "Gigabyte", "Razer", "Framework", "Fujitsu", "Xiaomi", "Google"}

func SeedLaptops(ctx context.Context, res http.ResponseWriter, req *http.Request) {
	log.Println("Seeding Laptops")
	wg := sync.WaitGroup{}
	ctx = context.WithValue(ctx, WG_CONTEXT, &wg)

	for _, b := range BRANDS {
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

func (l *Laptop) Serialize(data map[string]interface{}) {
	modelInfoList := data["model_info"].([]interface{})
	modelInfo := modelInfoList[0].(map[string]interface{})
	l.NotebId = int64(modelInfo["id"].(float64))
	l.Name = modelInfo["noteb_name"].(string)
	nameSplit := strings.Split(modelInfo["noteb_name"].(string), " ")
	l.Brand = nameSplit[0]
	moreResources := data["model_resources"].(map[string]interface{})
	l.Thumbnail = moreResources["thumbnail"].(string)
	link := moreResources["official_link"]
	if link != nil {
		l.Link = link.(string)
	}
	l.Launched, _ = moreResources["launch_date"].(string)
	l.PriceMin, _ = strconv.ParseFloat(data["config_price_min"].(string), 64)
	l.PriceMax, _ = strconv.ParseFloat(data["config_price_max"].(string), 64)
	for key, val := range moreResources {
		if strings.HasPrefix(key, "image_") {
			l.Images = append(l.Images, val.(string))
		}
	}
	cpus := data["cpu"].(map[string]interface{})
	for _, val := range cpus {
		_, ok := val.(map[string]interface{})
		if ok {
			var cpu CPU
			b, err := json.Marshal(val)
			if err != nil {
				panic(err)
			}
			json.Unmarshal(b, &cpu)
			l.CPU = append(l.CPU, cpu)
		}
	}
	gpus := data["gpu"].(map[string]interface{})
	for _, val := range gpus {
		_, ok := val.(map[string]interface{})
		if ok {
			var gpu GPU
			b, err := json.Marshal(val)
			if err != nil {
				panic(err)
			}
			json.Unmarshal(b, &gpu)
			l.GPU = append(l.GPU, gpu)
		}
	}
	displays := data["display"].(map[string]interface{})
	for _, val := range displays {
		_, ok := val.(map[string]interface{})
		if ok {
			var display Display
			b, err := json.Marshal(val)
			if err != nil {
				panic(err)
			}
			json.Unmarshal(b, &display)
			displayMap, ok := val.(map[string]interface{})
			if ok {
				display.Touchscreen = displayMap["touch"] != "no"
			}
			l.Display = append(l.Display, display)
		}
	}
	memorys := data["memory"].(map[string]interface{})
	for _, val := range memorys {
		_, ok := val.(map[string]interface{})
		if ok {
			var memory Memory
			b, err := json.Marshal(val)
			if err != nil {
				panic(err)
			}
			json.Unmarshal(b, &memory)
			l.Memory = append(l.Memory, memory)
		}
	}
	storages := data["primary_storage"].(map[string]interface{})
	for _, val := range storages {
		_, ok := val.(map[string]interface{})
		if ok {
			var storage Storage
			b, err := json.Marshal(val)
			if err != nil {
				panic(err)
			}
			json.Unmarshal(b, &storage)
			l.Storage = append(l.Storage, storage)
		}
	}
	batterys := data["battery"].(map[string]interface{})
	for _, val := range batterys {
		_, ok := val.(map[string]interface{})
		if ok {
			var battery Battery
			b, err := json.Marshal(val)
			if err != nil {
				panic(err)
			}
			json.Unmarshal(b, &battery)
			battery.Life, _ = strconv.ParseFloat(data["battery_life_raw"].(string), 64)
			l.Battery = append(l.Battery, battery)
		}
	}
	chassiss := data["chassis"].(map[string]interface{})
	for _, val := range chassiss {
		chassisMap, ok := val.(map[string]interface{})
		if ok {
			var chassis Chassis
			b, err := json.Marshal(val)
			if err != nil {
				panic(err)
			}
			json.Unmarshal(b, &chassis)

			if chassisMap["colors"] != nil {
				chassis.Colors = strings.Split(chassisMap["colors"].(string), ",")
			}
			periferals := chassisMap["peripheral_interfaces"].(string)
			video := chassisMap["video_interfaces"]
			if video != nil {
				periferals = strings.Join([]string{periferals, video.(string)}, ",")
			}
			ports := strings.Split(periferals, ",")
			for _, port := range ports {
				splits := strings.Split(port, "X")
				if len(splits) > 1 {
					count, _ := strconv.ParseUint(strings.Trim(splits[0], " "), 10, 64)
					chassis.Ports = append(chassis.Ports, Port{Count: count, Type: strings.Trim(splits[1], " ")})
				}
			}
			l.Chassis = append(l.Chassis, chassis)
		}
	}
	wireless := data["wireless"].(map[string]interface{})
	for _, val := range wireless {
		_, ok := val.(map[string]interface{})
		if ok {
			var wirelessCard WirelessCard
			b, err := json.Marshal(val)
			if err != nil {
				panic(err)
			}
			json.Unmarshal(b, &wirelessCard)
			l.WirelessCard = append(l.WirelessCard, wirelessCard)
		}
	}
	motherboards := data["motherboard"].(map[string]interface{})
	for _, val := range motherboards {
		_, ok := val.(map[string]interface{})
		if ok {
			var motherboard Motherboard
			b, err := json.Marshal(val)
			if err != nil {
				panic(err)
			}
			json.Unmarshal(b, &motherboard)
			l.Motherboard = append(l.Motherboard, motherboard)
		}
	}
}
