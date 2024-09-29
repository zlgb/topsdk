package domain


type TaobaoTbkDgGeneralLinkConvertLkItemDTO struct {
    /*
        商品ID     */
    ItemId  *string `json:"item_id,omitempty" `

    /*
        入参商品id下的skuid，传入时会透传至转链结果url中     */
    SkuId  *int64 `json:"sku_id,omitempty" `

    /*
        是否指定券，1-指定 0-不指定 默认为0（邀约制权限，仅面向KA淘宝客）     */
    IsTargetCoupon  *int64 `json:"is_target_coupon,omitempty" `

    /*
        优惠券id（邀约制权限，仅面向KA淘宝客）     */
    CouponId  *string `json:"coupon_id,omitempty" `

    /*
        淘宝客外部用户标记，如自身系统账户ID；微信ID等     */
    ExternalId  *string `json:"external_id,omitempty" `

    /*
        1表示商品转通用计划链接，其他值或不传表示转最优佣金率（含营销计划）链接     */
    Dx  *string `json:"dx,omitempty" `

}

func (s *TaobaoTbkDgGeneralLinkConvertLkItemDTO) SetItemId(v string) *TaobaoTbkDgGeneralLinkConvertLkItemDTO {
    s.ItemId = &v
    return s
}
func (s *TaobaoTbkDgGeneralLinkConvertLkItemDTO) SetSkuId(v int64) *TaobaoTbkDgGeneralLinkConvertLkItemDTO {
    s.SkuId = &v
    return s
}
func (s *TaobaoTbkDgGeneralLinkConvertLkItemDTO) SetIsTargetCoupon(v int64) *TaobaoTbkDgGeneralLinkConvertLkItemDTO {
    s.IsTargetCoupon = &v
    return s
}
func (s *TaobaoTbkDgGeneralLinkConvertLkItemDTO) SetCouponId(v string) *TaobaoTbkDgGeneralLinkConvertLkItemDTO {
    s.CouponId = &v
    return s
}
func (s *TaobaoTbkDgGeneralLinkConvertLkItemDTO) SetExternalId(v string) *TaobaoTbkDgGeneralLinkConvertLkItemDTO {
    s.ExternalId = &v
    return s
}
func (s *TaobaoTbkDgGeneralLinkConvertLkItemDTO) SetDx(v string) *TaobaoTbkDgGeneralLinkConvertLkItemDTO {
    s.Dx = &v
    return s
}
