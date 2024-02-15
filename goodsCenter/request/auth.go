package request

type AuthStatusReq struct {
	GoodsSource int        `json:"goodsSource"`
	PddPid      PddPidInfo `json:"pddPid"`
}
