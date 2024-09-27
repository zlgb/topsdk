package domain


type TaobaoTbkDgMaterialRecommendFinalPromotionPathMapData struct {
    /*
        优惠名称，如“商品券”、“跨店满减”、“单品直降”等     */
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

func (s *TaobaoTbkDgMaterialRecommendFinalPromotionPathMapData) SetPromotionTitle(v string) *TaobaoTbkDgMaterialRecommendFinalPromotionPathMapData {
    s.PromotionTitle = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendFinalPromotionPathMapData) SetPromotionDesc(v string) *TaobaoTbkDgMaterialRecommendFinalPromotionPathMapData {
    s.PromotionDesc = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendFinalPromotionPathMapData) SetPromotionFee(v string) *TaobaoTbkDgMaterialRecommendFinalPromotionPathMapData {
    s.PromotionFee = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendFinalPromotionPathMapData) SetPromotionStartTime(v string) *TaobaoTbkDgMaterialRecommendFinalPromotionPathMapData {
    s.PromotionStartTime = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendFinalPromotionPathMapData) SetPromotionEndTime(v string) *TaobaoTbkDgMaterialRecommendFinalPromotionPathMapData {
    s.PromotionEndTime = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendFinalPromotionPathMapData) SetPromotionId(v string) *TaobaoTbkDgMaterialRecommendFinalPromotionPathMapData {
    s.PromotionId = &v
    return s
}
