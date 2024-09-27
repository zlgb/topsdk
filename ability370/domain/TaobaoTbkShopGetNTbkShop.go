package domain


type TaobaoTbkShopGetNTbkShop struct {
    /*
        店铺地址     */
    ShopUrl  *string `json:"shop_url,omitempty" `

    /*
        卖家昵称     */
    SellerNick  *string `json:"seller_nick,omitempty" `

    /*
        店铺类型，B：天猫，C：淘宝     */
    ShopType  *string `json:"shop_type,omitempty" `

    /*
        店铺名称     */
    ShopTitle  *string `json:"shop_title,omitempty" `

    /*
        卖家ID     */
    UserId  *int64 `json:"user_id,omitempty" `

    /*
        店标图片     */
    PictUrl  *string `json:"pict_url,omitempty" `

}

func (s *TaobaoTbkShopGetNTbkShop) SetShopUrl(v string) *TaobaoTbkShopGetNTbkShop {
    s.ShopUrl = &v
    return s
}
func (s *TaobaoTbkShopGetNTbkShop) SetSellerNick(v string) *TaobaoTbkShopGetNTbkShop {
    s.SellerNick = &v
    return s
}
func (s *TaobaoTbkShopGetNTbkShop) SetShopType(v string) *TaobaoTbkShopGetNTbkShop {
    s.ShopType = &v
    return s
}
func (s *TaobaoTbkShopGetNTbkShop) SetShopTitle(v string) *TaobaoTbkShopGetNTbkShop {
    s.ShopTitle = &v
    return s
}
func (s *TaobaoTbkShopGetNTbkShop) SetUserId(v int64) *TaobaoTbkShopGetNTbkShop {
    s.UserId = &v
    return s
}
func (s *TaobaoTbkShopGetNTbkShop) SetPictUrl(v string) *TaobaoTbkShopGetNTbkShop {
    s.PictUrl = &v
    return s
}
