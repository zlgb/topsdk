package domain


type TaobaoTbkItemInfoUpgradeGetFinalPromotionPathMapData struct {
    /*
        优惠名称，如“商品券”、“跨店满减”、“单品直降”等     */
    PromotionTitle  *string `json:"promotion_title,omitempty" `

    /*
        优惠利益点文案，如“1件7.92折”、“每200减20”等     */
    PromotionDesc  *string `json:"promotion_desc,omitempty" `

    /*
        优惠金额（元）     */
    PromotionFee  *string `json:"promotion_fee,omitempty" `

    /*
        优惠开始时间     */
    PromotionStartTime  *string `json:"promotion_start_time,omitempty" `

    /*
        优惠结束时间     */
    PromotionEndTime  *string `json:"promotion_end_time,omitempty" `

    /*
        优惠ID     */
    PromotionId  *string `json:"promotion_id,omitempty" `

}

func (s *TaobaoTbkItemInfoUpgradeGetFinalPromotionPathMapData) SetPromotionTitle(v string) *TaobaoTbkItemInfoUpgradeGetFinalPromotionPathMapData {
    s.PromotionTitle = &v
    return s
}
func (s *TaobaoTbkItemInfoUpgradeGetFinalPromotionPathMapData) SetPromotionDesc(v string) *TaobaoTbkItemInfoUpgradeGetFinalPromotionPathMapData {
    s.PromotionDesc = &v
    return s
}
func (s *TaobaoTbkItemInfoUpgradeGetFinalPromotionPathMapData) SetPromotionFee(v string) *TaobaoTbkItemInfoUpgradeGetFinalPromotionPathMapData {
    s.PromotionFee = &v
    return s
}
func (s *TaobaoTbkItemInfoUpgradeGetFinalPromotionPathMapData) SetPromotionStartTime(v string) *TaobaoTbkItemInfoUpgradeGetFinalPromotionPathMapData {
    s.PromotionStartTime = &v
    return s
}
func (s *TaobaoTbkItemInfoUpgradeGetFinalPromotionPathMapData) SetPromotionEndTime(v string) *TaobaoTbkItemInfoUpgradeGetFinalPromotionPathMapData {
    s.PromotionEndTime = &v
    return s
}
func (s *TaobaoTbkItemInfoUpgradeGetFinalPromotionPathMapData) SetPromotionId(v string) *TaobaoTbkItemInfoUpgradeGetFinalPromotionPathMapData {
    s.PromotionId = &v
    return s
}
