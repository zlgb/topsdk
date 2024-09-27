package domain


type TaobaoTbkDgMaterialRecommendFavoritesInfo struct {
    /*
        选品库收藏夹总数量     */
    TotalCount  *int64 `json:"total_count,omitempty" `

    /*
        选品库收藏夹详细信息     */
    FavoritesList  *[]TaobaoTbkDgMaterialRecommendFavoritesDetail `json:"favorites_list,omitempty" `

}

func (s *TaobaoTbkDgMaterialRecommendFavoritesInfo) SetTotalCount(v int64) *TaobaoTbkDgMaterialRecommendFavoritesInfo {
    s.TotalCount = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendFavoritesInfo) SetFavoritesList(v []TaobaoTbkDgMaterialRecommendFavoritesDetail) *TaobaoTbkDgMaterialRecommendFavoritesInfo {
    s.FavoritesList = &v
    return s
}
