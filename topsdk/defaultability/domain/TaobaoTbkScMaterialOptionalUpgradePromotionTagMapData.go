package domain


type TaobaoTbkScMaterialOptionalUpgradePromotionTagMapData struct {
    /*
        标签名称，如“88VIP”、“花呗免息”、“猫超买返”     */
    TagName  *string `json:"tag_name,omitempty" `

}

func (s *TaobaoTbkScMaterialOptionalUpgradePromotionTagMapData) SetTagName(v string) *TaobaoTbkScMaterialOptionalUpgradePromotionTagMapData {
    s.TagName = &v
    return s
}
