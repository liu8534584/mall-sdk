package domain


type TaobaoTbkScMaterialRecommendBasicMapData struct {
    /*
        商品信息-商品标题     */
    Title  *string `json:"title,omitempty" `

    /*
        商品信息-商品短标题     */
    ShortTitle  *string `json:"short_title,omitempty" `

    /*
        商品信息-商品主图     */
    PictUrl  *string `json:"pict_url,omitempty" `

    /*
        商品信息-商品白底图     */
    WhiteImage  *string `json:"white_image,omitempty" `

    /*
        商品信息-一级类目ID     */
    LevelOneCategoryId  *int64 `json:"level_one_category_id,omitempty" `

    /*
        商品信息-一级类目名称     */
    LevelOneCategoryName  *string `json:"level_one_category_name,omitempty" `

    /*
        商品信息-叶子类目id     */
    CategoryId  *int64 `json:"category_id,omitempty" `

    /*
        商品信息-叶子类目名称     */
    CategoryName  *string `json:"category_name,omitempty" `

    /*
        店铺信息-卖家id     */
    SellerId  *int64 `json:"seller_id,omitempty" `

    /*
        店铺信息-卖家类型，0表示淘宝，1表示天猫，3表示特价版     */
    UserType  *int64 `json:"user_type,omitempty" `

    /*
        店铺信息-店铺名称     */
    ShopTitle  *string `json:"shop_title,omitempty" `

    /*
        商品信息-月销量     */
    Volume  *int64 `json:"volume,omitempty" `

    /*
        商品信息-商品子标题     */
    SubTitle  *string `json:"sub_title,omitempty" `

    /*
        商品信息-品牌名称     */
    BrandName  *string `json:"brand_name,omitempty" `

}

func (s *TaobaoTbkScMaterialRecommendBasicMapData) SetTitle(v string) *TaobaoTbkScMaterialRecommendBasicMapData {
    s.Title = &v
    return s
}
func (s *TaobaoTbkScMaterialRecommendBasicMapData) SetShortTitle(v string) *TaobaoTbkScMaterialRecommendBasicMapData {
    s.ShortTitle = &v
    return s
}
func (s *TaobaoTbkScMaterialRecommendBasicMapData) SetPictUrl(v string) *TaobaoTbkScMaterialRecommendBasicMapData {
    s.PictUrl = &v
    return s
}
func (s *TaobaoTbkScMaterialRecommendBasicMapData) SetWhiteImage(v string) *TaobaoTbkScMaterialRecommendBasicMapData {
    s.WhiteImage = &v
    return s
}
func (s *TaobaoTbkScMaterialRecommendBasicMapData) SetLevelOneCategoryId(v int64) *TaobaoTbkScMaterialRecommendBasicMapData {
    s.LevelOneCategoryId = &v
    return s
}
func (s *TaobaoTbkScMaterialRecommendBasicMapData) SetLevelOneCategoryName(v string) *TaobaoTbkScMaterialRecommendBasicMapData {
    s.LevelOneCategoryName = &v
    return s
}
func (s *TaobaoTbkScMaterialRecommendBasicMapData) SetCategoryId(v int64) *TaobaoTbkScMaterialRecommendBasicMapData {
    s.CategoryId = &v
    return s
}
func (s *TaobaoTbkScMaterialRecommendBasicMapData) SetCategoryName(v string) *TaobaoTbkScMaterialRecommendBasicMapData {
    s.CategoryName = &v
    return s
}
func (s *TaobaoTbkScMaterialRecommendBasicMapData) SetSellerId(v int64) *TaobaoTbkScMaterialRecommendBasicMapData {
    s.SellerId = &v
    return s
}
func (s *TaobaoTbkScMaterialRecommendBasicMapData) SetUserType(v int64) *TaobaoTbkScMaterialRecommendBasicMapData {
    s.UserType = &v
    return s
}
func (s *TaobaoTbkScMaterialRecommendBasicMapData) SetShopTitle(v string) *TaobaoTbkScMaterialRecommendBasicMapData {
    s.ShopTitle = &v
    return s
}
func (s *TaobaoTbkScMaterialRecommendBasicMapData) SetVolume(v int64) *TaobaoTbkScMaterialRecommendBasicMapData {
    s.Volume = &v
    return s
}
func (s *TaobaoTbkScMaterialRecommendBasicMapData) SetSubTitle(v string) *TaobaoTbkScMaterialRecommendBasicMapData {
    s.SubTitle = &v
    return s
}
func (s *TaobaoTbkScMaterialRecommendBasicMapData) SetBrandName(v string) *TaobaoTbkScMaterialRecommendBasicMapData {
    s.BrandName = &v
    return s
}
