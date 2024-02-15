package request

type TbFullItemIdReq struct {
	ItemId        string        `json:"itemId"`
	TaobaoPidInfo TaobaoPidInfo `json:"taobaoPid"`
}

type ShortItemIds2LongReq struct {
	ItemIds string `json:"itemIds"`
}
