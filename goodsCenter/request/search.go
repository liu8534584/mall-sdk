package request

type SearchReq struct {
	Content       string        `json:"content"`
	ItemType      int           `json:"itemType"`
	IsCoupon      bool          `json:"isCoupon"`
	Page          int           `json:"page"`
	PageSize      int           `json:"pageSize"`
	OptId         string        `json:"optId"`
	ActivityId    uint64        `json:"activityId"`
	SortType      int           `json:"sortType"`
	TopFilter     int           `json:"topFilter"`
	AppId         string        `json:"appId,omitempty"`
	TaobaoPidInfo TaobaoPidInfo `json:"taobaoPid"`
	PddPidInfo    PddPidInfo    `json:"pddPid"`
	JdChannelId   int           `json:"jdChannelId"`
	CustomParams  string        `json:"customParams"`
	CheckStatus   bool          `json:"checkStatus"`
	GoodsImgType  uint8         `json:"goodsImgType"`
	ChannelType   string        `json:"channelType"`
	IsTmall       bool          `json:"isTmall"`
	IsOwner       bool          `json:"isOwner"`
	Geo           SearchGeoInfo `json:"geo"`
	Cat           string        `json:"cat"`
	Cat0Id        int64         `json:"cat0Id"`
	Cat1Id        int64         `json:"cat1Id"`
	Cursor        int           `json:"cursor"`
}

type SearchGeoInfo struct {
	CityId     int64   `json:"cityId"`
	RegionId   int64   `json:"regionId"`
	DistrictId int64   `json:"districtId"`
	Lng        float64 `json:"lng"`
	Lat        float64 `json:"lat"`
	Radii      int64   `json:"radii"`
}

type PddPidInfo struct {
	PddPid          string `json:"pddPid"`
	Uid             string `json:"uid"`
	Sid             string `json:"sid"`
	PddCustomParams string `json:"pddCustomParams"`
}

type TaobaoPidInfo struct {
	Pid string `json:"pid"`
	Rid string `json:"rid"`
	Sid string `json:"sid"`
}
