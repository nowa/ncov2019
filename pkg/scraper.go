package scraper

var (
	DataURL = "https://view.inews.qq.com/g2/getOnsInfo?name=wuwei_ww_area_counts"
)

func GetAllData() (*[]model.City, error) {
	res, err := http.Get(DataURL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var (
		ret    int
		data   string
		cities []model.City
	)
	_ = json.Unmarshal(body, &struct {
		Ret  int     `json:"ret"`
		Data *string `json:"data"`
	}{ret, &data})

	_ = json.Unmarshal([]byte(data), &cities)
	return &cities, nil
}
