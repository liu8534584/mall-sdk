package doudian_sdk

type BaseDoudianOpApiResponse struct {
	ErrNo int64 `json:"err_no"`

	Message string `json:"message"`

	LogId string `json:"log_id"`

	Code int64 `json:"code"`

	Msg string `json:"msg"`

	SubCode string `json:"sub_code"`

	SubMsg string `json:"sub_msg"`
}
