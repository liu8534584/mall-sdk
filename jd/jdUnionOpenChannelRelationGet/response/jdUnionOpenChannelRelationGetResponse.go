package response

type JdUnionOpenChannelRelationGetResponse struct {
	JdUnionOpenChannelRelationGetResponce struct {
		Code      string `json:"code"`
		GetResult string `json:"getResult"`
	} `json:"jd_union_open_channel_relation_get_responce"`
}

type JdUnionOpenChannelRelationGetResponseResult struct {
	Code int `json:"code"`
	Data struct {
		ChannelId int `json:"channelId"`
	} `json:"data"`
	Message   string `json:"message"`
	RequestId string `json:"requestId"`
}
