package request

type ProductDetailRequest struct {
	Params *ProductDetailParams
}

func (r *ProductDetailRequest) GetParams() interface{} {
	return r.Params
}

func (r *ProductDetailRequest) GetUrlPath() string {
	return "/api/life/v1/outside_distribution/mget_product_by_id/"
}

type ProductDetailParams struct {
	UID           int64   `json:"uid,omitempty"`
	CityCode      string  `json:"city_code,omitempty"`
	Longitude     float64 `json:"longitude,omitempty"`
	Latitude      float64 `json:"latitude,omitempty"`
	ProductIDList []int64 `json:"product_id_list,omitempty"`
	SortBy        int     `json:"sort_by,omitempty"`
	OrderBy       int     `json:"order_by,omitempty"`
}
