package model

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
