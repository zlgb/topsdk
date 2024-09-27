package domain


type TaobaoTbkDgMaterialOptionalUpgradeMgcInfo struct {
    /*
        价格     */
    Price  *string `json:"price,omitempty" `

    /*
        价格描述     */
    PriceDesc  *string `json:"price_desc,omitempty" `

    /*
        文案     */
    PromotionSummary  *string `json:"promotion_summary,omitempty" `

    /*
        发布时间，13位毫秒时间戳     */
    PublishTime  *string `json:"publish_time,omitempty" `

    /*
        生效时间，实时线报为0，未来线报为13位毫秒时间戳     */
    ValidTime  *string `json:"valid_time,omitempty" `

}

func (s *TaobaoTbkDgMaterialOptionalUpgradeMgcInfo) SetPrice(v string) *TaobaoTbkDgMaterialOptionalUpgradeMgcInfo {
    s.Price = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradeMgcInfo) SetPriceDesc(v string) *TaobaoTbkDgMaterialOptionalUpgradeMgcInfo {
    s.PriceDesc = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradeMgcInfo) SetPromotionSummary(v string) *TaobaoTbkDgMaterialOptionalUpgradeMgcInfo {
    s.PromotionSummary = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradeMgcInfo) SetPublishTime(v string) *TaobaoTbkDgMaterialOptionalUpgradeMgcInfo {
    s.PublishTime = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradeMgcInfo) SetValidTime(v string) *TaobaoTbkDgMaterialOptionalUpgradeMgcInfo {
    s.ValidTime = &v
    return s
}
