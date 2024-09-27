package domain


type TaobaoTbkItemInfoUpgradeGetTbkItemDetail struct {
    /*
        商品ID     */
    ItemId  *string `json:"item_id,omitempty" `

    /*
        淘客推广信息     */
    PublishInfo  *TaobaoTbkItemInfoUpgradeGetPublishInfo `json:"publish_info,omitempty" `

    /*
        价格促销信息     */
    PricePromotionInfo  *TaobaoTbkItemInfoUpgradeGetPromotionInfoMapData `json:"price_promotion_info,omitempty" `

    /*
        输入的（新）商品ID     */
    InputItemIid  *string `json:"input_item_iid,omitempty" `

    /*
        商品基本信息     */
    ItemBasicInfo  *TaobaoTbkItemInfoUpgradeGetBasicMapData `json:"item_basic_info,omitempty" `

    /*
        预售信息     */
    PresaleInfo  *TaobaoTbkItemInfoUpgradeGetPresaleInfo `json:"presale_info,omitempty" `

}

func (s *TaobaoTbkItemInfoUpgradeGetTbkItemDetail) SetItemId(v string) *TaobaoTbkItemInfoUpgradeGetTbkItemDetail {
    s.ItemId = &v
    return s
}
func (s *TaobaoTbkItemInfoUpgradeGetTbkItemDetail) SetPublishInfo(v TaobaoTbkItemInfoUpgradeGetPublishInfo) *TaobaoTbkItemInfoUpgradeGetTbkItemDetail {
    s.PublishInfo = &v
    return s
}
func (s *TaobaoTbkItemInfoUpgradeGetTbkItemDetail) SetPricePromotionInfo(v TaobaoTbkItemInfoUpgradeGetPromotionInfoMapData) *TaobaoTbkItemInfoUpgradeGetTbkItemDetail {
    s.PricePromotionInfo = &v
    return s
}
func (s *TaobaoTbkItemInfoUpgradeGetTbkItemDetail) SetInputItemIid(v string) *TaobaoTbkItemInfoUpgradeGetTbkItemDetail {
    s.InputItemIid = &v
    return s
}
func (s *TaobaoTbkItemInfoUpgradeGetTbkItemDetail) SetItemBasicInfo(v TaobaoTbkItemInfoUpgradeGetBasicMapData) *TaobaoTbkItemInfoUpgradeGetTbkItemDetail {
    s.ItemBasicInfo = &v
    return s
}
func (s *TaobaoTbkItemInfoUpgradeGetTbkItemDetail) SetPresaleInfo(v TaobaoTbkItemInfoUpgradeGetPresaleInfo) *TaobaoTbkItemInfoUpgradeGetTbkItemDetail {
    s.PresaleInfo = &v
    return s
}
