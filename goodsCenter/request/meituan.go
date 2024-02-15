package request

type MeituanCategoryReq struct {
	CityId int64   `json:"cityId"`
	Lng    float64 `json:"lng"`
	Lat    float64 `json:"lat"`
}
