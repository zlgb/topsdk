package domain


type TaobaoTbkShopRecommendGetNTbkShop struct {
    /*
        卖家ID     */
    UserId  *int64 `json:"user_id,omitempty" `

    /*
        店铺名称     */
    ShopTitle  *string `json:"shop_title,omitempty" `

    /*
        卖家昵称     */
    SellerNick  *string `json:"seller_nick,omitempty" `

    /*
        店铺类型，B：天猫，C：淘宝     */
    ShopType  *string `json:"shop_type,omitempty" `

    /*
        店铺地址     */
    ShopUrl  *string `json:"shop_url,omitempty" `

    /*
        店标图片     */
    PictUrl  *string `json:"pict_url,omitempty" `

}

func (s *TaobaoTbkShopRecommendGetNTbkShop) SetUserId(v int64) *TaobaoTbkShopRecommendGetNTbkShop {
    s.UserId = &v
    return s
}
func (s *TaobaoTbkShopRecommendGetNTbkShop) SetShopTitle(v string) *TaobaoTbkShopRecommendGetNTbkShop {
    s.ShopTitle = &v
    return s
}
func (s *TaobaoTbkShopRecommendGetNTbkShop) SetSellerNick(v string) *TaobaoTbkShopRecommendGetNTbkShop {
    s.SellerNick = &v
    return s
}
func (s *TaobaoTbkShopRecommendGetNTbkShop) SetShopType(v string) *TaobaoTbkShopRecommendGetNTbkShop {
    s.ShopType = &v
    return s
}
func (s *TaobaoTbkShopRecommendGetNTbkShop) SetShopUrl(v string) *TaobaoTbkShopRecommendGetNTbkShop {
    s.ShopUrl = &v
    return s
}
func (s *TaobaoTbkShopRecommendGetNTbkShop) SetPictUrl(v string) *TaobaoTbkShopRecommendGetNTbkShop {
    s.PictUrl = &v
    return s
}
