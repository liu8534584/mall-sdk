package response

type SuccessResponse struct {
	SnResponseContent struct {
		SnBody interface{} `json:"sn_body"`
	} `json:"sn_responseContent"`
}

type FailedResponse struct {
	SnResponseContent struct {
		SnError struct {
			ErrorCode string `json:"error_code"`
			ErrorMsg  string `json:"error_msg"`
		} `json:"sn_error"`
	} `json:"sn_responseContent"`
}
