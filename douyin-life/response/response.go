package response

type ApiResponse struct {
	ErrNo  int         `json:"err_no"`
	ErrMsg string      `json:"err_msg"`
	LogId  string      `json:"log_id"`
	Data   interface{} `json:"data"`
}
