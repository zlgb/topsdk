package domain


type TaobaoTbkDgMaterialRecommendPromotionTagMapData struct {
    /*
        标签名称，如“88VIP”、“花呗免息”、“猫超买返”、“是否包邮”     */
    TagName  *string `json:"tag_name,omitempty" `

}

func (s *TaobaoTbkDgMaterialRecommendPromotionTagMapData) SetTagName(v string) *TaobaoTbkDgMaterialRecommendPromotionTagMapData {
    s.TagName = &v
    return s
}
