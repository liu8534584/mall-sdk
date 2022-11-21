package pddResponse

type PddDdkMemberAuthorityQueryRsp struct {
	AuthorityQueryResponse PddDdkMemberAuthorityQueryInfoRsp `json:"authority_query_response"`
	ErrorResponse          PddErrorResponse                  `json:"error_response"`
}

type PddDdkMemberAuthorityQueryInfoRsp struct {
	Bind uint8 `json:"bind"`
}
