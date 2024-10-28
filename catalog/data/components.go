package data

import (
	"strings"
	"sync"

	"github.com/jmoiron/sqlx"
)

func FetchCpus(wg *sync.WaitGroup, db *sqlx.DB, l *Laptop) {
	rows, _ := db.Queryx("SELECT * FROM cpus WHERE laptop=$1", l.Uid)
	for rows.Next() {
		var cpu CPU
		rows.StructScan(&cpu)
		l.CPU = append(l.CPU, cpu)
	}
	wg.Done()
}

func FetchGpus(wg *sync.WaitGroup, db *sqlx.DB, l *Laptop) {
	rows, _ := db.Queryx("SELECT * FROM gpus WHERE laptop=$1", l.Uid)
	for rows.Next() {
		var gpu GPU
		rows.StructScan(&gpu)
		l.GPU = append(l.GPU, gpu)
	}
	wg.Done()
}

func FetchDisplays(wg *sync.WaitGroup, db *sqlx.DB, l *Laptop) {
	rows, _ := db.Queryx("SELECT * FROM displays WHERE laptop=$1", l.Uid)
	for rows.Next() {
		var display Display
		rows.StructScan(&display)
		l.Display = append(l.Display, display)
	}
	wg.Done()
}

func FetchMemorys(wg *sync.WaitGroup, db *sqlx.DB, l *Laptop) {
	rows, _ := db.Queryx("SELECT * FROM memorys WHERE laptop=$1", l.Uid)
	for rows.Next() {
		var memory Memory
		rows.StructScan(&memory)
		l.Memory = append(l.Memory, memory)
	}
	wg.Done()
}

func FetchStorages(wg *sync.WaitGroup, db *sqlx.DB, l *Laptop) {
	rows, _ := db.Queryx("SELECT * FROM storages WHERE laptop=$1", l.Uid)
	for rows.Next() {
		var storage Storage
		rows.StructScan(&storage)
		l.Storage = append(l.Storage, storage)
	}
	wg.Done()
}

func FetchWirelessCards(wg *sync.WaitGroup, db *sqlx.DB, l *Laptop) {
	rows, _ := db.Queryx("SELECT * FROM wireless_cards WHERE laptop=$1", l.Uid)
	for rows.Next() {
		var wireless_card WirelessCard
		rows.StructScan(&wireless_card)
		l.WirelessCard = append(l.WirelessCard, wireless_card)
	}
	wg.Done()
}

func FetchMotherboards(wg *sync.WaitGroup, db *sqlx.DB, l *Laptop) {
	rows, _ := db.Queryx("SELECT * FROM motherboards WHERE laptop=$1", l.Uid)
	for rows.Next() {
		var motherboard Motherboard
		rows.StructScan(&motherboard)
		l.Motherboard = append(l.Motherboard, motherboard)
	}
	wg.Done()
}

func FetchBatteries(wg *sync.WaitGroup, db *sqlx.DB, l *Laptop) {
	rows, _ := db.Queryx("SELECT * FROM batteries WHERE laptop=$1", l.Uid)
	for rows.Next() {
		var battery Battery
		rows.StructScan(&battery)
		l.Battery = append(l.Battery, battery)
	}
	wg.Done()
}

func FetchChassis(wg *sync.WaitGroup, db *sqlx.DB, l *Laptop) {
	portsWg := sync.WaitGroup{}
	rows, _ := db.Queryx("SELECT * FROM chassis WHERE laptop=$1", l.Uid)
	for rows.Next() {
		var chassis Chassis
		rows.StructScan(&chassis)

		info := map[string]interface{}{}
		rows.MapScan(info)
		chassis.Colors = strings.Split(info["colors"].(string), ",")

		portsWg.Add(1)
		go func() {
			ports, _ := db.Queryx("SELECT * FROM ports WHERE chassis=$1", chassis.Uid)
			for ports.Next() {
				var port Port
				ports.StructScan(&port)
				chassis.Ports = append(chassis.Ports, port)
			}
			l.Chassis = append(l.Chassis, chassis)
			portsWg.Done()
		}()
	}
	portsWg.Wait()
	wg.Done()
}
