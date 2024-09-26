package data

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

const Schema = `
	CREATE TABLE IF NOT EXISTS laptops (
		uid					varchar(80) PRIMARY KEY,
		noteb_id			integer NULL,
		name				varchar(255) NOT NULL,
		brand				varchar(255) NOT NULL,
		thumbnail			varchar(255) NOT NULL,
		link				varchar(255) NULL,
		launched			TIMESTAMP NULL,
		price_min			float NOT NULL,
		price_max			float NOT NULL,

		created_at			TIMESTAMP NOT NULL DEFAULT NOW()
	);

	CREATE TABLE IF NOT EXISTS cpus (
		uid					varchar(80) PRIMARY KEY,
		laptop				varchar(80) REFERENCES laptops,
		model				varchar(255) NOT NULL,
		architecture		varchar(255) NULL,
		lithography			integer NULL,
		cache				float NULL,
		base_speed			float NULL,
		boost_speed			float NULL,
		cores				integer NULL,
		tdp					integer NULL,
		integrated_video	varchar(255) NULL,
		other_info			varchar(255) NULL,

		created_at			TIMESTAMP NOT NULL DEFAULT NOW()
	);

	CREATE TABLE IF NOT EXISTS gpus (
		uid					varchar(80) PRIMARY KEY,
		laptop				varchar(80) REFERENCES laptops,
		producer			varchar(255) NULL,
		model				varchar(255) NOT NULL,
		architecture		varchar(255) NULL,
		lithography			integer NULL,
		base_speed			float NULL,
		boost_speed			float NULL,
		memory_size			integer NULL,
		memory_speed		float NULL,
		memory_type			varchar(255) NULL,
		tdp					integer NULL,
		other_info			varchar(255) NULL,

		created_at			TIMESTAMP NOT NULL DEFAULT NOW()
	);

	CREATE TABLE IF NOT EXISTS displays (
		uid					varchar(80) PRIMARY KEY,
		laptop				varchar(80) REFERENCES laptops,
		size				float NOT NULL,
		h_resolution		integer NOT NULL,
		v_resolution		integer NOT NULL,
		type				varchar(255) NULL,
		srgb				float NULL,
		touchscreen			boolean NULL,
		other_info			varchar(255) NULL,

		created_at			TIMESTAMP NOT NULL DEFAULT NOW()
	);

	CREATE TABLE IF NOT EXISTS memorys (
		uid					varchar(80) PRIMARY KEY,
		laptop				varchar(80) REFERENCES laptops,
		size				integer NOT NULL,
		speed				float NULL,
		type				varchar(255) NULL,

		created_at			TIMESTAMP NOT NULL DEFAULT NOW()
	);

	CREATE TABLE IF NOT EXISTS storages (
		uid					varchar(80) PRIMARY KEY,
		laptop				varchar(80) REFERENCES laptops,
		capacity			integer NOT NULL,
		speed				integer NULL,
		model				varchar(255) NULL,

		created_at			TIMESTAMP NOT NULL DEFAULT NOW()
	);

	CREATE TABLE IF NOT EXISTS wireless_cards (
		uid					varchar(80) PRIMARY KEY,
		laptop				varchar(80) REFERENCES laptops,
		speed				integer NULL,
		model				varchar(255) NULL,
		other_info			varchar(255) NULL,

		created_at			TIMESTAMP NOT NULL DEFAULT NOW()
	);

	CREATE TABLE IF NOT EXISTS motherboards (
		uid					varchar(80) PRIMARY KEY,
		laptop				varchar(80) REFERENCES laptops,
		ram					varchar(255) NULL,
		storage				varchar(255) NULL,

		created_at			TIMESTAMP NOT NULL DEFAULT NOW()
	);

	CREATE TABLE IF NOT EXISTS chassis (
		uid					varchar(80) PRIMARY KEY,
		laptop				varchar(80) REFERENCES laptops,
		height				float NOT NULL,
		depth				float NOT NULL,
		width				float NOT NULL,
		weight				float NOT NULL,
		colors				varchar(255) NULL,
		webcam				float NULL,
		other_info			varchar(255) NULL,

		created_at			TIMESTAMP NOT NULL DEFAULT NOW()
	);

	CREATE TABLE IF NOT EXISTS ports (
		uid					varchar(80) PRIMARY KEY,
		laptop				varchar(80) REFERENCES laptops,
		chassis				varchar(80) REFERENCES chassis,
		count				integer NULL,
		type				varchar(255) NULL,

		created_at			TIMESTAMP NOT NULL DEFAULT NOW()
	);

	CREATE TABLE IF NOT EXISTS batteries (
		uid					varchar(80) PRIMARY KEY,
		laptop				varchar(80) REFERENCES laptops,
		capacity			float NOT NULL,
		type				varchar(255) NULL,
		life				float NULL,
		other_info			varchar(255) NULL,

		created_at			TIMESTAMP NOT NULL DEFAULT NOW()
	);

	CREATE TABLE IF NOT EXISTS images (
		uid					varchar(80) PRIMARY KEY,
		laptop				varchar(80) REFERENCES laptops,
		url					varchar(255) NOT NULL,

		created_at			TIMESTAMP NOT NULL DEFAULT NOW()
	);
`

type CPU struct {
	Uid             string
	Model           string  `json:"model"`
	Architecture    string  `json:"prod"`
	Lithography     uint64  `json:"lithography,string"`
	Cache           float64 `json:"cache,string"`
	BaseSpeed       float64 `json:"base_speed,string" db:"base_speed"`
	BoostSpeed      float64 `json:"boost_speed,string" db:"boost_speed"`
	Cores           int64   `json:"cores,string"`
	Tdp             int64   `json:"tdp,string"`
	IntegratedVideo string  `json:"integrated_video" db:"integrated_video"`
	OtherInfo       string  `json:"other_info" db:"other_info"`
}

type GPU struct {
	Uid          string
	Producer     string  `json:"prod"`
	Model        string  `json:"model"`
	Architecture string  `json:"architecture"`
	Lithography  uint64  `json:"lithography,string"`
	BaseSpeed    float64 `json:"base_speed,string" db:"base_speed"`
	BoostSpeed   float64 `json:"boost_speed,string" db:"boost_speed"`
	MemorySize   int64   `json:"memory_size,string" db:"memory_size"`
	MemorySpeed  float64 `json:"memory_speed,string" db:"memory_speed"`
	MemoryType   string  `json:"memory_type" db:"memory_type"`
	Tdp          int64   `json:"tdp,string"`
	OtherInfo    string  `json:"other_info" db:"other_info"`
}

type Display struct {
	Uid         string
	Size        float64 `json:"size,string"`
	HResolution int64   `json:"horizontal_resolution,string" db:"h_resolution"`
	VResolution int64   `json:"vertical_resolution,string" db:"v_resolution"`
	Type        string  `json:"type"`
	SRGB        float64 `json:"sRGB,string" db:"srgb"`
	Touchscreen bool
	OtherInfo   string `json:"other_info" db:"other_info"`
}

type Memory struct {
	Uid   string
	Size  int64   `json:"size,string"`
	Speed float64 `json:"speed,string"`
	Type  string  `json:"type"`
}

type Storage struct {
	Uid      string
	Capacity int64  `json:"cap,string"`
	Speed    int64  `json:"read_speed,string"`
	Model    string `json:"model"`
}

type WirelessCard struct {
	Uid       string
	Model     string `json:"model"`
	Speed     int64  `json:"speed,string"`
	OtherInfo string `json:"other_info" db:"other_info"`
}

type Motherboard struct {
	Uid     string
	Ram     string `json:"ram_slots"`
	Storage string `json:"storage_slots"`
}

type Port struct {
	Uid   string
	Count uint64
	Type  string
}

type Chassis struct {
	Uid       string
	Height    float64 `json:"height_cm,string"`
	Depth     float64 `json:"depth_cm,string"`
	Width     float64 `json:"width_cm,string"`
	Weight    float64 `json:"weight_kg,string"`
	Colors    []string
	ports     []Port
	Webcam    float64 `json:"webcam_mp,string"`
	OtherInfo string  `json:"other_info" db:"other_info"`
}

type Battery struct {
	Uid       string
	Capacity  float64 `json:"capacity,string"`
	Type      string  `json:"cell_type"`
	Life      float64
	OtherInfo string `json:"other_info" db:"other_info"`
}

type Laptop struct {
	Uid          string
	NotebId      int64 `db:"noteb_id"`
	Name         string
	Brand        string
	Thumbnail    string
	Images       []string
	Link         string
	Launched     time.Time
	PriceMin     float64 `db:"price_min"`
	PriceMax     float64 `db:"price_max"`
	CPU          []CPU
	GPU          []GPU
	Display      []Display
	Memory       []Memory
	Storage      []Storage
	WirelessCard []WirelessCard `db:"wireless_card"`
	Motherboard  []Motherboard
	Chassis      []Chassis
	Battery      []Battery
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
	l.Link = moreResources["official_link"].(string)
	l.Launched, _ = time.Parse("2006-01-02", moreResources["launch_date"].(string))
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

			chassis.Colors = strings.Split(chassisMap["colors"].(string), ",")
			ports := strings.Split(chassisMap["peripheral_interfaces"].(string), ",")
			for _, port := range ports {
				splits := strings.Split(port, "X")
				if len(splits) > 1 {
					count, _ := strconv.ParseUint(strings.Trim(splits[0], " "), 10, 64)
					chassis.ports = append(chassis.ports, Port{Count: count, Type: splits[1]})
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
