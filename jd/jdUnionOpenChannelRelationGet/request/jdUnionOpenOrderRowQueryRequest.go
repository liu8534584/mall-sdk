package request

import (
	"encoding/json"
	"errors"
	"github.com/liu8534584/mall-sdk/jd"
	"github.com/liu8534584/mall-sdk/jd/jdUnionOpenChannelRelationGet/response"
)

type JdUnionOpenChannelRelationGetRequest struct {
	jd.BaseJdApiRequest
	Param *JdUnionOpenChannelRelationGetParam
}

func (j *JdUnionOpenChannelRelationGetRequest) GetMethodName() string {
	return "jd.union.open.channel.relation.get"
}

func New(config *jd.JdBaseConfig) *JdUnionOpenChannelRelationGetRequest {
	request := JdUnionOpenChannelRelationGetRequest{
		Param: &JdUnionOpenChannelRelationGetParam{},
	}
	request.SetConfig(config)
	request.SetClient(jd.DefaultJdApiClient)

	return &request
}

func (j *JdUnionOpenChannelRelationGetRequest) Execute() (result *response.JdUnionOpenChannelRelationGetResponseResult, err error) {
	responseJson, err := j.GetClient().Request(j, true)
	if err != nil {
		return
	}

	resp := &response.JdUnionOpenChannelRelationGetResponse{}
	if err = json.Unmarshal([]byte(responseJson), resp); err != nil {
		return
	}
	if resp.JdUnionOpenChannelRelationGetResponce.GetResult != "" {
		if err = json.Unmarshal([]byte(resp.JdUnionOpenChannelRelationGetResponce.GetResult), &result); err != nil {
			return
		}
		return
	} else {
		err = errors.New("result is null")
	}
	return

}

func (j *JdUnionOpenChannelRelationGetRequest) GetParamsObject() interface{} {
	return j.Param
}

func (j *JdUnionOpenChannelRelationGetRequest) GetParams() *JdUnionOpenChannelRelationGetParam {
	return j.Param
}

type JdUnionOpenChannelRelationGetParam struct {
	ChannelRelationGetReq ChannelRelationGetReq `json:"channelRelationGetReq"`
}

type ChannelRelationGetReq struct {
	InviteCode string `json:"inviteCode"`
	Note       string `json:"note"` //仅支持传入中文、字母、数字、下划线或中划线，最多20个字符
}
