package domain


type TaobaoTbkDgMaterialRecommendMorePromotionMapData struct {
    /*
        优惠名称，如“商品券”、“跨店满减”、“单品直降”等     */
    PromotionTitle  *string `json:"promotion_title,omitempty" `

    /*
        优惠利益点文案，如“1件7.92折”、“每200减20”等     */
    PromotionDesc  *string `json:"promotion_desc,omitempty" `

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

func (s *TaobaoTbkDgMaterialRecommendMorePromotionMapData) SetPromotionTitle(v string) *TaobaoTbkDgMaterialRecommendMorePromotionMapData {
    s.PromotionTitle = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendMorePromotionMapData) SetPromotionDesc(v string) *TaobaoTbkDgMaterialRecommendMorePromotionMapData {
    s.PromotionDesc = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendMorePromotionMapData) SetPromotionStartTime(v string) *TaobaoTbkDgMaterialRecommendMorePromotionMapData {
    s.PromotionStartTime = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendMorePromotionMapData) SetPromotionEndTime(v string) *TaobaoTbkDgMaterialRecommendMorePromotionMapData {
    s.PromotionEndTime = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendMorePromotionMapData) SetPromotionId(v string) *TaobaoTbkDgMaterialRecommendMorePromotionMapData {
    s.PromotionId = &v
    return s
}
