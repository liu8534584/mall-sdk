package pddResponse

type PddDdkTmcActivityListGenerateRsp struct {
	TmcAtyListResponse TmcAtyListResponse `json:"tmc_aty_list_response"`
	ErrorResponse      PddErrorResponse   `json:"error_response"`
}

type TmcAtyVoList struct {
	EndTime   string `json:"end_time"`
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	StartTime string `json:"start_time"`
	Type      string `json:"type"`
}

type TmcAtyListResponse struct {
	QueryEndTime   string         `json:"query_end_time"`
	QueryStartTime string         `json:"query_start_time"`
	TmcAtyVoList   []TmcAtyVoList `json:"tmc_aty_vo_list"`
	Total          int64          `json:"total"`
}
