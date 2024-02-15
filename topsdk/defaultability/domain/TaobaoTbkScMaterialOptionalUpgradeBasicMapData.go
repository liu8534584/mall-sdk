package domain


type TaobaoTbkScMaterialOptionalUpgradeBasicMapData struct {
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
        商品信息-30天销量；数据统计截止昨日非实时更新     */
    Volume  *int64 `json:"volume,omitempty" `

    /*
        商品信息-商品子标题     */
    SubTitle  *string `json:"sub_title,omitempty" `

    /*
        商品信息-品牌名称     */
    BrandName  *string `json:"brand_name,omitempty" `

    /*
        商品信息-一级类目名称     */
    LevelOneCategoryName  *string `json:"level_one_category_name,omitempty" `

}

func (s *TaobaoTbkScMaterialOptionalUpgradeBasicMapData) SetTitle(v string) *TaobaoTbkScMaterialOptionalUpgradeBasicMapData {
    s.Title = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeBasicMapData) SetShortTitle(v string) *TaobaoTbkScMaterialOptionalUpgradeBasicMapData {
    s.ShortTitle = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeBasicMapData) SetPictUrl(v string) *TaobaoTbkScMaterialOptionalUpgradeBasicMapData {
    s.PictUrl = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeBasicMapData) SetWhiteImage(v string) *TaobaoTbkScMaterialOptionalUpgradeBasicMapData {
    s.WhiteImage = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeBasicMapData) SetLevelOneCategoryId(v int64) *TaobaoTbkScMaterialOptionalUpgradeBasicMapData {
    s.LevelOneCategoryId = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeBasicMapData) SetCategoryId(v int64) *TaobaoTbkScMaterialOptionalUpgradeBasicMapData {
    s.CategoryId = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeBasicMapData) SetCategoryName(v string) *TaobaoTbkScMaterialOptionalUpgradeBasicMapData {
    s.CategoryName = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeBasicMapData) SetSellerId(v int64) *TaobaoTbkScMaterialOptionalUpgradeBasicMapData {
    s.SellerId = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeBasicMapData) SetUserType(v int64) *TaobaoTbkScMaterialOptionalUpgradeBasicMapData {
    s.UserType = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeBasicMapData) SetShopTitle(v string) *TaobaoTbkScMaterialOptionalUpgradeBasicMapData {
    s.ShopTitle = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeBasicMapData) SetVolume(v int64) *TaobaoTbkScMaterialOptionalUpgradeBasicMapData {
    s.Volume = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeBasicMapData) SetSubTitle(v string) *TaobaoTbkScMaterialOptionalUpgradeBasicMapData {
    s.SubTitle = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeBasicMapData) SetBrandName(v string) *TaobaoTbkScMaterialOptionalUpgradeBasicMapData {
    s.BrandName = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeBasicMapData) SetLevelOneCategoryName(v string) *TaobaoTbkScMaterialOptionalUpgradeBasicMapData {
    s.LevelOneCategoryName = &v
    return s
}
