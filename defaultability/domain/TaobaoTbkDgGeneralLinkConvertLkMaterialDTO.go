package domain


type TaobaoTbkDgGeneralLinkConvertLkMaterialDTO struct {
    /*
        物料链接，可以为url或淘口令     */
    MaterialUrl  *string `json:"material_url,omitempty" `

    /*
        是否指定券，1-指定 0-不指定 默认为0（邀约制权限，仅面向KA淘宝客）     */
    IsTargetCoupon  *int64 `json:"is_target_coupon,omitempty" `

    /*
        优惠券id（邀约制权限，仅面向KA淘宝客）     */
    CouponId  *string `json:"coupon_id,omitempty" `

}

func (s *TaobaoTbkDgGeneralLinkConvertLkMaterialDTO) SetMaterialUrl(v string) *TaobaoTbkDgGeneralLinkConvertLkMaterialDTO {
    s.MaterialUrl = &v
    return s
}
func (s *TaobaoTbkDgGeneralLinkConvertLkMaterialDTO) SetIsTargetCoupon(v int64) *TaobaoTbkDgGeneralLinkConvertLkMaterialDTO {
    s.IsTargetCoupon = &v
    return s
}
func (s *TaobaoTbkDgGeneralLinkConvertLkMaterialDTO) SetCouponId(v string) *TaobaoTbkDgGeneralLinkConvertLkMaterialDTO {
    s.CouponId = &v
    return s
}
