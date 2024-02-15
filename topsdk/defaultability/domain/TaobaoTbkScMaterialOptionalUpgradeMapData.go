package domain


type TaobaoTbkScMaterialOptionalUpgradeMapData struct {
    /*
        商品信息-淘宝客新商品id；使用说明参考《淘宝客新商品ID升级》白皮书：https://www.yuque.com/taobaolianmengguanfangxiaoer/zmig94/tfyt0pahmlpzu2ud     */
    ItemId  *string `json:"item_id,omitempty" `

    /*
        价格促销信息     */
    PricePromotionInfo  *TaobaoTbkScMaterialOptionalUpgradePromotionInfoMapData `json:"price_promotion_info,omitempty" `

    /*
        淘客推广信息     */
    PublishInfo  *TaobaoTbkScMaterialOptionalUpgradePublishInfo `json:"publish_info,omitempty" `

    /*
        商品基础信息     */
    ItemBasicInfo  *TaobaoTbkScMaterialOptionalUpgradeBasicMapData `json:"item_basic_info,omitempty" `

    /*
        天猫榜单信息     */
    TmallRankInfo  *TaobaoTbkScMaterialOptionalUpgradeTmallRankInfo `json:"tmall_rank_info,omitempty" `

    /*
        预售信息     */
    PresaleInfo  *TaobaoTbkScMaterialOptionalUpgradePresaleInfo `json:"presale_info,omitempty" `

    /*
        商品库范围信息     */
    ScopeInfo  *TaobaoTbkScMaterialOptionalUpgradeScopeInfo `json:"scope_info,omitempty" `

    /*
        线报内容     */
    MgcInfo  *TaobaoTbkScMaterialOptionalUpgradeMgcInfo `json:"mgc_info,omitempty" `

}

func (s *TaobaoTbkScMaterialOptionalUpgradeMapData) SetItemId(v string) *TaobaoTbkScMaterialOptionalUpgradeMapData {
    s.ItemId = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeMapData) SetPricePromotionInfo(v TaobaoTbkScMaterialOptionalUpgradePromotionInfoMapData) *TaobaoTbkScMaterialOptionalUpgradeMapData {
    s.PricePromotionInfo = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeMapData) SetPublishInfo(v TaobaoTbkScMaterialOptionalUpgradePublishInfo) *TaobaoTbkScMaterialOptionalUpgradeMapData {
    s.PublishInfo = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeMapData) SetItemBasicInfo(v TaobaoTbkScMaterialOptionalUpgradeBasicMapData) *TaobaoTbkScMaterialOptionalUpgradeMapData {
    s.ItemBasicInfo = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeMapData) SetTmallRankInfo(v TaobaoTbkScMaterialOptionalUpgradeTmallRankInfo) *TaobaoTbkScMaterialOptionalUpgradeMapData {
    s.TmallRankInfo = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeMapData) SetPresaleInfo(v TaobaoTbkScMaterialOptionalUpgradePresaleInfo) *TaobaoTbkScMaterialOptionalUpgradeMapData {
    s.PresaleInfo = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeMapData) SetScopeInfo(v TaobaoTbkScMaterialOptionalUpgradeScopeInfo) *TaobaoTbkScMaterialOptionalUpgradeMapData {
    s.ScopeInfo = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeMapData) SetMgcInfo(v TaobaoTbkScMaterialOptionalUpgradeMgcInfo) *TaobaoTbkScMaterialOptionalUpgradeMapData {
    s.MgcInfo = &v
    return s
}
