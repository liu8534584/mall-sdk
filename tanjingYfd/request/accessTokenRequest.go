package request

import (
	"encoding/json"
	tanjing "github.com/liu8534584/mall-sdk/tanjingYfd"
	"github.com/liu8534584/mall-sdk/tanjingYfd/tanjingResponse"
)

type AccessTokenRequest struct {
	tanjing.TanjingBaseRequest
	params *AccessTokenParams
}

func NewAccessToken(config *tanjing.TanjingBaseConfig) *AccessTokenRequest {
	request := AccessTokenRequest{}
	request.SetConfig(config)
	request.SetClient(tanjing.DefaultTanjingClient)
	request.params = &AccessTokenParams{}

	return &request
}

func (a *AccessTokenRequest) Execute() (*tanjingResponse.AccessTokenResp, error) {
	execute, err := a.GetClient().Execute(a, "")
	if err != nil {
		return nil, err
	}

	var resp *tanjingResponse.AccessTokenResp
	_ = json.Unmarshal([]byte(execute), &resp)
	return resp, nil
}

func (a *AccessTokenRequest) GetApiPath() string {
	return "/api/oauth/get_token"
}
func (a *AccessTokenRequest) GetApiParamsObject() interface{} {
	return a.params
}

func (a *AccessTokenRequest) GetApiParams() *AccessTokenParams {
	return a.params
}

type AccessTokenParams struct {
	Merchant string `json:"merchant"`
	Secret   string `json:"secret"`
}
