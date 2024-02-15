package buyin_liveShareMaterial_response

import (
	"doudian.com/open/sdk_golang/core"
)

type BuyinLiveShareMaterialResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *BuyinLiveShareMaterialData `json:"data"`
}
type DataItem struct {
	// 达人open_id
	AuthorOpenid string `json:"author_openid"`
	// 达人昵称
	AuthorName string `json:"author_name"`
	// 头像
	AuthorPic string `json:"author_pic"`
	// 达人等级
	AuthorLevel int64 `json:"author_level"`
	// 商品类别
	ProductCategory []string `json:"product_category"`
	// 作者省份
	Province string `json:"province"`
	// 作者城市
	City string `json:"city"`
	// 场均gmv
	AverageGmv string `json:"average_gmv"`
	// 粉丝数量
	FansNum int64 `json:"fans_num"`
	// 平均佣金率
	AverageCommissionRate string `json:"average_commission_rate"`
	// 直播间id
	RoomId string `json:"room_id"`
	// 正在直播
	IsLive bool `json:"is_live"`
	// 正在带货直播
	IsEcom bool `json:"is_ecom"`
	// 性别 male/ female/ 未知
	Gender string `json:"gender"`
	// 主播百应ID
	BuyinId string `json:"buyin_id"`
	// 该达人是否开通了直播预告权限
	LivePreviewAuth bool `json:"live_preview_auth"`
	// 该达人是否配置了生效中的抖客红包
	HasRedpack bool `json:"has_redpack"`
}
type BuyinLiveShareMaterialData struct {
	// 总数
	Total int64 `json:"total"`
	// 列表
	Data []DataItem `json:"data"`
}
