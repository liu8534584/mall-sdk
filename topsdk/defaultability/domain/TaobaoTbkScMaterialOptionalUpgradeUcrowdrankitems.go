package domain


type TaobaoTbkScMaterialOptionalUpgradeUcrowdrankitems struct {
    /*
        物料评估-商品ID，material_id=41377时必填     */
    ItemId  *string `json:"item_id,omitempty" `

    /*
        物料评估-商品佣金率，如：1234表示12.34%，material_id=41377时选填     */
    Commirate  *int64 `json:"commirate,omitempty" `

    /*
        物料评估-商品价格，单位：元，material_id=41377时选填     */
    Price  *string `json:"price,omitempty" `

}

func (s *TaobaoTbkScMaterialOptionalUpgradeUcrowdrankitems) SetItemId(v string) *TaobaoTbkScMaterialOptionalUpgradeUcrowdrankitems {
    s.ItemId = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeUcrowdrankitems) SetCommirate(v int64) *TaobaoTbkScMaterialOptionalUpgradeUcrowdrankitems {
    s.Commirate = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeUcrowdrankitems) SetPrice(v string) *TaobaoTbkScMaterialOptionalUpgradeUcrowdrankitems {
    s.Price = &v
    return s
}
