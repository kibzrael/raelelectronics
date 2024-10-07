package data

import (
	"context"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

const Schema = `
	CREATE TABLE IF NOT EXISTS laptops (
		uid					varchar(80) PRIMARY KEY,
		noteb_id			integer NOT NULL UNIQUE,
		name				varchar(255) NOT NULL,
		brand				varchar(255) NOT NULL,
		thumbnail			varchar(255) NOT NULL,
		link				varchar(255) NULL,
		launched			Date NULL,
		price_min			float NOT NULL,
		price_max			float NOT NULL,

		created_at			TIMESTAMP NOT NULL DEFAULT NOW()
	);

	CREATE TABLE IF NOT EXISTS cpus (
		uid					varchar(80) PRIMARY KEY,
		laptop				varchar(80) REFERENCES laptops ON DELETE CASCADE,
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
		laptop				varchar(80) REFERENCES laptops ON DELETE CASCADE,
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
		laptop				varchar(80) REFERENCES laptops ON DELETE CASCADE,
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
		laptop				varchar(80) REFERENCES laptops ON DELETE CASCADE,
		size				integer NOT NULL,
		speed				float NULL,
		type				varchar(255) NULL,

		created_at			TIMESTAMP NOT NULL DEFAULT NOW()
	);

	CREATE TABLE IF NOT EXISTS storages (
		uid					varchar(80) PRIMARY KEY,
		laptop				varchar(80) REFERENCES laptops ON DELETE CASCADE,
		capacity			integer NOT NULL,
		speed				integer NULL,
		model				varchar(255) NULL,

		created_at			TIMESTAMP NOT NULL DEFAULT NOW()
	);

	CREATE TABLE IF NOT EXISTS wireless_cards (
		uid					varchar(80) PRIMARY KEY,
		laptop				varchar(80) REFERENCES laptops ON DELETE CASCADE,
		speed				integer NULL,
		model				varchar(255) NULL,
		other_info			varchar(255) NULL,

		created_at			TIMESTAMP NOT NULL DEFAULT NOW()
	);

	CREATE TABLE IF NOT EXISTS motherboards (
		uid					varchar(80) PRIMARY KEY,
		laptop				varchar(80) REFERENCES laptops ON DELETE CASCADE,
		ram					varchar(255) NULL,
		storage				varchar(255) NULL,

		created_at			TIMESTAMP NOT NULL DEFAULT NOW()
	);

	CREATE TABLE IF NOT EXISTS chassis (
		uid					varchar(80) PRIMARY KEY,
		laptop				varchar(80) REFERENCES laptops ON DELETE CASCADE,
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
		chassis				varchar(80) REFERENCES chassis ON DELETE CASCADE,
		count				integer NULL,
		type				varchar(255) NULL,

		created_at			TIMESTAMP NOT NULL DEFAULT NOW()
	);

	CREATE TABLE IF NOT EXISTS batteries (
		uid					varchar(80) PRIMARY KEY,
		laptop				varchar(80) REFERENCES laptops ON DELETE CASCADE,
		capacity			float NOT NULL,
		type				varchar(255) NULL,
		life				float NULL,
		other_info			varchar(255) NULL,

		created_at			TIMESTAMP NOT NULL DEFAULT NOW()
	);

	CREATE TABLE IF NOT EXISTS images (
		uid					varchar(80) PRIMARY KEY,
		laptop				varchar(80) REFERENCES laptops ON DELETE CASCADE,
		url					varchar(255) NOT NULL,

		created_at			TIMESTAMP NOT NULL DEFAULT NOW()
	);
`

type LaptopComponent interface {
	Save(tx *sqlx.Tx, l *Laptop) error
}

type CPU struct {
	Uid          string
	Laptop       string
	Model        string `json:"model"`
	Architecture string `json:"prod"`
	// Nanometres (nm)
	Lithography uint64 `json:"lithography,string"`
	// Megabytes (MB)
	Cache float64 `json:"cache,string"`
	// Megahertz (MHz)
	BaseSpeed float64 `json:"base_speed,string" db:"base_speed"`
	// Megahertz (MHz)
	BoostSpeed      float64   `json:"boost_speed,string" db:"boost_speed"`
	Cores           int64     `json:"cores,string"`
	Tdp             int64     `json:"tdp,string"`
	IntegratedVideo string    `json:"integrated_video" db:"integrated_video"`
	OtherInfo       string    `json:"other_info" db:"other_info"`
	CreatedAt       time.Time `db:"created_at"`
}

func (c *CPU) Save(tx *sqlx.Tx, l *Laptop) error {
	c.Uid = uuid.New().String()
	c.Laptop = l.Uid
	_, err := tx.NamedExec(`
		INSERT INTO cpus (uid, laptop, model, architecture, lithography, cache, base_speed, boost_speed, cores, tdp, integrated_video, other_info) 
		VALUES (:uid, :laptop, :model, :architecture, :lithography, :cache, :base_speed, :boost_speed, :cores, :tdp, :integrated_video, :other_info)
	`, c)
	return err
}

type GPU struct {
	Uid          string
	Laptop       string
	Producer     string `json:"prod"`
	Model        string `json:"model"`
	Architecture string `json:"architecture"`
	// Nanometres (nm)
	Lithography uint64 `json:"lithography,string"`
	// Megahertz (MHz)
	BaseSpeed float64 `json:"base_speed,string" db:"base_speed"`
	// Megahertz (MHz)
	BoostSpeed float64 `json:"boost_speed,string" db:"boost_speed"`
	// Megabytes (MB)
	MemorySize int64 `json:"memory_size,string" db:"memory_size"`
	// Megahertz (MHz)
	MemorySpeed float64   `json:"memory_speed,string" db:"memory_speed"`
	MemoryType  string    `json:"memory_type" db:"memory_type"`
	Tdp         int64     `json:"tdp,string"`
	OtherInfo   string    `json:"other_info" db:"other_info"`
	CreatedAt   time.Time `db:"created_at"`
}

func (g *GPU) Save(tx *sqlx.Tx, l *Laptop) error {
	g.Uid = uuid.New().String()
	g.Laptop = l.Uid
	_, err := tx.NamedExec(`
		INSERT INTO gpus (uid, laptop, model, architecture, lithography, producer, base_speed, boost_speed, memory_size, tdp, memory_speed, memory_type, other_info) 
		VALUES (:uid, :laptop, :model, :architecture, :lithography, :producer, :base_speed, :boost_speed, :memory_size, :tdp, :memory_speed, :memory_type, :other_info)
	`, g)
	return err
}

type Display struct {
	Uid    string
	Laptop string
	// Inches
	Size float64 `json:"size,string"`
	// Pixels
	HResolution int64 `json:"horizontal_resolution,string" db:"h_resolution"`
	// Pixels
	VResolution int64  `json:"vertical_resolution,string" db:"v_resolution"`
	Type        string `json:"type"`
	// Percentage
	SRGB        float64 `json:"sRGB,string" db:"srgb"`
	Touchscreen bool
	OtherInfo   string    `json:"other_info" db:"other_info"`
	CreatedAt   time.Time `db:"created_at"`
}

func (d *Display) Save(tx *sqlx.Tx, l *Laptop) error {
	d.Uid = uuid.New().String()
	d.Laptop = l.Uid
	_, err := tx.NamedExec(`
		INSERT INTO displays (uid, laptop, size, h_resolution, v_resolution, type, srgb, touchscreen, other_info) 
		VALUES (:uid, :laptop, :size, :h_resolution, :v_resolution, :type, :srgb, :touchscreen, :other_info)
	`, d)
	return err
}

type Memory struct {
	Uid    string
	Laptop string
	// Gigabytes (GB)
	Size int64 `json:"size,string"`
	// MB/s
	Speed     float64   `json:"speed,string"`
	Type      string    `json:"type"`
	CreatedAt time.Time `db:"created_at"`
}

func (m *Memory) Save(tx *sqlx.Tx, l *Laptop) error {
	m.Uid = uuid.New().String()
	m.Laptop = l.Uid
	_, err := tx.NamedExec(`
		INSERT INTO memorys (uid, laptop, size, speed, type) 
		VALUES (:uid, :laptop, :size, :speed, :type)
	`, m)
	return err
}

type Storage struct {
	Uid    string
	Laptop string
	// Gigabytes (GB)
	Capacity int64 `json:"cap,string"`
	// MB/s
	Speed     int64     `json:"read_speed,string"`
	Model     string    `json:"model"`
	CreatedAt time.Time `db:"created_at"`
}

func (s *Storage) Save(tx *sqlx.Tx, l *Laptop) error {
	s.Uid = uuid.New().String()
	s.Laptop = l.Uid
	_, err := tx.NamedExec(`
		INSERT INTO storages (uid, laptop, capacity, speed, model) 
		VALUES (:uid, :laptop, :capacity, :speed, :model)
	`, s)
	return err
}

type WirelessCard struct {
	Uid    string
	Laptop string
	Model  string `json:"model"`
	// MB/s
	Speed     int64     `json:"speed,string"`
	OtherInfo string    `json:"other_info" db:"other_info"`
	CreatedAt time.Time `db:"created_at"`
}

func (w *WirelessCard) Save(tx *sqlx.Tx, l *Laptop) error {
	w.Uid = uuid.New().String()
	w.Laptop = l.Uid
	_, err := tx.NamedExec(`
		INSERT INTO wireless_cards (uid, laptop, model, speed, other_info) 
		VALUES (:uid, :laptop, :model, :speed, :other_info)
	`, w)
	return err
}

type Motherboard struct {
	Uid       string
	Laptop    string
	Ram       string    `json:"ram_slots"`
	Storage   string    `json:"storage_slots"`
	CreatedAt time.Time `db:"created_at"`
}

func (m *Motherboard) Save(tx *sqlx.Tx, l *Laptop) error {
	m.Uid = uuid.New().String()
	m.Laptop = l.Uid
	_, err := tx.NamedExec(`
		INSERT INTO motherboards (uid, laptop, ram, storage) 
		VALUES (:uid, :laptop, :ram, :storage)
	`, m)
	return err
}

type Port struct {
	Uid       string
	Chassis   string
	Count     uint64
	Type      string
	CreatedAt time.Time `db:"created_at"`
}

func (p *Port) Save(tx *sqlx.Tx, c *Chassis) error {
	p.Uid = uuid.New().String()
	p.Chassis = c.Uid
	_, err := tx.NamedExec(`
		INSERT INTO ports (uid, chassis, count, type) 
		VALUES (:uid, :chassis, :count, :type)
	`, p)
	return err
}

type Chassis struct {
	Uid    string
	Laptop string
	// Centimetres (cm)
	Height float64 `json:"height_cm,string"`
	// Centimetres (cm)
	Depth float64 `json:"depth_cm,string"`
	// Centimetres (cm)
	Width float64 `json:"width_cm,string"`
	// Kilograms (kg)
	Weight float64 `json:"weight_kg,string"`
	Colors []string
	Ports  []Port
	// Megapixels (MP)
	Webcam    float64   `json:"webcam_mp,string"`
	OtherInfo string    `json:"other_info" db:"other_info"`
	CreatedAt time.Time `db:"created_at"`
}

func (c *Chassis) Save(tx *sqlx.Tx, l *Laptop) error {
	c.Uid = uuid.New().String()
	c.Laptop = l.Uid
	_, err := tx.Exec(`
		INSERT INTO chassis (uid, laptop, height, depth, width, weight, colors, webcam, other_info) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`, c.Uid, c.Laptop, c.Height, c.Depth, c.Width, c.Weight, strings.Join(c.Colors, ","), c.Webcam, c.OtherInfo)
	for _, port := range c.Ports {
		err = port.Save(tx, c)
	}
	return err
}

type Battery struct {
	Uid    string
	Laptop string
	// Watts per hour (WH)
	Capacity float64 `json:"capacity,string"`
	Type     string  `json:"cell_type"`
	// Hours
	Life      float64
	OtherInfo string    `json:"other_info" db:"other_info"`
	CreatedAt time.Time `db:"created_at"`
}

func (b *Battery) Save(tx *sqlx.Tx, l *Laptop) error {
	b.Uid = uuid.New().String()
	b.Laptop = l.Uid
	_, err := tx.NamedExec(`
		INSERT INTO batteries (uid, laptop, capacity, type, life, other_info) 
		VALUES (:uid, :laptop, :capacity, :type, :life, :other_info)
	`, b)
	return err
}

type Laptop struct {
	Uid       string
	NotebId   int64  `db:"noteb_id"`
	Name      string `db:"name"`
	Brand     string
	Thumbnail string
	Images    []string
	Link      string
	// Launched date YYYY-MM-DD
	Launched string
	// Price in dollars
	PriceMin float64 `db:"price_min"`
	// Price in dollars
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
	CreatedAt    time.Time `db:"created_at"`
}

func (l *Laptop) Save(ctx context.Context) error {
	db := ctx.Value(DB_CONTEXT).(*sqlx.DB)

	tx := db.MustBegin()
	l.Uid = uuid.New().String()
	_, err := tx.NamedExec(`
		INSERT INTO laptops (uid, noteb_id, name, brand, thumbnail, link, launched, price_min, price_max) 
		VALUES (:uid, :noteb_id, :name, :brand, :thumbnail, :link, :launched, :price_min, :price_max)
	`, l)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	for _, img := range l.Images {
		_, err = tx.Exec(`
			INSERT INTO images (uid, laptop, url) VALUES ($1, $2, $3)
		`, uuid.New().String(), l.Uid, img)
	}

	// fields := []any{l.CPU, l.GPU, l.Display, l.Memory, l.Storage, l.WirelessCard, l.Motherboard, l.Chassis, l.Battery}
	// for _, f := range fields {
	// 	components := f.([]LaptopComponent)
	// 	for _, c := range components {
	// 		err = c.Save(tx, l)
	// 	}
	// }

	for _, cpu := range l.CPU {
		err = cpu.Save(tx, l)
	}
	for _, gpu := range l.GPU {
		err = gpu.Save(tx, l)
	}
	for _, display := range l.Display {
		err = display.Save(tx, l)
	}
	for _, memory := range l.Memory {
		err = memory.Save(tx, l)
	}
	for _, storage := range l.Storage {
		err = storage.Save(tx, l)
	}
	for _, wireless_card := range l.WirelessCard {
		err = wireless_card.Save(tx, l)
	}
	for _, motherboard := range l.Motherboard {
		err = motherboard.Save(tx, l)
	}
	for _, chassis := range l.Chassis {
		err = chassis.Save(tx, l)
	}
	for _, battery := range l.Battery {
		err = battery.Save(tx, l)
	}

	if err != nil {
		_ = tx.Rollback()
		return err
	}
	err = tx.Commit()
	return err
}
