package domain


type TaobaoTbkItemInfoUpgradeGetPromotionInfoMapData struct {
    /*
        商品一口价格     */
    ReservePrice  *string `json:"reserve_price,omitempty" `

    /*
        折扣价（元） 若属于预售商品，付定金时间内，折扣价=预售价     */
    ZkFinalPrice  *string `json:"zk_final_price,omitempty" `

    /*
        促销信息-预估到手价(元)。若属于预售商品，付定金时间内，预估到手价=预售尾款预估到手价     */
    FinalPromotionPrice  *string `json:"final_promotion_price,omitempty" `

    /*
        优惠促销列表     */
    FinalPromotionPathList  *[]TaobaoTbkItemInfoUpgradeGetFinalPromotionPathMapData `json:"final_promotion_path_list,omitempty" `

    /*
        标签信息列表     */
    PromotionTagList  *[]TaobaoTbkItemInfoUpgradeGetPromotionTagMapData `json:"promotion_tag_list,omitempty" `

    /*
        促销信息-预估凑单价（元）。预估凑单叠加优惠后的商品单价     */
    PredictRoundingUpPrice  *string `json:"predict_rounding_up_price,omitempty" `

    /*
        促销信息-凑单价说明，描述凑单价的实现说明。如 “可凑单”或“需买X件”     */
    PredictRoundingUpPriceDesc  *string `json:"predict_rounding_up_price_desc,omitempty" `

    /*
        更多活动优惠     */
    MorePromotionList  *[]TaobaoTbkItemInfoUpgradeGetMorePromotionMapData `json:"more_promotion_list,omitempty" `

    /*
        货品展示标识列表     */
    ActivityTagList  *[]TaobaoTbkItemInfoUpgradeGetActivityTagMapData `json:"activity_tag_list,omitempty" `

}

func (s *TaobaoTbkItemInfoUpgradeGetPromotionInfoMapData) SetReservePrice(v string) *TaobaoTbkItemInfoUpgradeGetPromotionInfoMapData {
    s.ReservePrice = &v
    return s
}
func (s *TaobaoTbkItemInfoUpgradeGetPromotionInfoMapData) SetZkFinalPrice(v string) *TaobaoTbkItemInfoUpgradeGetPromotionInfoMapData {
    s.ZkFinalPrice = &v
    return s
}
func (s *TaobaoTbkItemInfoUpgradeGetPromotionInfoMapData) SetFinalPromotionPrice(v string) *TaobaoTbkItemInfoUpgradeGetPromotionInfoMapData {
    s.FinalPromotionPrice = &v
    return s
}
func (s *TaobaoTbkItemInfoUpgradeGetPromotionInfoMapData) SetFinalPromotionPathList(v []TaobaoTbkItemInfoUpgradeGetFinalPromotionPathMapData) *TaobaoTbkItemInfoUpgradeGetPromotionInfoMapData {
    s.FinalPromotionPathList = &v
    return s
}
func (s *TaobaoTbkItemInfoUpgradeGetPromotionInfoMapData) SetPromotionTagList(v []TaobaoTbkItemInfoUpgradeGetPromotionTagMapData) *TaobaoTbkItemInfoUpgradeGetPromotionInfoMapData {
    s.PromotionTagList = &v
    return s
}
func (s *TaobaoTbkItemInfoUpgradeGetPromotionInfoMapData) SetPredictRoundingUpPrice(v string) *TaobaoTbkItemInfoUpgradeGetPromotionInfoMapData {
    s.PredictRoundingUpPrice = &v
    return s
}
func (s *TaobaoTbkItemInfoUpgradeGetPromotionInfoMapData) SetPredictRoundingUpPriceDesc(v string) *TaobaoTbkItemInfoUpgradeGetPromotionInfoMapData {
    s.PredictRoundingUpPriceDesc = &v
    return s
}
func (s *TaobaoTbkItemInfoUpgradeGetPromotionInfoMapData) SetMorePromotionList(v []TaobaoTbkItemInfoUpgradeGetMorePromotionMapData) *TaobaoTbkItemInfoUpgradeGetPromotionInfoMapData {
    s.MorePromotionList = &v
    return s
}
func (s *TaobaoTbkItemInfoUpgradeGetPromotionInfoMapData) SetActivityTagList(v []TaobaoTbkItemInfoUpgradeGetActivityTagMapData) *TaobaoTbkItemInfoUpgradeGetPromotionInfoMapData {
    s.ActivityTagList = &v
    return s
}
