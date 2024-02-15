package request

type SearchProductRequest struct {
	Params *SearchProductParams
}

func (r *SearchProductRequest) GetParams() interface{} {
	return r.Params
}

func (r *SearchProductRequest) GetUrlPath() string {
	return "/api/life/v1/outside_distribution/search_product/"
}

type SearchProductParams struct {
	UID         int64   `json:"uid,omitempty"`    //在抖⾳创建的达⼈账号的uid
	Cursor      int     `json:"cursor,omitempty"` //查询光标起点，默认0
	Count       int     `json:"count,omitempty"`
	Pid         string  `json:"pid,omitempty"`
	CityCode    string  `json:"city_code,omitempty"`
	Longitude   float64 `json:"longitude,omitempty"`
	Latitude    float64 `json:"latitude,omitempty"`
	KeyWord     string  `json:"key_word,omitempty"`
	SortBy      int     `json:"sort_by,omitempty"`  //1-距离(根据city_code城市下适⽤⻔店的距 离排序), 2-销量 3-优惠前价格 (按优惠前价格排序，因此优 惠后价格更⾼的商品可能排在前⾯) 4-佣⾦率
	OrderBy     int     `json:"order_by,omitempty"` //
	CategoryID  string  `json:"category_id,omitempty"`
	NeedInvoice bool    `json:"need_invoice,omitempty"` //筛选项：是否需要发票（商家给出分佣佣 ⾦后需要发票）
	DistanceMax int     `json:"distance_max,omitempty"` //最⼤距离，单位m
}
