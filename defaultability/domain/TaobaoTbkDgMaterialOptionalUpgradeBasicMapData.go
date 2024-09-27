package domain


type TaobaoTbkDgMaterialOptionalUpgradeBasicMapData struct {
    /*
        商品信息-商品标题     */
    Title  *string `json:"title,omitempty" `

    /*
        商品信息-商品短标题     */
    ShortTitle  *string `json:"short_title,omitempty" `

    /*
        商品信息-商品主图     */
    PictUrl  *string `json:"pict_url,omitempty" `

    /*
        商品信息-商品白底图     */
    WhiteImage  *string `json:"white_image,omitempty" `

    /*
        商品信息-一级类目ID     */
    LevelOneCategoryId  *int64 `json:"level_one_category_id,omitempty" `

    /*
        商品信息-叶子类目id     */
    CategoryId  *int64 `json:"category_id,omitempty" `

    /*
        商品信息-叶子类目名称     */
    CategoryName  *string `json:"category_name,omitempty" `

    /*
        店铺信息-卖家id     */
    SellerId  *int64 `json:"seller_id,omitempty" `

    /*
        店铺信息-卖家类型，0表示淘宝，1表示天猫，3表示特价版     */
    UserType  *int64 `json:"user_type,omitempty" `

    /*
        店铺信息-店铺名称     */
    ShopTitle  *string `json:"shop_title,omitempty" `

    /*
        商品信息-30天销量；数据统计截止昨日非实时更新     */
    Volume  *int64 `json:"volume,omitempty" `

    /*
        商品信息-商品子标题     */
    SubTitle  *string `json:"sub_title,omitempty" `

    /*
        商品信息-品牌名称     */
    BrandName  *string `json:"brand_name,omitempty" `

    /*
        商品信息-一级类目名称     */
    LevelOneCategoryName  *string `json:"level_one_category_name,omitempty" `

    /*
        商品信息-淘客30天推广量     */
    TkTotalSales  *string `json:"tk_total_sales,omitempty" `

    /*
        商品信息-宝贝所在地     */
    Provcity  *string `json:"provcity,omitempty" `

    /*
        商品信息-商品小图列表     */
    SmallImages  *[]string `json:"small_images,omitempty" `

    /*
        年销量     */
    AnnualVol  *string `json:"annual_vol,omitempty" `

    /*
        商品邮费     */
    RealPostFee  *string `json:"real_post_fee,omitempty" `

}

func (s *TaobaoTbkDgMaterialOptionalUpgradeBasicMapData) SetTitle(v string) *TaobaoTbkDgMaterialOptionalUpgradeBasicMapData {
    s.Title = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradeBasicMapData) SetShortTitle(v string) *TaobaoTbkDgMaterialOptionalUpgradeBasicMapData {
    s.ShortTitle = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradeBasicMapData) SetPictUrl(v string) *TaobaoTbkDgMaterialOptionalUpgradeBasicMapData {
    s.PictUrl = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradeBasicMapData) SetWhiteImage(v string) *TaobaoTbkDgMaterialOptionalUpgradeBasicMapData {
    s.WhiteImage = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradeBasicMapData) SetLevelOneCategoryId(v int64) *TaobaoTbkDgMaterialOptionalUpgradeBasicMapData {
    s.LevelOneCategoryId = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradeBasicMapData) SetCategoryId(v int64) *TaobaoTbkDgMaterialOptionalUpgradeBasicMapData {
    s.CategoryId = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradeBasicMapData) SetCategoryName(v string) *TaobaoTbkDgMaterialOptionalUpgradeBasicMapData {
    s.CategoryName = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradeBasicMapData) SetSellerId(v int64) *TaobaoTbkDgMaterialOptionalUpgradeBasicMapData {
    s.SellerId = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradeBasicMapData) SetUserType(v int64) *TaobaoTbkDgMaterialOptionalUpgradeBasicMapData {
    s.UserType = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradeBasicMapData) SetShopTitle(v string) *TaobaoTbkDgMaterialOptionalUpgradeBasicMapData {
    s.ShopTitle = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradeBasicMapData) SetVolume(v int64) *TaobaoTbkDgMaterialOptionalUpgradeBasicMapData {
    s.Volume = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradeBasicMapData) SetSubTitle(v string) *TaobaoTbkDgMaterialOptionalUpgradeBasicMapData {
    s.SubTitle = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradeBasicMapData) SetBrandName(v string) *TaobaoTbkDgMaterialOptionalUpgradeBasicMapData {
    s.BrandName = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradeBasicMapData) SetLevelOneCategoryName(v string) *TaobaoTbkDgMaterialOptionalUpgradeBasicMapData {
    s.LevelOneCategoryName = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradeBasicMapData) SetTkTotalSales(v string) *TaobaoTbkDgMaterialOptionalUpgradeBasicMapData {
    s.TkTotalSales = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradeBasicMapData) SetProvcity(v string) *TaobaoTbkDgMaterialOptionalUpgradeBasicMapData {
    s.Provcity = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradeBasicMapData) SetSmallImages(v []string) *TaobaoTbkDgMaterialOptionalUpgradeBasicMapData {
    s.SmallImages = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradeBasicMapData) SetAnnualVol(v string) *TaobaoTbkDgMaterialOptionalUpgradeBasicMapData {
    s.AnnualVol = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradeBasicMapData) SetRealPostFee(v string) *TaobaoTbkDgMaterialOptionalUpgradeBasicMapData {
    s.RealPostFee = &v
    return s
}
