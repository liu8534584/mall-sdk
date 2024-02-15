package meituanDdResponse

type DdRegionsResponse struct {
	Code int64          `json:"code"`
	Msg  MtCityRegionVO `json:"msg"`
}

type MtCityRegionVO struct {
	CityId                 int64                        `json:"cityId"`
	CityName               string                       `json:"cityName"`
	HotRegionInfos         map[string]string            `json:"hotRegionInfos"`
	MainRegionInfos        map[string]string            `json:"mainRegionInfos"`
	SubRegionInfos         map[string]map[string]string `json:"subRegionInfos"`
	SubCityInfos           map[string]string            `json:"subCityInfos"`
	SubCityMainRegionInfos map[string]map[string]string `json:"subCityMainRegionInfos"`
	DistrictInfos          map[string]string            `json:"districtInfos"`
	DistrictToRegionMaps   map[string]map[string]string `json:"districtToRegionMaps"`
}
