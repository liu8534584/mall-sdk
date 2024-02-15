package kaolaResponse

type QueryRecommendGoodsListResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data []GoodsInfo `json:"data"`
}
