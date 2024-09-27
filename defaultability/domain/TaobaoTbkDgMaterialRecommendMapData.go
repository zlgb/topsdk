package domain


type TaobaoTbkDgMaterialRecommendMapData struct {
    /*
        商品信息-淘宝客新商品id；使用说明参考《淘宝客新商品ID升级》白皮书：https://www.yuque.com/taobaolianmengguanfangxiaoer/zmig94/tfyt0pahmlpzu2ud     */
    ItemId  *string `json:"item_id,omitempty" `

    /*
        商品基础信息     */
    ItemBasicInfo  *TaobaoTbkDgMaterialRecommendBasicMapData `json:"item_basic_info,omitempty" `

    /*
        价格促销信息     */
    PricePromotionInfo  *TaobaoTbkDgMaterialRecommendPromotionInfoMapData `json:"price_promotion_info,omitempty" `

    /*
        淘客推广信息     */
    PublishInfo  *TaobaoTbkDgMaterialRecommendPublishInfo `json:"publish_info,omitempty" `

    /*
        天猫榜单信息     */
    TmallRankInfo  *TaobaoTbkDgMaterialRecommendTmallRankInfo `json:"tmall_rank_info,omitempty" `

    /*
        预售信息     */
    PresaleInfo  *TaobaoTbkDgMaterialRecommendPresaleInfo `json:"presale_info,omitempty" `

    /*
        选品库信息     */
    FavoritesInfo  *TaobaoTbkDgMaterialRecommendFavoritesInfo `json:"favorites_info,omitempty" `

    /*
        前N件佣金信息-前N件佣金生效或预热时透出以下字段     */
    TopnInfo  *TaobaoTbkDgMaterialRecommendTopNInfoDTO `json:"topn_info,omitempty" `

    /*
        商品库范围信息     */
    ScopeInfo  *TaobaoTbkDgMaterialRecommendScopeInfo `json:"scope_info,omitempty" `

}

func (s *TaobaoTbkDgMaterialRecommendMapData) SetItemId(v string) *TaobaoTbkDgMaterialRecommendMapData {
    s.ItemId = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendMapData) SetItemBasicInfo(v TaobaoTbkDgMaterialRecommendBasicMapData) *TaobaoTbkDgMaterialRecommendMapData {
    s.ItemBasicInfo = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendMapData) SetPricePromotionInfo(v TaobaoTbkDgMaterialRecommendPromotionInfoMapData) *TaobaoTbkDgMaterialRecommendMapData {
    s.PricePromotionInfo = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendMapData) SetPublishInfo(v TaobaoTbkDgMaterialRecommendPublishInfo) *TaobaoTbkDgMaterialRecommendMapData {
    s.PublishInfo = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendMapData) SetTmallRankInfo(v TaobaoTbkDgMaterialRecommendTmallRankInfo) *TaobaoTbkDgMaterialRecommendMapData {
    s.TmallRankInfo = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendMapData) SetPresaleInfo(v TaobaoTbkDgMaterialRecommendPresaleInfo) *TaobaoTbkDgMaterialRecommendMapData {
    s.PresaleInfo = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendMapData) SetFavoritesInfo(v TaobaoTbkDgMaterialRecommendFavoritesInfo) *TaobaoTbkDgMaterialRecommendMapData {
    s.FavoritesInfo = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendMapData) SetTopnInfo(v TaobaoTbkDgMaterialRecommendTopNInfoDTO) *TaobaoTbkDgMaterialRecommendMapData {
    s.TopnInfo = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendMapData) SetScopeInfo(v TaobaoTbkDgMaterialRecommendScopeInfo) *TaobaoTbkDgMaterialRecommendMapData {
    s.ScopeInfo = &v
    return s
}
