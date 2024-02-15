package buyin_kolPidList_response

import (
	"doudian.com/open/sdk_golang/core"
)

type BuyinKolPidListResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *BuyinKolPidListData `json:"data"`
}
type PidsItem struct {
	// PID
	Pid string `json:"pid"`
	// 推广媒体类型
	MediaType int32 `json:"media_type"`
	// 渠道ID
	MediaId int64 `json:"media_id"`
	// 渠道名称
	MediaName string `json:"media_name"`
	// 推广位名称
	SiteName string `json:"site_name"`
	// 创建时间
	CreateTime int64 `json:"create_time"`
	// 更新时间
	UpdateTime int64 `json:"update_time"`
}
type Data struct {
	// 总数
	Total int32 `json:"total"`
	// PID列表
	Pids []PidsItem `json:"pids"`
}
type BuyinKolPidListData struct {
	// 结果
	Data Data `json:"data"`
}
