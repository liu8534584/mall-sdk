package kaolaRequest

type QueryRecommendGoodsListRequest struct {
	pageIndex  uint64
	sortType   uint8
	categoryId string
	apiParas   map[string]interface{}
}

func NewQueryRecommendGoodsListRequest() *QueryRecommendGoodsListRequest {
	p := &QueryRecommendGoodsListRequest{
		apiParas: make(map[string]interface{}),
	}
	return p
}

func (k *QueryRecommendGoodsListRequest) PageIndex() uint64 {
	return k.pageIndex
}

func (k *QueryRecommendGoodsListRequest) SetPageIndex(pageIndex uint64) {
	k.pageIndex = pageIndex
	k.apiParas["pageIndex"] = pageIndex
}

func (k *QueryRecommendGoodsListRequest) SortType() uint8 {
	return k.sortType
}

func (k *QueryRecommendGoodsListRequest) SetSortType(sortType uint8) {
	k.sortType = sortType
	k.apiParas["sortType"] = sortType
}

func (k *QueryRecommendGoodsListRequest) CategoryId() string {
	return k.categoryId
}

func (k *QueryRecommendGoodsListRequest) SetCategoryId(categoryId string) {
	k.categoryId = categoryId
	k.apiParas["categoryId"] = categoryId
}

func (k *QueryRecommendGoodsListRequest) GetApiMethodName() string {
	return "kaola.zhuanke.api.queryRecommendGoodsList"
}

func (k *QueryRecommendGoodsListRequest) GetApiParams() map[string]interface{} {

	return k.apiParas
}
