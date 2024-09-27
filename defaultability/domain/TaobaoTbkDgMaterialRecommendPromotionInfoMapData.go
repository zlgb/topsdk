package domain


type TaobaoTbkDgMaterialRecommendPromotionInfoMapData struct {
    /*
        商品信息-一口价通常显示为划线价     */
    ReservePrice  *string `json:"reserve_price,omitempty" `

    /*
        促销信息-销售价格，无促销时等于一口价，有促销时为促销价。若属于预售商品，付定金时间内，在线售卖价=预售价     */
    ZkFinalPrice  *string `json:"zk_final_price,omitempty" `

    /*
        促销信息-预估到手价(元)。若属于预售商品，付定金时间内，预估到手价价=定金+尾款的预估到手价     */
    FinalPromotionPrice  *string `json:"final_promotion_price,omitempty" `

    /*
        到手价优惠路径列表     */
    FinalPromotionPathList  *[]TaobaoTbkDgMaterialRecommendFinalPromotionPathMapData `json:"final_promotion_path_list,omitempty" `

    /*
        预热预估到手价（元）     */
    FutureActivityPromotionPrice  *string `json:"future_activity_promotion_price,omitempty" `

    /*
        预热到手价优惠路径列表     */
    FutureActivityPromotionPathList  *[]TaobaoTbkDgMaterialRecommendFutureActivityPromotionPathMapData `json:"future_activity_promotion_path_list,omitempty" `

    /*
        标签信息列表     */
    PromotionTagList  *[]TaobaoTbkDgMaterialRecommendPromotionTagMapData `json:"promotion_tag_list,omitempty" `

    /*
        促销信息-预估凑单价（元）。预估凑单叠加优惠后的商品单价     */
    PredictRoundingUpPrice  *string `json:"predict_rounding_up_price,omitempty" `

    /*
        促销信息-凑单价说明，描述凑单价的实现说明。如 “可凑单”或“需买X件”     */
    PredictRoundingUpPriceDesc  *string `json:"predict_rounding_up_price_desc,omitempty" `

    /*
        更多活动优惠     */
    MorePromotionList  *[]TaobaoTbkDgMaterialRecommendMorePromotionMapData `json:"more_promotion_list,omitempty" `

    /*
        预估凑单优惠路径     */
    PredictRoundingUpPathList  *[]TaobaoTbkDgMaterialRecommendPredictRoundingUpPathMapData `json:"predict_rounding_up_path_list,omitempty" `

}

func (s *TaobaoTbkDgMaterialRecommendPromotionInfoMapData) SetReservePrice(v string) *TaobaoTbkDgMaterialRecommendPromotionInfoMapData {
    s.ReservePrice = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendPromotionInfoMapData) SetZkFinalPrice(v string) *TaobaoTbkDgMaterialRecommendPromotionInfoMapData {
    s.ZkFinalPrice = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendPromotionInfoMapData) SetFinalPromotionPrice(v string) *TaobaoTbkDgMaterialRecommendPromotionInfoMapData {
    s.FinalPromotionPrice = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendPromotionInfoMapData) SetFinalPromotionPathList(v []TaobaoTbkDgMaterialRecommendFinalPromotionPathMapData) *TaobaoTbkDgMaterialRecommendPromotionInfoMapData {
    s.FinalPromotionPathList = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendPromotionInfoMapData) SetFutureActivityPromotionPrice(v string) *TaobaoTbkDgMaterialRecommendPromotionInfoMapData {
    s.FutureActivityPromotionPrice = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendPromotionInfoMapData) SetFutureActivityPromotionPathList(v []TaobaoTbkDgMaterialRecommendFutureActivityPromotionPathMapData) *TaobaoTbkDgMaterialRecommendPromotionInfoMapData {
    s.FutureActivityPromotionPathList = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendPromotionInfoMapData) SetPromotionTagList(v []TaobaoTbkDgMaterialRecommendPromotionTagMapData) *TaobaoTbkDgMaterialRecommendPromotionInfoMapData {
    s.PromotionTagList = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendPromotionInfoMapData) SetPredictRoundingUpPrice(v string) *TaobaoTbkDgMaterialRecommendPromotionInfoMapData {
    s.PredictRoundingUpPrice = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendPromotionInfoMapData) SetPredictRoundingUpPriceDesc(v string) *TaobaoTbkDgMaterialRecommendPromotionInfoMapData {
    s.PredictRoundingUpPriceDesc = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendPromotionInfoMapData) SetMorePromotionList(v []TaobaoTbkDgMaterialRecommendMorePromotionMapData) *TaobaoTbkDgMaterialRecommendPromotionInfoMapData {
    s.MorePromotionList = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendPromotionInfoMapData) SetPredictRoundingUpPathList(v []TaobaoTbkDgMaterialRecommendPredictRoundingUpPathMapData) *TaobaoTbkDgMaterialRecommendPromotionInfoMapData {
    s.PredictRoundingUpPathList = &v
    return s
}
