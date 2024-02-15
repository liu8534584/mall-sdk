package response

type SubCode struct {
	SubCode string `json:"subCode"`
}

type QuerySearchcommodityResponse struct {
	SnResponseContent struct {
		SnBody struct {
			QuerySearchcommodity []Detail `json:"querySearchcommodity"`
		} `json:"sn_body"`
	} `json:"sn_responseContent"`
}
