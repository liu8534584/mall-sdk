package pddRequest

import "encoding/json"

type PddDdkRpPromUrlGenerate struct {
	typeName            string
	apiParas            map[string]interface{}
	amount              uint64
	channelType         uint64
	customParameters    string
	diyOneYuanParam     string
	diyRedPacketParam   string
	diySpRedPacketParam string
	generateQqApp       bool
	generateSchemaUrl   bool
	generateShortUrl    bool
	generateWeApp       bool
	pIdList             string
	scratchCardAmount   uint32
	zsDuoId             uint64
	tmccParam           string
}

func (p *PddDdkRpPromUrlGenerate) TmccParam() string {
	return p.tmccParam
}

func (p *PddDdkRpPromUrlGenerate) DiyRedPacketParam() string {
	return p.diyRedPacketParam
}

func (p *PddDdkRpPromUrlGenerate) SetTmccParam(goodsSignList []string, configId int64) {
	data := make(map[string]interface{})
	if len(goodsSignList) > 0 {
		data["goods_signs"] = goodsSignList
	}
	if configId > 0 {
		data["tmc_config_id"] = configId
	}

	dataByte, _ := json.Marshal(data)
	dataStr := string(dataByte)
	p.apiParas["tmcc_param"] = dataStr
	p.tmccParam = dataStr
}

func (p *PddDdkRpPromUrlGenerate) SetDiyRedPacketParam(amountProbability []uint32, disText bool, notShowBackground bool, optId int, rangeFrom int, rangeId int, rangTo int) {
	var diyRedPacketParamData map[string]interface{}
	diyRedPacketParamData = make(map[string]interface{})
	diyRedPacketParamData["amount_probability"] = amountProbability
	diyRedPacketParamData["dis_text"] = disText
	diyRedPacketParamData["not_show_background"] = notShowBackground
	diyRedPacketParamData["opt_id"] = optId

	var rangeItems map[string]interface{}
	rangeItems = make(map[string]interface{})
	rangeItems["range_from"] = rangeFrom
	rangeItems["range_id"] = rangeId
	rangeItems["range_to"] = rangTo

	diyRedPacketParamData["range_items"] = rangeItems

	diyRedPacketParamStr, _ := json.Marshal(diyRedPacketParamData)
	p.apiParas["diy_red_packet_param"] = string(diyRedPacketParamStr)
	p.diyRedPacketParam = string(diyRedPacketParamStr)
}

func (p *PddDdkRpPromUrlGenerate) SetDiySpRedPacketParam(goodsSign string) {
	diySpRedPacketParam := make(map[string]string)
	diySpRedPacketParam["goods_sign"] = goodsSign
	diySpRedPacketParamStr, _ := json.Marshal(diySpRedPacketParam)
	p.apiParas["diy_sp_red_packet_param"] = string(diySpRedPacketParamStr)
	p.diySpRedPacketParam = string(diySpRedPacketParamStr)
}

func (p *PddDdkRpPromUrlGenerate) DiyOneYuanParam() string {
	return p.diyOneYuanParam
}

func (p *PddDdkRpPromUrlGenerate) SetDiyOneYuanParam(goodsSign string) {
	var diyOneYuanParamData map[string]string
	diyOneYuanParamData = make(map[string]string)
	diyOneYuanParamData["goods_sign"] = goodsSign
	diyOneYuanParamStr, _ := json.Marshal(diyOneYuanParamData)
	p.apiParas["diy_one_yuan_param"] = string(diyOneYuanParamStr)
	p.diyOneYuanParam = string(diyOneYuanParamStr)
}

func (p *PddDdkRpPromUrlGenerate) GenerateQqApp() bool {
	return p.generateQqApp
}

func (p *PddDdkRpPromUrlGenerate) SetGenerateQqApp(generateQqApp bool) {
	p.apiParas["generate_qq_app"] = generateQqApp
	p.generateQqApp = generateQqApp
}

func (p *PddDdkRpPromUrlGenerate) GenerateSchemaUrl() bool {
	return p.generateSchemaUrl
}

func (p *PddDdkRpPromUrlGenerate) SetGenerateSchemaUrl(generateSchemaUrl bool) {
	p.apiParas["generate_schema_url"] = generateSchemaUrl
	p.generateSchemaUrl = generateSchemaUrl
}

func (p *PddDdkRpPromUrlGenerate) SetZsDuoId(zsDuoId uint64) {
	p.apiParas["zs_duo_id"] = zsDuoId
	p.zsDuoId = zsDuoId
}

func (p *PddDdkRpPromUrlGenerate) GenerateZsDuoId(zsDuoId uint64) uint64 {
	return zsDuoId
}

func (p *PddDdkRpPromUrlGenerate) GenerateShortUrl() bool {
	return p.generateShortUrl
}

func (p *PddDdkRpPromUrlGenerate) SetGenerateShortUrl(generateShortUrl bool) {
	p.apiParas["generate_short_url"] = generateShortUrl
	p.generateShortUrl = generateShortUrl
}

func (p *PddDdkRpPromUrlGenerate) GenerateWeApp() bool {
	return p.generateWeApp
}

func (p *PddDdkRpPromUrlGenerate) SetGenerateWeApp(generateWeApp bool) {
	p.apiParas["generate_we_app"] = generateWeApp
	p.generateWeApp = generateWeApp
}

func (p *PddDdkRpPromUrlGenerate) PIdList() string {
	return p.pIdList
}

func (p *PddDdkRpPromUrlGenerate) SetPIdList(pIdList []string) {
	pIdListStr, _ := json.Marshal(pIdList)

	p.apiParas["p_id_list"] = string(pIdListStr)
	p.pIdList = string(pIdListStr)
}

func (p *PddDdkRpPromUrlGenerate) ScratchCardAmount() uint32 {
	return p.scratchCardAmount
}

func (p *PddDdkRpPromUrlGenerate) SetScratchCardAmount(scratchCardAmount uint32) {
	p.apiParas["scratch_card_amount"] = scratchCardAmount
	p.scratchCardAmount = scratchCardAmount
}

func (p *PddDdkRpPromUrlGenerate) Amount() uint64 {
	return p.amount
}

func (p *PddDdkRpPromUrlGenerate) SetAmount(amount uint64) {
	p.apiParas["amount"] = amount
	p.amount = amount
}

func (p *PddDdkRpPromUrlGenerate) ChannelType() uint64 {
	return p.channelType
}

func (p *PddDdkRpPromUrlGenerate) SetChannelType(channelType uint64) {
	p.apiParas["channel_type"] = channelType
	p.channelType = channelType
}

func (p *PddDdkRpPromUrlGenerate) CustomParameters() string {
	return p.customParameters
}

func (p *PddDdkRpPromUrlGenerate) SetCustomParameters(customParameters string) {
	p.apiParas["custom_parameters"] = customParameters
	p.customParameters = customParameters
}

func (p *PddDdkRpPromUrlGenerate) TypeName() string {
	return p.typeName
}

func (p *PddDdkRpPromUrlGenerate) SetTypeName(typeName string) {
	p.apiParas["type"] = typeName
	p.typeName = typeName
}

func NewPddDdkRpPromUrlGenerate() *PddDdkRpPromUrlGenerate {
	p := &PddDdkRpPromUrlGenerate{
		apiParas: make(map[string]interface{}),
	}
	p.SetTypeName(p.GetApiMethodName())
	return p
}

func (p *PddDdkRpPromUrlGenerate) GetApiParams() map[string]interface{} {
	return p.apiParas
}

func (p *PddDdkRpPromUrlGenerate) GetApiMethodName() string {
	return "pdd.ddk.rp.prom.url.generate"
}
