package domain


type TaobaoTbkDgMaterialRecommendPredictRoundingUpPathMapData struct {
    /*
        优惠名称，如“商品券”、“跨店满减”、“单品直降”等     */
    PromotionTitle  *string `json:"promotion_title,omitempty" `

    /*
        优惠利益点文案，如“2件5折”、“每200减20”等     */
    PromotionDesc  *string `json:"promotion_desc,omitempty" `

}

func (s *TaobaoTbkDgMaterialRecommendPredictRoundingUpPathMapData) SetPromotionTitle(v string) *TaobaoTbkDgMaterialRecommendPredictRoundingUpPathMapData {
    s.PromotionTitle = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendPredictRoundingUpPathMapData) SetPromotionDesc(v string) *TaobaoTbkDgMaterialRecommendPredictRoundingUpPathMapData {
    s.PromotionDesc = &v
    return s
}
