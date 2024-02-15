package response

type ProductDetailResponse struct {
	ApiResponse
	Data ProductDetailData `json:"data"`
}

type ProductDetailData struct {
	ProductMap map[string]ProductDetail `json:"product_map"`
}
