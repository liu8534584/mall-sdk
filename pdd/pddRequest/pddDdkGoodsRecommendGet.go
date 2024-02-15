package pddRequest

import "encoding/json"

type PddDdkGoodsRecommendGet struct {
	typeName         string
	apiParas         map[string]interface{}
	activityTags     []int64
	catId            int64
	channelType      int64
	customParameters string
	goodsImgType     int
	goodsSignList    []string
	limit            uint64
	listId           string
	offset           uint64
	pid              string
}

func (p *PddDdkGoodsRecommendGet) Pid() string {
	return p.pid
}

func (p *PddDdkGoodsRecommendGet) SetPid(pid string) {
	p.apiParas["pid"] = pid
	p.pid = pid
}

func (p *PddDdkGoodsRecommendGet) Offset() uint64 {
	return p.offset
}

func (p *PddDdkGoodsRecommendGet) SetOffset(offset uint64) {
	p.apiParas["offset"] = offset
	p.offset = offset
}

func (p *PddDdkGoodsRecommendGet) ListId() string {
	return p.listId
}

func (p *PddDdkGoodsRecommendGet) SetListId(listId string) {
	p.apiParas["list_id"] = listId
	p.listId = listId
}

func (p *PddDdkGoodsRecommendGet) Limit() uint64 {
	return p.limit
}

func (p *PddDdkGoodsRecommendGet) SetLimit(limit uint64) {
	p.apiParas["limit"] = limit
	p.limit = limit
}

func (p *PddDdkGoodsRecommendGet) GoodsSignList() []string {
	return p.goodsSignList
}

func (p *PddDdkGoodsRecommendGet) SetGoodsSignList(goodsSignList []string) {
	goodsSignListStr, _ := json.Marshal(goodsSignList)
	p.apiParas["goods_sign_list"] = string(goodsSignListStr)
	p.goodsSignList = goodsSignList
}

func (p *PddDdkGoodsRecommendGet) GoodsImgType() int {
	return p.goodsImgType
}

func (p *PddDdkGoodsRecommendGet) SetGoodsImgType(goodsImgType int) {
	p.apiParas["goods_img_type"] = goodsImgType
	p.goodsImgType = goodsImgType
}

func (p *PddDdkGoodsRecommendGet) CustomParameters() string {
	return p.customParameters
}

func (p *PddDdkGoodsRecommendGet) SetCustomParameters(customParameters string) {
	p.apiParas["custom_parameters"] = customParameters
	p.customParameters = customParameters
}

func (p *PddDdkGoodsRecommendGet) ChannelType() int64 {
	return p.channelType
}

func (p *PddDdkGoodsRecommendGet) SetChannelType(channelType int64) {
	p.apiParas["channel_type"] = channelType
	p.channelType = channelType
}

func (p *PddDdkGoodsRecommendGet) CatId() int64 {
	return p.catId
}

func (p *PddDdkGoodsRecommendGet) SetCatId(catId int64) {
	p.apiParas["cat_id"] = catId
	p.catId = catId
}

func (p *PddDdkGoodsRecommendGet) ActivityTags() []int64 {
	return p.activityTags
}

func (p *PddDdkGoodsRecommendGet) SetActivityTags(activityTags []int64) {
	activityTagsStr, _ := json.Marshal(activityTags)
	p.apiParas["activity_tags"] = string(activityTagsStr)
	p.activityTags = activityTags
}

func (p *PddDdkGoodsRecommendGet) TypeName() string {
	return p.typeName
}

func (p *PddDdkGoodsRecommendGet) SetTypeName(typeName string) {
	p.apiParas["type"] = typeName
	p.typeName = typeName
}

func NewPddDdkGoodsRecommendGet() *PddDdkGoodsRecommendGet {
	p := &PddDdkGoodsRecommendGet{
		apiParas: make(map[string]interface{}),
	}

	p.SetTypeName(p.GetApiMethodName())
	return p
}

func (p *PddDdkGoodsRecommendGet) GetApiParams() map[string]interface{} {
	return p.apiParas
}

func (p *PddDdkGoodsRecommendGet) GetApiMethodName() string {
	return "pdd.ddk.goods.recommend.get"
}
