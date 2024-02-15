package tanjingResponse

type PushDownFileMessageResp struct {
	NResult    int    `json:"nResult"`
	VcResult   string `json:"vcResult"`
	VcSerialNo string `json:"vcSerialNo"`
	Status     int    `json:"status"`
}
