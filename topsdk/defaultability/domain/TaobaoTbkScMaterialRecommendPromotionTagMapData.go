package domain


type TaobaoTbkScMaterialRecommendPromotionTagMapData struct {
    /*
        标签名称，如“88VIP”、“花呗免息”、“猫超买返”     */
    TagName  *string `json:"tag_name,omitempty" `

}

func (s *TaobaoTbkScMaterialRecommendPromotionTagMapData) SetTagName(v string) *TaobaoTbkScMaterialRecommendPromotionTagMapData {
    s.TagName = &v
    return s
}
