package domain


type TaobaoTbkScMaterialOptionalUpgradeScopeInfo struct {
    /*
        是否品牌精选，0不是，1是     */
    SuperiorBrand  *int64 `json:"superior_brand,omitempty" `

}

func (s *TaobaoTbkScMaterialOptionalUpgradeScopeInfo) SetSuperiorBrand(v int64) *TaobaoTbkScMaterialOptionalUpgradeScopeInfo {
    s.SuperiorBrand = &v
    return s
}
