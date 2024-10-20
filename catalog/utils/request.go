package utils

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	c "github.com/kibzrael/raelelectronics/common/api/catalog"
)

func SortQuery(req *c.FeedRequest) (sort string) {
	switch req.Sort {
	case c.FeedSort_New:
		sort = "l.launched DESC"
	case c.FeedSort_PriceAsc:
		sort = "l.price_min ASC"
	case c.FeedSort_PriceDesc:
		sort = "l.price_min DESC"
	default:
		sort = "l.featured DESC"
	}
	return
}

func FilterQuery(req *c.FeedRequest) (query string) {
	queries := []string{}
	for _, filter := range req.Filters {
		switch filter.Key {
		case "brand":
			arr := []string{}
			for _, b := range strings.Split(filter.Val, ",") {
				if len(b) > 0 {
					arr = append(arr, fmt.Sprintf(`'%s'`, b))
				}
			}
			if len(arr) > 0 {
				brands := strings.Join(arr, ",")
				queries = append(queries, fmt.Sprintf("laptops.brand IN (%s)", brands))
			}
		case "price_min":
			if price, err := strconv.ParseFloat(filter.Val, 64); err != nil {
				log.Println("Failed to parse price_min", err)
			} else {
				queries = append(queries, fmt.Sprintf("laptops.price_min >= %v", price))
			}
		case "price_max":
			if price, err := strconv.ParseFloat(filter.Val, 64); err != nil {
				log.Println("Failed to parse price_max", err)
			} else {
				queries = append(queries, fmt.Sprintf("laptops.price_max <= %v", price))
			}
		case "memory":
			arr := []string{}
			for _, m := range strings.Split(filter.Val, ",") {
				if memory, err := strconv.ParseUint(m, 10, 64); err != nil {
					log.Println("Failed to parse memory", err)
				} else {
					arr = append(arr, fmt.Sprintf("memorys.size = %v", memory))
				}
			}
			if len(arr) > 0 {
				memorys := strings.Join(arr, " OR ")
				queries = append(queries, fmt.Sprintf("(%s)", memorys))
			}
		case "size":
			arr := []string{}
			for _, s := range strings.Split(filter.Val, ",") {
				if size, err := strconv.ParseUint(s, 10, 64); err != nil {
					log.Println("Failed to parse size", err)
				} else {
					arr = append(arr, fmt.Sprintf("displays.size >= %v AND displays.size < %v", size, size+1))
				}
			}
			if len(arr) > 0 {
				sizes := strings.Join(arr, " OR ")
				queries = append(queries, fmt.Sprintf("(%s)", sizes))
			}
		case "display":
			arr := []string{}
			for _, d := range strings.Split(filter.Val, ",") {
				if len(d) > 0 {
					arr = append(arr, fmt.Sprintf(`'%s'`, d))
				}
			}
			if len(arr) > 0 {
				displays := strings.Join(arr, ",")
				queries = append(queries, fmt.Sprintf("displays.type IN (%s)", displays))
			}
		case "storage":
			arr := []string{}
			for _, s := range strings.Split(filter.Val, ",") {
				if storage, err := strconv.ParseUint(s, 10, 64); err != nil {
					log.Println("Failed to parse storage", err)
				} else {
					arr = append(arr, fmt.Sprintf("storages.capacity = %v", storage))
				}
			}
			if len(arr) > 0 {
				storages := strings.Join(arr, " OR ")
				queries = append(queries, fmt.Sprintf("(%s)", storages))
			}
		case "cpu":
			arr := []string{}
			for _, c := range strings.Split(filter.Val, ",") {
				if len(c) > 0 {
					arr = append(arr, fmt.Sprintf(`concat_ws(' ', cpus.architecture, cpus.model) ILIKE '%s%%'`, c))
				}
			}
			if len(arr) > 0 {
				cpus := strings.Join(arr, " OR ")
				queries = append(queries, fmt.Sprintf("(%s)", cpus))
			}
		case "gpu":
			arr := []string{}
			for _, g := range strings.Split(filter.Val, ",") {
				if len(g) > 0 {
					arr = append(arr, fmt.Sprintf(`'%s'`, g))
				}
			}
			if len(arr) > 0 {
				gpus := strings.Join(arr, ",")
				queries = append(queries, fmt.Sprintf("gpus.producer IN (%s)", gpus))
			}
		case "colors":
			arr := []string{}
			for _, c := range strings.Split(filter.Val, ",") {
				if len(c) > 0 {
					arr = append(arr, fmt.Sprintf(`chassis.colors ILIKE '%%%s%%'`, c))
				}
			}
			if len(arr) > 0 {
				colors := strings.Join(arr, " OR ")
				queries = append(queries, fmt.Sprintf("(%s)", colors))
			}
		}
	}

	if len(queries) > 0 {
		query = fmt.Sprintf(`
		WHERE
			%s
		`, strings.Join(queries, " AND "))
	}
	return
}
