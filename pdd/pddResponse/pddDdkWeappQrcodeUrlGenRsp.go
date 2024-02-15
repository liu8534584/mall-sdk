package pddResponse

type PddDdkWeappQrcodeUrlGenRsp struct {
	WeappQrcodeGenerateResponse WeappQrcodeGenerateResponse `json:"weapp_qrcode_generate_response"`
}

type WeappQrcodeGenerateResponse struct {
	Url string `json:"url"`
}
