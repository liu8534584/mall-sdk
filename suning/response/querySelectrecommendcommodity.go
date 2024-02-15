package response

type QuerySelectrecommendcommodityResponse struct {
	SnResponseContent struct {
		SnBody struct {
			QuerySelectrecommendcommodity struct {
				CommodityList []Detail `json:"commodityList"`
				IsHaveData    string   `json:"isHaveData"`
			} `json:"querySelectrecommendcommodity"`
		} `json:"sn_body"`
	} `json:"sn_responseContent"`
}
