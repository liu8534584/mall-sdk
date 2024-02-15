package response

type CommandParseAndShareResponse struct {
	ApiResponse
	Data CommandParseAndShareData `json:"data"`
}

type CommandParseAndShareData struct {
	CommandType int         `json:"command_type"`
	ProductInfo ProductInfo `json:"product_info"`
}

type QrCode struct {
	URL string `json:"url"`
}
type ShareInfo struct {
	DyPassword string `json:"dy_password"`
	ShareLink  string `json:"share_link"`
	DyDeeplink string `json:"dy_deeplink"`
	DyZlink    string `json:"dy_zlink"`
	QrCode     QrCode `json:"qr_code"`
}
type ProductInfo struct {
	ProductID int64     `json:"product_id"`
	ShareInfo ShareInfo `json:"share_info"`
}
