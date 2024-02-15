package kaolaRequest

type QueryGoodsInfoRequest struct {
	goodsIds      string
	goodsType     uint8
	trackingCode1 string
	trackingCode2 string
	goodsUrl      string
	apiParas      map[string]interface{}
}

func NewQueryGoodsInfoRequest() *QueryGoodsInfoRequest {
	p := &QueryGoodsInfoRequest{
		apiParas: make(map[string]interface{}),
	}
	return p
}

func (k *QueryGoodsInfoRequest) GoodsIds() string {
	return k.goodsIds
}

func (k *QueryGoodsInfoRequest) SetGoodsIds(goodsIds string) {
	k.goodsIds = goodsIds
	k.apiParas["goodsIds"] = goodsIds
}

func (k *QueryGoodsInfoRequest) GoodsUrl() string {
	return k.goodsUrl
}

func (k *QueryGoodsInfoRequest) SetGoodsUrl(goodsUrl string) {
	k.goodsUrl = goodsUrl
	k.apiParas["goodsUrl"] = goodsUrl
}

func (k *QueryGoodsInfoRequest) GoodsType() uint8 {
	return k.goodsType
}

func (k *QueryGoodsInfoRequest) SetGoodsType(goodsType uint8) {
	k.goodsType = goodsType
	k.apiParas["type"] = goodsType
}

func (k *QueryGoodsInfoRequest) TrackingCode1() string {
	return k.trackingCode1
}

func (k *QueryGoodsInfoRequest) SetTrackingCode1(trackingCode1 string) {
	k.trackingCode1 = trackingCode1
	k.apiParas["trackingCode1"] = trackingCode1
}

func (k *QueryGoodsInfoRequest) TrackingCode2() string {
	return k.trackingCode2
}

func (k *QueryGoodsInfoRequest) SetTrackingCode2(trackingCode2 string) {
	k.trackingCode2 = trackingCode2
	k.apiParas["trackingCode2"] = trackingCode2
}

func (k *QueryGoodsInfoRequest) GetApiMethodName() string {
	return "kaola.zhuanke.api.queryGoodsInfo"
}

func (k *QueryGoodsInfoRequest) GetApiParams() map[string]interface{} {

	return k.apiParas
}
