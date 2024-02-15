package pddResponse

type PddErrorResponse struct {
	ErrorCode uint64 `json:"error_code"`
	SubCode   string `json:"sub_code"`
	ErrorMsg  string `json:"error_msg"`
	SubMsg    string `json:"sub_msg"`
}
