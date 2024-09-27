package domain


type TaobaoTbkDgMaterialRecommendFavoritesDetail struct {
    /*
        选品库收藏夹id     */
    FavoritesId  *int64 `json:"favorites_id,omitempty" `

    /*
        选品库收藏夹标题     */
    FavoritesTitle  *string `json:"favorites_title,omitempty" `

}

func (s *TaobaoTbkDgMaterialRecommendFavoritesDetail) SetFavoritesId(v int64) *TaobaoTbkDgMaterialRecommendFavoritesDetail {
    s.FavoritesId = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendFavoritesDetail) SetFavoritesTitle(v string) *TaobaoTbkDgMaterialRecommendFavoritesDetail {
    s.FavoritesTitle = &v
    return s
}
