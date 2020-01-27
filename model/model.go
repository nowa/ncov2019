package model

import (
	"log"
)

type Country struct {
	Name string `json:"country"`
}

type Province struct {
	Name string `json:"area"`
}

type City struct {
	Name string `json:"city"`
	Province
	Country
	Municipality bool `json:"municipality"`
	Confirmed    int  `json:"confirm"`
	Suspected    int  `json:"suspect"`
	Cured        int  `json:"heal"`
	Dead         int  `json:"dead"`
}

func ParseData(cities []*City) map[string]int {
	m, c, s, h, d := make(map[string]int), make(map[string]int), make(map[string]int), make(map[string]int), make(map[string]int)

	for _, city := range cities {
		if city.Name == "" && city.Country.Name == "中国" {
			log.Println("Municipality: ", city)
			city.Name = city.Province.Name
			city.Municipality = true
		}

		country := city.Country
		c[country.Name] += city.Confirmed
		s[country.Name] += city.Suspected
		h[country.Name] += city.Cured
		d[country.Name] += city.Dead
	}

	log.Println("c: ", c)

	return m
}
