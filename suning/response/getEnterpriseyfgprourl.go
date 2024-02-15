package response

type GetEnterPriseyfgproUrlResponse struct {
	SnResponseContent struct {
		SnBody struct {
			GetEnterpriseyfgprourl struct {
				ActName    string `json:"actName"`
				ActType    string `json:"actType"`
				CommandURL string `json:"commandUrl"`
				Desc       string `json:"desc"`
				HtmlURL    string `json:"htmlUrl"`
				SpQRCode   string `json:"spQRCode"`
				SpURL      string `json:"spUrl"`
			} `json:"getEnterpriseyfgprourl"`
		} `json:"sn_body"`
	} `json:"sn_responseContent"`
}
