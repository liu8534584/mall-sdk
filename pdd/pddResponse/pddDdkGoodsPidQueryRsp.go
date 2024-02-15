package pddResponse

type PddDdkGoodsPidQueryRsp struct {
	PIdQueryResponse PddDdkGoodsPidQueryInfoRsp `json:"p_id_query_response"`
	ErrorResponse    PddErrorResponse           `json:"error_response"`
}

type PddDdkGoodsPidQueryInfoRsp struct {
	PIdList    []PddDdkGoodsPidQueryPidInfoRsp `json:"p_id_list"`
	TotalCount uint64                          `json:"total_count"`
}

type PddDdkGoodsPidQueryPidInfoRsp struct {
	CreateTime uint64 `json:"create_time"`
	PidName    string `json:"pid_name"`
	PId        string `json:"p_id"`
	Status     uint16 `json:"status"`
}
