package pddResponse

type PddDdkGoodsPidGenerateRsp struct {
	PIdGenerateResponse PddDdkGoodsPidGenerateInfoRsp `json:"p_id_generate_response"`
	ErrorResponse       PddErrorResponse              `json:"error_response"`
}

type PddDdkGoodsPidGenerateInfoRsp struct {
	PIdList        []PddDdkGoodsPidGeneratePidInfoRsp `json:"p_id_list"`
	RemainPidCount uint64                             `json:"remain_pid_count"`
}

type PddDdkGoodsPidGeneratePidInfoRsp struct {
	CreateTime uint64 `json:"create_time"`
	PidName    string `json:"pid_name"`
	PId        string `json:"p_id"`
	MediaId    uint64 `json:"media_id"`
}
