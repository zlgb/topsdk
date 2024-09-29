package domain


type TaobaoTbkDgGeneralLinkConvertLinkInfoDTO struct {
    /*
        二合一长链接     */
    CouponLongUrl  *string `json:"coupon_long_url,omitempty" `

    /*
        1—单品 2—店铺 3—会场 4-承接开放 5-优惠券 6-直播间 7-淘积木页     */
    MaterialType  *int64 `json:"material_type,omitempty" `

    /*
        淘口令原始链接，入参物料不为淘口令时，此字段返回null     */
    TpwdOriginUrl  *string `json:"tpwd_origin_url,omitempty" `

    /*
        物料ID，需根据material_type判断物料类型， 可能为商品ID、卖家ID、会场ID     */
    MaterialId  *string `json:"material_id,omitempty" `

    /*
        CPS链接长链     */
    CpsLongUrl  *string `json:"cps_long_url,omitempty" `

    /*
        CPS链接的短口令     */
    CpsShortTpwd  *string `json:"cps_short_tpwd,omitempty" `

    /*
        1-联盟口令，2-主站口令。入参物料不为淘口令时，此字段返回null     */
    TkBizType  *int64 `json:"tk_biz_type,omitempty" `

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

    /*
        淘积木页面链接，定向开放     */
    TaoBrickUrl  *string `json:"tao_brick_url,omitempty" `

}

func (s *TaobaoTbkDgGeneralLinkConvertLinkInfoDTO) SetCouponLongUrl(v string) *TaobaoTbkDgGeneralLinkConvertLinkInfoDTO {
    s.CouponLongUrl = &v
    return s
}
func (s *TaobaoTbkDgGeneralLinkConvertLinkInfoDTO) SetMaterialType(v int64) *TaobaoTbkDgGeneralLinkConvertLinkInfoDTO {
    s.MaterialType = &v
    return s
}
func (s *TaobaoTbkDgGeneralLinkConvertLinkInfoDTO) SetTpwdOriginUrl(v string) *TaobaoTbkDgGeneralLinkConvertLinkInfoDTO {
    s.TpwdOriginUrl = &v
    return s
}
func (s *TaobaoTbkDgGeneralLinkConvertLinkInfoDTO) SetMaterialId(v string) *TaobaoTbkDgGeneralLinkConvertLinkInfoDTO {
    s.MaterialId = &v
    return s
}
func (s *TaobaoTbkDgGeneralLinkConvertLinkInfoDTO) SetCpsLongUrl(v string) *TaobaoTbkDgGeneralLinkConvertLinkInfoDTO {
    s.CpsLongUrl = &v
    return s
}
func (s *TaobaoTbkDgGeneralLinkConvertLinkInfoDTO) SetCpsShortTpwd(v string) *TaobaoTbkDgGeneralLinkConvertLinkInfoDTO {
    s.CpsShortTpwd = &v
    return s
}
func (s *TaobaoTbkDgGeneralLinkConvertLinkInfoDTO) SetTkBizType(v int64) *TaobaoTbkDgGeneralLinkConvertLinkInfoDTO {
    s.TkBizType = &v
    return s
}
func (s *TaobaoTbkDgGeneralLinkConvertLinkInfoDTO) SetCouponShortTpwd(v string) *TaobaoTbkDgGeneralLinkConvertLinkInfoDTO {
    s.CouponShortTpwd = &v
    return s
}
func (s *TaobaoTbkDgGeneralLinkConvertLinkInfoDTO) SetCpsShortUrl(v string) *TaobaoTbkDgGeneralLinkConvertLinkInfoDTO {
    s.CpsShortUrl = &v
    return s
}
func (s *TaobaoTbkDgGeneralLinkConvertLinkInfoDTO) SetCouponShortUrl(v string) *TaobaoTbkDgGeneralLinkConvertLinkInfoDTO {
    s.CouponShortUrl = &v
    return s
}
func (s *TaobaoTbkDgGeneralLinkConvertLinkInfoDTO) SetCouponFullTpwd(v string) *TaobaoTbkDgGeneralLinkConvertLinkInfoDTO {
    s.CouponFullTpwd = &v
    return s
}
func (s *TaobaoTbkDgGeneralLinkConvertLinkInfoDTO) SetCpsFullTpwd(v string) *TaobaoTbkDgGeneralLinkConvertLinkInfoDTO {
    s.CpsFullTpwd = &v
    return s
}
func (s *TaobaoTbkDgGeneralLinkConvertLinkInfoDTO) SetShareUnitUrl(v string) *TaobaoTbkDgGeneralLinkConvertLinkInfoDTO {
    s.ShareUnitUrl = &v
    return s
}
func (s *TaobaoTbkDgGeneralLinkConvertLinkInfoDTO) SetTaoBrickUrl(v string) *TaobaoTbkDgGeneralLinkConvertLinkInfoDTO {
    s.TaoBrickUrl = &v
    return s
}
