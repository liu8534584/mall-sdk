package pddRequest

type PddDdkGoodsDetail struct {
	typeName         string
	apiParas         map[string]interface{}
	customParameters string
	goodsImgType     uint8
	goodsSign        string
	pid              string
	searchId         string
	zsDuoId          uint64
}

func (p *PddDdkGoodsDetail) ZsDuoId() uint64 {
	return p.zsDuoId
}

func (p *PddDdkGoodsDetail) SetZsDuoId(zsDuoId uint64) {
	p.apiParas["zs_duo_id"] = zsDuoId
	p.zsDuoId = zsDuoId
}

func (p *PddDdkGoodsDetail) SearchId() string {
	return p.searchId
}

func (p *PddDdkGoodsDetail) SetSearchId(searchId string) {
	p.apiParas["search_id"] = searchId
	p.searchId = searchId
}

func (p *PddDdkGoodsDetail) Pid() string {
	return p.pid
}

func (p *PddDdkGoodsDetail) SetPid(pid string) {
	p.apiParas["pid"] = pid
	p.pid = pid
}

func (p *PddDdkGoodsDetail) GoodsSign() string {
	return p.goodsSign
}

func (p *PddDdkGoodsDetail) SetGoodsSign(goodsSign string) {
	p.apiParas["goods_sign"] = goodsSign
	p.goodsSign = goodsSign
}

func (p *PddDdkGoodsDetail) GoodsImgType() uint8 {
	return p.goodsImgType
}

func (p *PddDdkGoodsDetail) SetGoodsImgType(goodsImgType uint8) {
	p.apiParas["goods_img_type"] = goodsImgType
	p.goodsImgType = goodsImgType
}

func (p *PddDdkGoodsDetail) CustomParameters() string {
	return p.customParameters
}

func (p *PddDdkGoodsDetail) SetCustomParameters(customParameters string) {
	p.apiParas["custom_parameters"] = customParameters
	p.customParameters = customParameters
}

func (p *PddDdkGoodsDetail) TypeName() string {
	return p.typeName
}

func (p *PddDdkGoodsDetail) SetTypeName(typeName string) {
	p.apiParas["type"] = typeName
	p.typeName = typeName
}

func NewPddDdkGoodsDetail() *PddDdkGoodsDetail {
	p := &PddDdkGoodsDetail{
		apiParas: make(map[string]interface{}),
	}
	//fmt.Println(p.GetApiMethodName())
	p.SetTypeName(p.GetApiMethodName())
	return p
}

func (p *PddDdkGoodsDetail) GetApiParams() map[string]interface{} {
	return p.apiParas
}

func (p *PddDdkGoodsDetail) GetApiMethodName() string {
	return "pdd.ddk.goods.detail"
}
