package utils

import (
	"strings"

	c "github.com/kibzrael/raelelectronics/common/api/catalog"

	"github.com/kibzrael/raelelectronics/catalog/data"
)

func ParseLaptopCard(l data.LaptopCard) c.LaptopCard {
	return c.LaptopCard{
		Uid:         l.Uid,
		Name:        l.Name,
		Brand:       l.Brand,
		Thumbnail:   l.Thumbnail,
		Launched:    l.Launched,
		PriceMin:    l.PriceMin,
		PriceMax:    l.PriceMax,
		Colors:      strings.Split(l.Colors, ","),
		Size:        l.Size,
		Cpu:         l.Cpu,
		Cores:       l.Cores,
		BaseSpeed:   l.BaseSpeed,
		Memory:      l.Memory,
		Storage:     l.Storage,
		BatteryLife: l.BatteryLife,
	}
}

func ParseLaptop(l data.Laptop) (laptop c.Laptop) {
	laptop = c.Laptop{
		Uid:       l.Uid,
		NotebId:   l.NotebId,
		Name:      l.Name,
		Brand:     l.Brand,
		Thumbnail: l.Thumbnail,
		Images:    l.Images,
		Link:      l.Link,
		Launched:  l.Launched,
		PriceMin:  l.PriceMin,
		PriceMax:  l.PriceMax,
	}

	for _, cpu := range l.CPU {
		laptop.CPU = append(laptop.CPU, &c.CPU{
			Uid:             cpu.Uid,
			Model:           cpu.Model,
			Architecture:    cpu.Architecture,
			Lithography:     int64(cpu.Lithography),
			Cache:           cpu.Cache,
			BaseSpeed:       cpu.BaseSpeed,
			BoostSpeed:      cpu.BoostSpeed,
			Cores:           cpu.Cores,
			Tdp:             cpu.Tdp,
			IntegratedVideo: cpu.IntegratedVideo,
			OtherInfo:       cpu.OtherInfo,
		})
	}

	for _, gpu := range l.GPU {
		laptop.GPU = append(laptop.GPU, &c.GPU{
			Uid:          gpu.Uid,
			Producer:     gpu.Producer,
			Model:        gpu.Model,
			Architecture: gpu.Architecture,
			Lithography:  int64(gpu.Lithography),
			BaseSpeed:    gpu.BaseSpeed,
			BoostSpeed:   gpu.BoostSpeed,
			MemorySize:   gpu.MemorySize,
			MemorySpeed:  gpu.MemorySpeed,
			MemoryType:   gpu.MemoryType,
			Tdp:          gpu.Tdp,
			OtherInfo:    gpu.OtherInfo,
		})
	}

	for _, display := range l.Display {
		laptop.Display = append(laptop.Display, &c.Display{
			Uid:         display.Uid,
			Size:        display.Size,
			HResolution: display.HResolution,
			VResolution: display.VResolution,
			Type:        display.Type,
			SRGB:        display.SRGB,
			Touchscreen: display.Touchscreen,
			OtherInfo:   display.OtherInfo,
		})
	}

	for _, memory := range l.Memory {
		laptop.Memory = append(laptop.Memory, &c.Memory{
			Uid:   memory.Uid,
			Size:  memory.Size,
			Speed: memory.Speed,
			Type:  memory.Type,
		})
	}

	for _, storage := range l.Storage {
		laptop.Storage = append(laptop.Storage, &c.Storage{
			Uid:      storage.Uid,
			Capacity: storage.Capacity,
			Speed:    storage.Speed,
			Model:    storage.Model,
		})
	}

	for _, card := range l.WirelessCard {
		laptop.WirelessCard = append(laptop.WirelessCard, &c.WirelessCard{
			Uid:       card.Uid,
			Model:     card.Model,
			Speed:     card.Speed,
			OtherInfo: card.OtherInfo,
		})
	}

	for _, board := range l.Motherboard {
		laptop.Motherboard = append(laptop.Motherboard, &c.Motherboard{
			Uid:     board.Uid,
			Ram:     board.Ram,
			Storage: board.Storage,
		})
	}

	for _, chassis := range l.Chassis {
		ch := &c.Chasis{
			Uid:       chassis.Uid,
			Height:    chassis.Height,
			Depth:     chassis.Depth,
			Width:     chassis.Width,
			Weight:    chassis.Weight,
			Colors:    chassis.Colors,
			Webcam:    chassis.Webcam,
			OtherInfo: chassis.OtherInfo,
		}

		for _, port := range chassis.Ports {
			ch.Ports = append(ch.Ports, &c.Chasis_Port{
				Uid:   port.Uid,
				Count: int64(port.Count),
				Type:  port.Type,
			})
		}
		laptop.Chasis = append(laptop.Chasis, ch)
	}

	for _, battery := range l.Battery {
		laptop.Battery = append(laptop.Battery, &c.Battery{
			Uid:       battery.Uid,
			Capacity:  battery.Capacity,
			Type:      battery.Type,
			Life:      battery.Life,
			OtherInfo: battery.OtherInfo,
		})
	}
	return
}
