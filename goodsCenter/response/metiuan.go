package response

import "github.com/liu8534584/mall-sdk/goodsCenter/response/base"

type MeituanCategoryResp struct {
	Code     base.Code          `json:"code"`
	Message  string             `json:"msg"`
	Demotion base.Demotion      `json:"demotion"`
	Data     MeituanCategoryRes `json:"data"`
}

type MeituanCategoryRes struct {
	Category []MeituanCategory `json:"category"`
	Area     []MeituanArea     `json:"area"`
	SortType []MeituanSortType `json:"sortType"`
}

type MeituanCategory struct {
	Cat0Id  int    `json:"cat0Id"`
	CatName string `json:"catName"`
	Cat1    []struct {
		Cat1Id  int    `json:"cat1Id"`
		CatName string `json:"catName"`
	} `json:"cat1"`
}
type MeituanArea struct {
	DistrictId   int             `json:"districtId"`
	DistrictName string          `json:"districtName"`
	Region       []MeituanRegion `json:"region"`
}

type MeituanRegion struct {
	RegionId   int    `json:"regionId"`
	RegionName string `json:"regionName"`
}
type MeituanSortType struct {
	Type int    `json:"type"`
	Name string `json:"name"`
}
