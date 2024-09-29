package domain


type TaobaoTbkDgGeneralLinkConvertLkShopDTO struct {
    /*
        店铺ID     */
    ShopId  *string `json:"shop_id,omitempty" `

}

func (s *TaobaoTbkDgGeneralLinkConvertLkShopDTO) SetShopId(v string) *TaobaoTbkDgGeneralLinkConvertLkShopDTO {
    s.ShopId = &v
    return s
}
