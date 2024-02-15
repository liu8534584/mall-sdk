package domain


type TaobaoTbkScMaterialOptionalUpgradeTmallRankInfo struct {
    /*
        榜单排行描述     */
    TmallRankText  *string `json:"tmall_rank_text,omitempty" `

    /*
        榜单url     */
    TmallRankUrl  *string `json:"tmall_rank_url,omitempty" `

}

func (s *TaobaoTbkScMaterialOptionalUpgradeTmallRankInfo) SetTmallRankText(v string) *TaobaoTbkScMaterialOptionalUpgradeTmallRankInfo {
    s.TmallRankText = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeTmallRankInfo) SetTmallRankUrl(v string) *TaobaoTbkScMaterialOptionalUpgradeTmallRankInfo {
    s.TmallRankUrl = &v
    return s
}
