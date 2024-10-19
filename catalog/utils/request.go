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
		sort = "laptops.launched DESC"
	case c.FeedSort_PriceAsc:
		sort = "laptops.price_min ASC"
	case c.FeedSort_PriceDesc:
		sort = "laptops.price_min DESC"
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
			if memory, err := strconv.ParseUint(filter.Val, 10, 64); err != nil {
				log.Println("Failed to parse memory", err)
			} else {
				queries = append(queries, fmt.Sprintf("memorys.size = %v", memory))
			}
		case "size":
			if size, err := strconv.ParseUint(filter.Val, 10, 64); err != nil {
				log.Println("Failed to parse size", err)
			} else {
				queries = append(queries, fmt.Sprintf("displays.size >= %v AND displays.size < %v", size, size+1))
			}
		case "storage":
			if storage, err := strconv.ParseUint(filter.Val, 10, 64); err != nil {
				log.Println("Failed to parse storage", err)
			} else {
				queries = append(queries, fmt.Sprintf("storages.capacity = %v", storage))
			}
		case "cpu":
			arr := []string{}
			for _, b := range strings.Split(filter.Val, ",") {
				if len(b) > 0 {
					arr = append(arr, fmt.Sprintf(`concat_ws(' ', cpus.architecture, cpus.model) ILIKE '%s%%'`, b))
				}
			}
			if len(arr) > 0 {
				cpus := strings.Join(arr, " OR ")
				queries = append(queries, fmt.Sprintf("(%s)", cpus))
			}
		case "gpu":
			arr := []string{}
			for _, b := range strings.Split(filter.Val, ",") {
				if len(b) > 0 {
					arr = append(arr, fmt.Sprintf(`'%s'`, b))
				}
			}
			if len(arr) > 0 {
				gpus := strings.Join(arr, ",")
				queries = append(queries, fmt.Sprintf("gpus.producer IN (%s)", gpus))
			}
		case "colors":
			arr := []string{}
			for _, b := range strings.Split(filter.Val, ",") {
				if len(b) > 0 {
					arr = append(arr, fmt.Sprintf(`chassis.colors ILIKE '%%%s%%'`, b))
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
