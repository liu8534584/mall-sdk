package pddRequest

import "encoding/json"

type PddDdkGoodsPromotionUrlGenerate struct {
	typeName                  string
	apiParas                  map[string]interface{}
	cashGiftId                uint64
	cashGiftName              string
	customParameters          string
	generateAuthorityUrl      bool
	generateMallCollectCoupon bool
	generateQqApp             bool
	generateSchemaUrl         bool
	generateShortUrl          bool
	generateWeApp             bool
	goodsSignList             []string
	materialId                string
	multiGroup                bool
	pId                       string
	searchId                  string
	zsDuoId                   uint64
}

func (p *PddDdkGoodsPromotionUrlGenerate) ZsDuoId() uint64 {
	return p.zsDuoId
}

func (p *PddDdkGoodsPromotionUrlGenerate) SetZsDuoId(zsDuoId uint64) {
	p.apiParas["zs_duo_id"] = zsDuoId
	p.zsDuoId = zsDuoId
}

func (p *PddDdkGoodsPromotionUrlGenerate) SearchId() string {
	return p.searchId
}

func (p *PddDdkGoodsPromotionUrlGenerate) SetSearchId(searchId string) {
	p.apiParas["search_id"] = searchId
	p.searchId = searchId
}

func (p *PddDdkGoodsPromotionUrlGenerate) CustomParameters() string {
	return p.customParameters
}

func (p *PddDdkGoodsPromotionUrlGenerate) SetCustomParameters(customParameters string) {
	p.apiParas["custom_parameters"] = customParameters
	p.customParameters = customParameters
}

func (p *PddDdkGoodsPromotionUrlGenerate) PId() string {
	return p.pId
}

func (p *PddDdkGoodsPromotionUrlGenerate) SetPId(pId string) {
	p.apiParas["p_id"] = pId
	p.pId = pId
}

func (p *PddDdkGoodsPromotionUrlGenerate) MultiGroup() bool {
	return p.multiGroup
}

func (p *PddDdkGoodsPromotionUrlGenerate) SetMultiGroup(multiGroup bool) {
	p.apiParas["multi_group"] = multiGroup
	p.multiGroup = multiGroup
}

func (p *PddDdkGoodsPromotionUrlGenerate) MaterialId() string {
	return p.materialId
}

func (p *PddDdkGoodsPromotionUrlGenerate) SetMaterialId(materialId string) {
	p.apiParas["material_id"] = materialId
	p.materialId = materialId
}

func (p *PddDdkGoodsPromotionUrlGenerate) GoodsSignList() []string {
	return p.goodsSignList
}

func (p *PddDdkGoodsPromotionUrlGenerate) SetGoodsSignList(goodsSignList []string) {
	goodsSignListStr, _ := json.Marshal(goodsSignList)
	p.apiParas["goods_sign_list"] = string(goodsSignListStr)
	p.goodsSignList = goodsSignList
}

func (p *PddDdkGoodsPromotionUrlGenerate) GenerateWeApp() bool {
	return p.generateWeApp
}

func (p *PddDdkGoodsPromotionUrlGenerate) SetGenerateWeApp(generateWeApp bool) {
	p.apiParas["generate_we_app"] = generateWeApp
	p.generateWeApp = generateWeApp
}

func (p *PddDdkGoodsPromotionUrlGenerate) GenerateShortUrl() bool {
	return p.generateShortUrl
}

func (p *PddDdkGoodsPromotionUrlGenerate) SetGenerateShortUrl(generateShortUrl bool) {
	p.apiParas["generate_short_url"] = generateShortUrl
	p.generateShortUrl = generateShortUrl
}

func (p *PddDdkGoodsPromotionUrlGenerate) GenerateSchemaUrl() bool {
	return p.generateSchemaUrl
}

func (p *PddDdkGoodsPromotionUrlGenerate) SetGenerateSchemaUrl(generateSchemaUrl bool) {
	p.apiParas["generate_schema_url"] = generateSchemaUrl
	p.generateSchemaUrl = generateSchemaUrl
}

func (p *PddDdkGoodsPromotionUrlGenerate) GenerateQqApp() bool {
	return p.generateQqApp
}

func (p *PddDdkGoodsPromotionUrlGenerate) SetGenerateQqApp(generateQqApp bool) {
	p.apiParas["generate_qq_app"] = generateQqApp
	p.generateQqApp = generateQqApp
}

func (p *PddDdkGoodsPromotionUrlGenerate) GenerateMallCollectCoupon() bool {
	return p.generateMallCollectCoupon
}

func (p *PddDdkGoodsPromotionUrlGenerate) SetGenerateMallCollectCoupon(generateMallCollectCoupon bool) {
	p.apiParas["generate_mall_collect_coupon"] = generateMallCollectCoupon
	p.generateMallCollectCoupon = generateMallCollectCoupon
}

func (p *PddDdkGoodsPromotionUrlGenerate) GenerateAuthorityUrl() bool {
	return p.generateAuthorityUrl
}

func (p *PddDdkGoodsPromotionUrlGenerate) SetGenerateAuthorityUrl(generateAuthorityUrl bool) {
	p.apiParas["generate_authority_url"] = generateAuthorityUrl
	p.generateAuthorityUrl = generateAuthorityUrl
}

func (p *PddDdkGoodsPromotionUrlGenerate) CashGiftName() string {
	return p.cashGiftName
}

func (p *PddDdkGoodsPromotionUrlGenerate) SetCashGiftName(cashGiftName string) {
	p.apiParas["cash_gift_name"] = cashGiftName
	p.cashGiftName = cashGiftName
}

func (p *PddDdkGoodsPromotionUrlGenerate) CashGiftId() uint64 {
	return p.cashGiftId
}

func (p *PddDdkGoodsPromotionUrlGenerate) SetCashGiftId(cashGiftId uint64) {
	p.apiParas["cash_gift_id"] = cashGiftId
	p.cashGiftId = cashGiftId
}

func (p *PddDdkGoodsPromotionUrlGenerate) ApiParas() map[string]interface{} {
	return p.apiParas
}

func (p *PddDdkGoodsPromotionUrlGenerate) SetApiParas(apiParas map[string]interface{}) {
	p.apiParas = apiParas
}

func (p *PddDdkGoodsPromotionUrlGenerate) TypeName() string {
	return p.typeName
}

func (p *PddDdkGoodsPromotionUrlGenerate) SetTypeName(typeName string) {
	p.apiParas["type"] = typeName
	p.typeName = typeName
}

func NewPddDdkGoodsPromotionUrlGenerate() *PddDdkGoodsPromotionUrlGenerate {
	p := &PddDdkGoodsPromotionUrlGenerate{
		apiParas: make(map[string]interface{}),
	}
	p.SetTypeName(p.GetApiMethodName())
	return p
}

func (p *PddDdkGoodsPromotionUrlGenerate) GetApiParams() map[string]interface{} {
	return p.apiParas
}

func (p *PddDdkGoodsPromotionUrlGenerate) GetApiMethodName() string {
	return "pdd.ddk.goods.promotion.url.generate"
}
