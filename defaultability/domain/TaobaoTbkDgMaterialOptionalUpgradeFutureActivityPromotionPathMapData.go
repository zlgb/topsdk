package domain


type TaobaoTbkDgMaterialOptionalUpgradeFutureActivityPromotionPathMapData struct {
    /*
        优惠名称，如“商品券”、“跨店满减”、“单品直降”等  	     */
    PromotionTitle  *string `json:"promotion_title,omitempty" `

    /*
        预热优惠利益点文案，如“1件7.92折”、“每200减20”等     */
    PromotionDesc  *string `json:"promotion_desc,omitempty" `

    /*
        预热实际优惠金额（元）     */
    PromotionFee  *string `json:"promotion_fee,omitempty" `

    /*
        优惠开始时间     */
    PromotionStartTime  *string `json:"promotion_start_time,omitempty" `

    /*
        优惠结束时间     */
    PromotionEndTime  *string `json:"promotion_end_time,omitempty" `

}

func (s *TaobaoTbkDgMaterialOptionalUpgradeFutureActivityPromotionPathMapData) SetPromotionTitle(v string) *TaobaoTbkDgMaterialOptionalUpgradeFutureActivityPromotionPathMapData {
    s.PromotionTitle = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradeFutureActivityPromotionPathMapData) SetPromotionDesc(v string) *TaobaoTbkDgMaterialOptionalUpgradeFutureActivityPromotionPathMapData {
    s.PromotionDesc = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradeFutureActivityPromotionPathMapData) SetPromotionFee(v string) *TaobaoTbkDgMaterialOptionalUpgradeFutureActivityPromotionPathMapData {
    s.PromotionFee = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradeFutureActivityPromotionPathMapData) SetPromotionStartTime(v string) *TaobaoTbkDgMaterialOptionalUpgradeFutureActivityPromotionPathMapData {
    s.PromotionStartTime = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradeFutureActivityPromotionPathMapData) SetPromotionEndTime(v string) *TaobaoTbkDgMaterialOptionalUpgradeFutureActivityPromotionPathMapData {
    s.PromotionEndTime = &v
    return s
}
