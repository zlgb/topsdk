package domain


type TaobaoTbkDgGeneralLinkConvertTbkLinkDTO struct {
    /*
        转链结果     */
    MaterialUrlList  *[]TaobaoTbkDgGeneralLinkConvertMaterialUrlList `json:"material_url_list,omitempty" `

    /*
        店铺转链结果     */
    ShopUrlList  *[]TaobaoTbkDgGeneralLinkConvertShopUrlList `json:"shop_url_list,omitempty" `

    /*
        活动转链信息     */
    EventUrlList  *[]TaobaoTbkDgGeneralLinkConvertEventUrlList `json:"event_url_list,omitempty" `

    /*
        单品转链信息     */
    ItemUrlList  *[]TaobaoTbkDgGeneralLinkConvertItemUrlList `json:"item_url_list,omitempty" `

}

func (s *TaobaoTbkDgGeneralLinkConvertTbkLinkDTO) SetMaterialUrlList(v []TaobaoTbkDgGeneralLinkConvertMaterialUrlList) *TaobaoTbkDgGeneralLinkConvertTbkLinkDTO {
    s.MaterialUrlList = &v
    return s
}
func (s *TaobaoTbkDgGeneralLinkConvertTbkLinkDTO) SetShopUrlList(v []TaobaoTbkDgGeneralLinkConvertShopUrlList) *TaobaoTbkDgGeneralLinkConvertTbkLinkDTO {
    s.ShopUrlList = &v
    return s
}
func (s *TaobaoTbkDgGeneralLinkConvertTbkLinkDTO) SetEventUrlList(v []TaobaoTbkDgGeneralLinkConvertEventUrlList) *TaobaoTbkDgGeneralLinkConvertTbkLinkDTO {
    s.EventUrlList = &v
    return s
}
func (s *TaobaoTbkDgGeneralLinkConvertTbkLinkDTO) SetItemUrlList(v []TaobaoTbkDgGeneralLinkConvertItemUrlList) *TaobaoTbkDgGeneralLinkConvertTbkLinkDTO {
    s.ItemUrlList = &v
    return s
}
