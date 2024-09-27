package domain


type TaobaoTbkDgMaterialOptionalUpgradePromotionTagMapData struct {
    /*
        标签名称，如“88VIP”、“花呗免息”、“猫超买返”、“是否包邮”     */
    TagName  *string `json:"tag_name,omitempty" `

}

func (s *TaobaoTbkDgMaterialOptionalUpgradePromotionTagMapData) SetTagName(v string) *TaobaoTbkDgMaterialOptionalUpgradePromotionTagMapData {
    s.TagName = &v
    return s
}
