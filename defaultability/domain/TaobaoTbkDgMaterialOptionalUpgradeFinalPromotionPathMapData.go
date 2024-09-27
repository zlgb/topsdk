package domain


type TaobaoTbkDgMaterialOptionalUpgradeFinalPromotionPathMapData struct {
    /*
        优惠名称，如“商品券”、“跨店满减”、“单品直降”等  	     */
    PromotionTitle  *string `json:"promotion_title,omitempty" `

    /*
        优惠利益点文案，如“1件7.92折”、“每200减20”等     */
    PromotionDesc  *string `json:"promotion_desc,omitempty" `

    /*
        实际优惠金额（元）     */
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

func (s *TaobaoTbkDgMaterialOptionalUpgradeFinalPromotionPathMapData) SetPromotionTitle(v string) *TaobaoTbkDgMaterialOptionalUpgradeFinalPromotionPathMapData {
    s.PromotionTitle = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradeFinalPromotionPathMapData) SetPromotionDesc(v string) *TaobaoTbkDgMaterialOptionalUpgradeFinalPromotionPathMapData {
    s.PromotionDesc = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradeFinalPromotionPathMapData) SetPromotionFee(v string) *TaobaoTbkDgMaterialOptionalUpgradeFinalPromotionPathMapData {
    s.PromotionFee = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradeFinalPromotionPathMapData) SetPromotionStartTime(v string) *TaobaoTbkDgMaterialOptionalUpgradeFinalPromotionPathMapData {
    s.PromotionStartTime = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradeFinalPromotionPathMapData) SetPromotionEndTime(v string) *TaobaoTbkDgMaterialOptionalUpgradeFinalPromotionPathMapData {
    s.PromotionEndTime = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradeFinalPromotionPathMapData) SetPromotionId(v string) *TaobaoTbkDgMaterialOptionalUpgradeFinalPromotionPathMapData {
    s.PromotionId = &v
    return s
}
