package pddRequest

import "encoding/json"

type PddDdkWeappQrcodeUrlGen struct {
	cashGiftId                uint64
	apiParas                  map[string]interface{}
	generateMallCollectCoupon bool
	customParameters          string
	goodsSignList             string
	pId                       string
	zsDuoId                   uint64
	typeName                  string
}

func (p *PddDdkWeappQrcodeUrlGen) SetCashGiftId(cashGiftId uint64) {
	p.cashGiftId = cashGiftId
	p.apiParas["cash_gift_id"] = cashGiftId
}
func (p *PddDdkWeappQrcodeUrlGen) SetGenerateMallCollectCoupon(generateMallCollectCoupon bool) {
	p.generateMallCollectCoupon = generateMallCollectCoupon
	p.apiParas["generate_mall_collect_coupon"] = generateMallCollectCoupon
}
func (p *PddDdkWeappQrcodeUrlGen) SetCustomParameters(customParameters string) {
	p.apiParas["custom_parameters"] = customParameters
	p.customParameters = customParameters
}
func (p *PddDdkWeappQrcodeUrlGen) SetGoodsSignList(goodsSignList []string) {
	marshal, _ := json.Marshal(goodsSignList)
	p.goodsSignList = string(marshal)

	p.apiParas["goods_sign_list"] = string(marshal)
}
func (p *PddDdkWeappQrcodeUrlGen) SetPId(pid string) {
	p.pId = pid
	p.apiParas["p_id"] = pid
}
func (p *PddDdkWeappQrcodeUrlGen) SetZsDuoId(zsDuoId uint64) {
	p.zsDuoId = zsDuoId
	p.apiParas["zs_duo_id"] = zsDuoId
}

func (p *PddDdkWeappQrcodeUrlGen) TypeName() string {
	return p.typeName
}

func (p *PddDdkWeappQrcodeUrlGen) SetTypeName(typeName string) {
	p.apiParas["type"] = typeName
	p.typeName = typeName
}

func NewPddDdkWeappQrcodeUrlGen() *PddDdkWeappQrcodeUrlGen {
	p := &PddDdkWeappQrcodeUrlGen{
		apiParas: make(map[string]interface{}),
	}
	p.SetTypeName(p.GetApiMethodName())
	return p
}

func (p *PddDdkWeappQrcodeUrlGen) GetApiParams() map[string]interface{} {
	return p.apiParas
}

func (p *PddDdkWeappQrcodeUrlGen) GetApiMethodName() string {
	return "pdd.ddk.weapp.qrcode.url.gen"
}
