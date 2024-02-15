package domain


type TaobaoTbkScMaterialRecommendMapData struct {
    /*
        商品信息-宝贝id     */
    ItemId  *string `json:"item_id,omitempty" `

    /*
        商品基础信息     */
    ItemBasicInfo  *TaobaoTbkScMaterialRecommendBasicMapData `json:"item_basic_info,omitempty" `

    /*
        价格促销信息     */
    PricePromotionInfo  *TaobaoTbkScMaterialRecommendPromotionInfoMapData `json:"price_promotion_info,omitempty" `

    /*
        淘客推广信息     */
    PublishInfo  *TaobaoTbkScMaterialRecommendPublishInfo `json:"publish_info,omitempty" `

}

func (s *TaobaoTbkScMaterialRecommendMapData) SetItemId(v string) *TaobaoTbkScMaterialRecommendMapData {
    s.ItemId = &v
    return s
}
func (s *TaobaoTbkScMaterialRecommendMapData) SetItemBasicInfo(v TaobaoTbkScMaterialRecommendBasicMapData) *TaobaoTbkScMaterialRecommendMapData {
    s.ItemBasicInfo = &v
    return s
}
func (s *TaobaoTbkScMaterialRecommendMapData) SetPricePromotionInfo(v TaobaoTbkScMaterialRecommendPromotionInfoMapData) *TaobaoTbkScMaterialRecommendMapData {
    s.PricePromotionInfo = &v
    return s
}
func (s *TaobaoTbkScMaterialRecommendMapData) SetPublishInfo(v TaobaoTbkScMaterialRecommendPublishInfo) *TaobaoTbkScMaterialRecommendMapData {
    s.PublishInfo = &v
    return s
}
