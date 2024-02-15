package request

type GetSecretCodeRequest struct {
	PlatformId string `json:"platformId"`
	Phone      string `json:"phone"`
}

func NewGetSecretCodeRequest() *GetSecretCodeRequest {
	return &GetSecretCodeRequest{}
}

func (r *GetSecretCodeRequest) GetParams() map[string]interface{} {
	m := make(map[string]interface{})
	m["platformId"] = r.PlatformId
	m["phone"] = r.Phone
	return m
}

func (r *GetSecretCodeRequest) GetUrlPath() string {
	return "/services/v3/begin/getSecretCode"
}
