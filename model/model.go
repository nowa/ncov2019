package model

import ()

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

func ParseCountryData(cities []*City) map[string]map[string]int {
	c := make(map[string]map[string]int)

	for _, city := range cities {
		if city.Name == "" && city.Country.Name == "中国" {
			// log.Println("Municipality: ", city)
			city.Name = city.Province.Name
			city.Municipality = true
		}

		country := city.Country

		if c[country.Name] == nil {
			c[country.Name] = make(map[string]int)
		}
		c[country.Name]["C"] += city.Confirmed
		c[country.Name]["S"] += city.Suspected
		c[country.Name]["H"] += city.Cured
		c[country.Name]["D"] += city.Dead
	}

	// log.Println("c: ", c)

	return c
}
