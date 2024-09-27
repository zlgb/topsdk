package domain


type TaobaoTbkDgMaterialOptionalUpgradeTmallRankInfo struct {
    /*
        榜单排行描述     */
    TmallRankText  *string `json:"tmall_rank_text,omitempty" `

    /*
        榜单url     */
    TmallRankUrl  *string `json:"tmall_rank_url,omitempty" `

}

func (s *TaobaoTbkDgMaterialOptionalUpgradeTmallRankInfo) SetTmallRankText(v string) *TaobaoTbkDgMaterialOptionalUpgradeTmallRankInfo {
    s.TmallRankText = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradeTmallRankInfo) SetTmallRankUrl(v string) *TaobaoTbkDgMaterialOptionalUpgradeTmallRankInfo {
    s.TmallRankUrl = &v
    return s
}
