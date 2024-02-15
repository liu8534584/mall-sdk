package request

import (
	"encoding/json"
	tanjing "github.com/liu8534584/mall-sdk/tanjingYfd"
	"github.com/liu8534584/mall-sdk/tanjingYfd/tanjingResponse"
)

// GetAllTagsRequest 获取标签列表/**
type GetAllTagsRequest struct {
	tanjing.TanjingBaseRequest
	params *GetAllTagsRequestParams
}

func NewGetAllTagsRequest(config *tanjing.TanjingBaseConfig) *GetAllTagsRequest {
	sign := GetAllTagsRequest{
		params: &GetAllTagsRequestParams{},
	}
	sign.SetConfig(config)
	sign.SetClient(tanjing.DefaultTanjingClient)

	return &sign
}

func (r *GetAllTagsRequest) Execute(token string) (*tanjingResponse.GetAllTagsResp, error) {
	execute, err := r.GetClient().Execute(r, token)
	if err != nil {
		return nil, err
	}

	var resp *tanjingResponse.GetAllTagsResp
	_ = json.Unmarshal([]byte(execute), &resp)
	return resp, nil
}

func (r *GetAllTagsRequest) GetApiParamsObject() interface{} {
	return r.params
}

func (r *GetAllTagsRequest) GetApiParams() *GetAllTagsRequestParams {
	return r.params
}

func (r *GetAllTagsRequest) GetApiPath() string {
	return "/scrm/Friend/GetAllTags"
}

type GetAllTagsRequestParams struct {
	VcMerchantNo    string `json:"vcMerchantNo"`    //商家编号
	VcRobotSerialNo string `json:"vcRobotSerialNo"` // 机器人编号
}
