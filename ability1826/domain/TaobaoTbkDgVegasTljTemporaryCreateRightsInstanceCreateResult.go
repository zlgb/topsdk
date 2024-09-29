package domain


type TaobaoTbkDgVegasTljTemporaryCreateRightsInstanceCreateResult struct {
    /*
        淘礼金Id     */
    RightsId  *string `json:"rights_id,omitempty" `

    /*
        淘礼金领取Url     */
    SendUrl  *string `json:"send_url,omitempty" `

    /*
        投放code【百川商品详情页业务专用】     */
    VegasCode  *string `json:"vegas_code,omitempty" `

    /*
        创建完成后资金账户可用资金，单位元，保留2位小数     */
    AvailableFee  *string `json:"available_fee,omitempty" `

    /*
        媒体针对此商品今日剩余可领取淘礼金数量     */
    ItemTodayNumLeft  *int64 `json:"item_today_num_left,omitempty" `

}

func (s *TaobaoTbkDgVegasTljTemporaryCreateRightsInstanceCreateResult) SetRightsId(v string) *TaobaoTbkDgVegasTljTemporaryCreateRightsInstanceCreateResult {
    s.RightsId = &v
    return s
}
func (s *TaobaoTbkDgVegasTljTemporaryCreateRightsInstanceCreateResult) SetSendUrl(v string) *TaobaoTbkDgVegasTljTemporaryCreateRightsInstanceCreateResult {
    s.SendUrl = &v
    return s
}
func (s *TaobaoTbkDgVegasTljTemporaryCreateRightsInstanceCreateResult) SetVegasCode(v string) *TaobaoTbkDgVegasTljTemporaryCreateRightsInstanceCreateResult {
    s.VegasCode = &v
    return s
}
func (s *TaobaoTbkDgVegasTljTemporaryCreateRightsInstanceCreateResult) SetAvailableFee(v string) *TaobaoTbkDgVegasTljTemporaryCreateRightsInstanceCreateResult {
    s.AvailableFee = &v
    return s
}
func (s *TaobaoTbkDgVegasTljTemporaryCreateRightsInstanceCreateResult) SetItemTodayNumLeft(v int64) *TaobaoTbkDgVegasTljTemporaryCreateRightsInstanceCreateResult {
    s.ItemTodayNumLeft = &v
    return s
}
