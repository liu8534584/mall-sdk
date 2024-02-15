package alliance_materialsProductCategory_response

import (
	"doudian.com/open/sdk_golang/core"
)

type AllianceMaterialsProductCategoryResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *AllianceMaterialsProductCategoryData `json:"data"`
}
type CategoryListItem struct {
	// 类目ID
	Id int64 `json:"id"`
	// 类目名称
	Name string `json:"name"`
	// 类目层级
	Level int64 `json:"level"`
}
type AllianceMaterialsProductCategoryData struct {
	// 类目总数
	Total int64 `json:"total"`
	// 类目列表（包含查询的父类目和下一级子类目）
	CategoryList []CategoryListItem `json:"category_list"`
}
