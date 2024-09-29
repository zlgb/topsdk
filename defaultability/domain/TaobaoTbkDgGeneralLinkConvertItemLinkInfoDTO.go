package domain


type TaobaoTbkDgGeneralLinkConvertItemLinkInfoDTO struct {
    /*
        二合一长链接     */
    CouponLongUrl  *string `json:"coupon_long_url,omitempty" `

    /*
        1—单品 2—店铺 3—会场     */
    MaterialType  *int64 `json:"material_type,omitempty" `

    /*
        商品ID     */
    ItemId  *string `json:"item_id,omitempty" `

    /*
        CPS链接长链     */
    CpsLongUrl  *string `json:"cps_long_url,omitempty" `

    /*
        CPS链接的短口令     */
    CpsShortTpwd  *string `json:"cps_short_tpwd,omitempty" `

    /*
        二合一链接的短口令     */
    CouponShortTpwd  *string `json:"coupon_short_tpwd,omitempty" `

    /*
        CPS链接短链     */
    CpsShortUrl  *string `json:"cps_short_url,omitempty" `

    /*
        二合一链接长链     */
    CouponShortUrl  *string `json:"coupon_short_url,omitempty" `

    /*
        二合一链接的长口令     */
    CouponFullTpwd  *string `json:"coupon_full_tpwd,omitempty" `

    /*
        CPS链接的长口令     */
    CpsFullTpwd  *string `json:"cps_full_tpwd,omitempty" `

    /*
        种草页链接，定向开放     */
    ShareUnitUrl  *string `json:"share_unit_url,omitempty" `

}

func (s *TaobaoTbkDgGeneralLinkConvertItemLinkInfoDTO) SetCouponLongUrl(v string) *TaobaoTbkDgGeneralLinkConvertItemLinkInfoDTO {
    s.CouponLongUrl = &v
    return s
}
func (s *TaobaoTbkDgGeneralLinkConvertItemLinkInfoDTO) SetMaterialType(v int64) *TaobaoTbkDgGeneralLinkConvertItemLinkInfoDTO {
    s.MaterialType = &v
    return s
}
func (s *TaobaoTbkDgGeneralLinkConvertItemLinkInfoDTO) SetItemId(v string) *TaobaoTbkDgGeneralLinkConvertItemLinkInfoDTO {
    s.ItemId = &v
    return s
}
func (s *TaobaoTbkDgGeneralLinkConvertItemLinkInfoDTO) SetCpsLongUrl(v string) *TaobaoTbkDgGeneralLinkConvertItemLinkInfoDTO {
    s.CpsLongUrl = &v
    return s
}
func (s *TaobaoTbkDgGeneralLinkConvertItemLinkInfoDTO) SetCpsShortTpwd(v string) *TaobaoTbkDgGeneralLinkConvertItemLinkInfoDTO {
    s.CpsShortTpwd = &v
    return s
}
func (s *TaobaoTbkDgGeneralLinkConvertItemLinkInfoDTO) SetCouponShortTpwd(v string) *TaobaoTbkDgGeneralLinkConvertItemLinkInfoDTO {
    s.CouponShortTpwd = &v
    return s
}
func (s *TaobaoTbkDgGeneralLinkConvertItemLinkInfoDTO) SetCpsShortUrl(v string) *TaobaoTbkDgGeneralLinkConvertItemLinkInfoDTO {
    s.CpsShortUrl = &v
    return s
}
func (s *TaobaoTbkDgGeneralLinkConvertItemLinkInfoDTO) SetCouponShortUrl(v string) *TaobaoTbkDgGeneralLinkConvertItemLinkInfoDTO {
    s.CouponShortUrl = &v
    return s
}
func (s *TaobaoTbkDgGeneralLinkConvertItemLinkInfoDTO) SetCouponFullTpwd(v string) *TaobaoTbkDgGeneralLinkConvertItemLinkInfoDTO {
    s.CouponFullTpwd = &v
    return s
}
func (s *TaobaoTbkDgGeneralLinkConvertItemLinkInfoDTO) SetCpsFullTpwd(v string) *TaobaoTbkDgGeneralLinkConvertItemLinkInfoDTO {
    s.CpsFullTpwd = &v
    return s
}
func (s *TaobaoTbkDgGeneralLinkConvertItemLinkInfoDTO) SetShareUnitUrl(v string) *TaobaoTbkDgGeneralLinkConvertItemLinkInfoDTO {
    s.ShareUnitUrl = &v
    return s
}
