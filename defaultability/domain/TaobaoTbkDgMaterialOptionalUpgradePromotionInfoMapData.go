package domain


type TaobaoTbkDgMaterialOptionalUpgradePromotionInfoMapData struct {
    /*
        促销信息-预估凑单价（元）。预估凑单叠加优惠后的商品单价     */
    PredictRoundingUpPrice  *string `json:"predict_rounding_up_price,omitempty" `

    /*
        促销信息-凑单价说明，描述凑单价的实现说明。如 “可凑单”或“需买X件”     */
    PredictRoundingUpPriceDesc  *string `json:"predict_rounding_up_price_desc,omitempty" `

    /*
        更多活动优惠     */
    MorePromotionList  *[]TaobaoTbkDgMaterialOptionalUpgradeMorePromotionMapData `json:"more_promotion_list,omitempty" `

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
    FinalPromotionPathList  *[]TaobaoTbkDgMaterialOptionalUpgradeFinalPromotionPathMapData `json:"final_promotion_path_list,omitempty" `

    /*
        预热预估到手价（元）     */
    FutureActivityPromotionPrice  *string `json:"future_activity_promotion_price,omitempty" `

    /*
        预热到手价优惠路径列表     */
    FutureActivityPromotionPathList  *[]TaobaoTbkDgMaterialOptionalUpgradeFutureActivityPromotionPathMapData `json:"future_activity_promotion_path_list,omitempty" `

    /*
        标签信息列表     */
    PromotionTagList  *[]TaobaoTbkDgMaterialOptionalUpgradePromotionTagMapData `json:"promotion_tag_list,omitempty" `

}

func (s *TaobaoTbkDgMaterialOptionalUpgradePromotionInfoMapData) SetPredictRoundingUpPrice(v string) *TaobaoTbkDgMaterialOptionalUpgradePromotionInfoMapData {
    s.PredictRoundingUpPrice = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradePromotionInfoMapData) SetPredictRoundingUpPriceDesc(v string) *TaobaoTbkDgMaterialOptionalUpgradePromotionInfoMapData {
    s.PredictRoundingUpPriceDesc = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradePromotionInfoMapData) SetMorePromotionList(v []TaobaoTbkDgMaterialOptionalUpgradeMorePromotionMapData) *TaobaoTbkDgMaterialOptionalUpgradePromotionInfoMapData {
    s.MorePromotionList = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradePromotionInfoMapData) SetReservePrice(v string) *TaobaoTbkDgMaterialOptionalUpgradePromotionInfoMapData {
    s.ReservePrice = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradePromotionInfoMapData) SetZkFinalPrice(v string) *TaobaoTbkDgMaterialOptionalUpgradePromotionInfoMapData {
    s.ZkFinalPrice = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradePromotionInfoMapData) SetFinalPromotionPrice(v string) *TaobaoTbkDgMaterialOptionalUpgradePromotionInfoMapData {
    s.FinalPromotionPrice = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradePromotionInfoMapData) SetFinalPromotionPathList(v []TaobaoTbkDgMaterialOptionalUpgradeFinalPromotionPathMapData) *TaobaoTbkDgMaterialOptionalUpgradePromotionInfoMapData {
    s.FinalPromotionPathList = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradePromotionInfoMapData) SetFutureActivityPromotionPrice(v string) *TaobaoTbkDgMaterialOptionalUpgradePromotionInfoMapData {
    s.FutureActivityPromotionPrice = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradePromotionInfoMapData) SetFutureActivityPromotionPathList(v []TaobaoTbkDgMaterialOptionalUpgradeFutureActivityPromotionPathMapData) *TaobaoTbkDgMaterialOptionalUpgradePromotionInfoMapData {
    s.FutureActivityPromotionPathList = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradePromotionInfoMapData) SetPromotionTagList(v []TaobaoTbkDgMaterialOptionalUpgradePromotionTagMapData) *TaobaoTbkDgMaterialOptionalUpgradePromotionInfoMapData {
    s.PromotionTagList = &v
    return s
}
