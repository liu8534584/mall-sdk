package token_refresh_response

import (
	"doudian.com/open/sdk_golang/core"
)

type TokenRefreshResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *TokenRefreshData `json:"data"`
}
type TokenRefreshData struct {
	// token值
	AccessToken string `json:"access_token"`
	// 过期时间(秒级时间戳)
	ExpiresIn int64 `json:"expires_in"`
	// 刷新token值
	RefreshToken string `json:"refresh_token"`
	// 范围
	Scope string `json:"scope"`
	// 店铺ID
	ShopId int64 `json:"shop_id"`
	// 店铺名称
	ShopName string `json:"shop_name"`
	// 授权ID
	AuthorityId string `json:"authority_id"`
	//过期时间点(秒级时间戳)
	ExpiresTime int64 `json:"expires_time"`
}
