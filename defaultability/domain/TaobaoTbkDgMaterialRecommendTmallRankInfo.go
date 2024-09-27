package domain


type TaobaoTbkDgMaterialRecommendTmallRankInfo struct {
    /*
        榜单排行描述     */
    TmallRankText  *string `json:"tmall_rank_text,omitempty" `

    /*
        榜单url     */
    TmallRankUrl  *string `json:"tmall_rank_url,omitempty" `

}

func (s *TaobaoTbkDgMaterialRecommendTmallRankInfo) SetTmallRankText(v string) *TaobaoTbkDgMaterialRecommendTmallRankInfo {
    s.TmallRankText = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendTmallRankInfo) SetTmallRankUrl(v string) *TaobaoTbkDgMaterialRecommendTmallRankInfo {
    s.TmallRankUrl = &v
    return s
}
