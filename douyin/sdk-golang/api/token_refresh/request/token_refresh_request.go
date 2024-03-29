package token_refresh_request

import (
	"doudian.com/open/sdk_golang/api/token_refresh/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type TokenRefreshRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *TokenRefreshParam
}

func (c *TokenRefreshRequest) GetUrlPath() string {
	return "/token/refresh"
}

func New() *TokenRefreshRequest {
	request := &TokenRefreshRequest{
		Param: &TokenRefreshParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}

func (c *TokenRefreshRequest) Execute(accessToken *doudian_sdk.AccessToken) (*token_refresh_response.TokenRefreshResponse, error) {
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &token_refresh_response.TokenRefreshResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}

func (c *TokenRefreshRequest) GetParamObject() interface{} {
	return c.Param
}

func (c *TokenRefreshRequest) GetParams() *TokenRefreshParam {
	return c.Param
}

type TokenRefreshParam struct {
	// 刷新token的值
	RefreshToken string `json:"refresh_token"`
	// 授权类型
	GrantType string `json:"grant_type"`
}
