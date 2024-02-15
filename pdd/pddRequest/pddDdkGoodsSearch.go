package pddRequest

import (
	"encoding/json"
	"github.com/spf13/cast"
)

type pddDdkGoodsSearch struct {
	typeName string
	keyword  string

	activityTags      []uint32
	blockCVatPackages []uint32
	blockCats         []uint32
	catId             uint64
	customParameters  string
	goodsImgType      uint8
	goodsSignList     []string
	isBrandGoods      bool
	listId            string
	merchantType      uint8
	merchantTypeList  []uint8
	optId             uint64
	page              uint64
	pageSize          uint64
	pid               string
	rangeList         string
	sortType          uint8
	useCustomized     bool
	withCoupon        bool

	apiParas map[string]interface{}
}

func (p *pddDdkGoodsSearch) RangeList() string {
	return p.rangeList
}

func (p *pddDdkGoodsSearch) SetRangeList(rangeFrom uint64, rangeTo int, rangeId int) {
	var rangeListData map[string]interface{}
	rangeListData = make(map[string]interface{})
	rangeListData["range_from"] = rangeFrom
	rangeListData["range_id"] = cast.ToString(rangeId)
	rangeListData["range_to"] = rangeTo
	rangeListStr, _ := json.Marshal([]map[string]interface{}{rangeListData})
	p.rangeList = cast.ToString(rangeListStr)
	p.apiParas["range_list"] = p.rangeList
}

func (p *pddDdkGoodsSearch) WithCoupon() bool {
	return p.withCoupon
}

func (p *pddDdkGoodsSearch) SetWithCoupon(withCoupon bool) {
	p.apiParas["with_coupon"] = withCoupon
	p.withCoupon = withCoupon
}

func (p *pddDdkGoodsSearch) UseCustomized() bool {
	return p.useCustomized
}

func (p *pddDdkGoodsSearch) SetUseCustomized(useCustomized bool) {
	p.apiParas["use_customized"] = useCustomized
	p.useCustomized = useCustomized
}

func (p *pddDdkGoodsSearch) SortType() uint8 {
	return p.sortType
}

func (p *pddDdkGoodsSearch) SetSortType(sortType uint8) {
	p.apiParas["sort_type"] = sortType
	p.sortType = sortType
}

func (p *pddDdkGoodsSearch) OptId() uint64 {
	return p.optId
}

func (p *pddDdkGoodsSearch) SetOptId(optId uint64) {
	p.apiParas["opt_id"] = optId
	p.optId = optId
}

func (p *pddDdkGoodsSearch) MerchantTypeList() []uint8 {
	return p.merchantTypeList
}

func (p *pddDdkGoodsSearch) SetMerchantTypeList(merchantTypeList []uint8) {
	merchantTypeListStr, _ := json.Marshal(merchantTypeList)
	p.apiParas["merchant_type_list"] = string(merchantTypeListStr)
	p.merchantTypeList = merchantTypeList
}

func (p *pddDdkGoodsSearch) MerchantType() uint8 {
	return p.merchantType
}

func (p *pddDdkGoodsSearch) SetMerchantType(merchantType uint8) {
	p.apiParas["merchant_type"] = merchantType
	p.merchantType = merchantType
}

func (p *pddDdkGoodsSearch) ListId() string {
	return p.listId
}

func (p *pddDdkGoodsSearch) SetListId(listId string) {
	p.apiParas["list_id"] = listId
	p.listId = listId
}

func (p *pddDdkGoodsSearch) IsBrandGoods() bool {
	return p.isBrandGoods
}

func (p *pddDdkGoodsSearch) SetIsBrandGoods(isBrandGoods bool) {
	p.apiParas["is_brand_goods"] = isBrandGoods
	p.isBrandGoods = isBrandGoods
}

func (p *pddDdkGoodsSearch) GoodsSignList() []string {
	return p.goodsSignList
}

func (p *pddDdkGoodsSearch) SetGoodsSignList(goodsSignList []string) {
	goodsSignListStr, _ := json.Marshal(goodsSignList)
	p.apiParas["goods_sign_list"] = string(goodsSignListStr)
	p.goodsSignList = goodsSignList
}

func (p *pddDdkGoodsSearch) GoodsImgType() uint8 {
	return p.goodsImgType
}

func (p *pddDdkGoodsSearch) SetGoodsImgType(goodsImgType uint8) {
	p.apiParas["goods_img_type"] = goodsImgType
	p.goodsImgType = goodsImgType
}

func (p *pddDdkGoodsSearch) CatId() uint64 {
	return p.catId
}

func (p *pddDdkGoodsSearch) SetCatId(catId uint64) {
	p.apiParas["cat_id"] = catId
	p.catId = catId
}

func (p *pddDdkGoodsSearch) BlockCats() []uint32 {
	return p.blockCats
}

func (p *pddDdkGoodsSearch) SetBlockCats(blockCats []uint32) {
	blockCatsStr, _ := json.Marshal(blockCats)
	p.apiParas["block_cats"] = string(blockCatsStr)
	p.blockCats = blockCats
}

func (p *pddDdkGoodsSearch) ActivityTags() []uint32 {
	return p.activityTags
}

func (p *pddDdkGoodsSearch) SetActivityTags(activityTags []uint32) {
	activityTagsStr, _ := json.Marshal(activityTags)
	p.apiParas["activity_tags"] = string(activityTagsStr)
	p.activityTags = activityTags
}

func (p *pddDdkGoodsSearch) PageSize() uint64 {
	return p.pageSize
}

func (p *pddDdkGoodsSearch) SetPageSize(pageSize uint64) {
	p.apiParas["page_size"] = pageSize
	p.pageSize = pageSize
}

func (p *pddDdkGoodsSearch) Page() uint64 {
	return p.page
}

func (p *pddDdkGoodsSearch) SetPage(page uint64) {
	p.apiParas["page"] = page
	p.page = page
}

func (p *pddDdkGoodsSearch) Pid() string {
	return p.pid
}

func (p *pddDdkGoodsSearch) SetPid(pid string) {
	p.apiParas["pid"] = pid
	p.pid = pid
}

func (p *pddDdkGoodsSearch) CustomParameters() string {
	return p.customParameters
}

func (p *pddDdkGoodsSearch) SetCustomParameters(customParameters string) {
	p.apiParas["custom_parameters"] = customParameters
	p.customParameters = customParameters
}

func (p *pddDdkGoodsSearch) BlockCVatPackages() []uint32 {
	return p.blockCVatPackages
}

func (p *pddDdkGoodsSearch) SetBlockCVatPackages(blockCVatPackages []uint32) {
	blockCVatPackagesStr, _ := json.Marshal(blockCVatPackages)
	p.apiParas["block_cat_packages"] = string(blockCVatPackagesStr)
	p.blockCVatPackages = blockCVatPackages
}

func (p *pddDdkGoodsSearch) Keyword() string {
	return p.keyword
}

func (p *pddDdkGoodsSearch) SetKeyword(keyword string) {
	p.apiParas["keyword"] = keyword
	p.keyword = keyword
}

func (p *pddDdkGoodsSearch) TypeName() string {
	return p.typeName
}

func (p *pddDdkGoodsSearch) SetTypeName(typeName string) {
	p.apiParas["type"] = typeName
	p.typeName = typeName
}

func NewPddDdkGoodsSearch() *pddDdkGoodsSearch {
	p := &pddDdkGoodsSearch{
		apiParas: make(map[string]interface{}),
	}
	p.SetTypeName(p.GetApiMethodName())
	return p
}

func (p *pddDdkGoodsSearch) GetApiParams() map[string]interface{} {
	return p.apiParas
}

func (p *pddDdkGoodsSearch) GetApiMethodName() string {
	return "pdd.ddk.goods.search"
}
