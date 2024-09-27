package domain


type TaobaoTbkItemInfoUpgradeGetBasicMapData struct {
    /*
        商品标题     */
    Title  *string `json:"title,omitempty" `

    /*
        商品主图     */
    PictUrl  *string `json:"pict_url,omitempty" `

    /*
        商品小图列表     */
    SmallImages  *[]string `json:"small_images,omitempty" `

    /*
        一级类目名称     */
    LevelOneCategoryName  *string `json:"level_one_category_name,omitempty" `

    /*
        叶子类目名称     */
    CategoryName  *string `json:"category_name,omitempty" `

    /*
        商品所在地     */
    Provcity  *string `json:"provcity,omitempty" `

    /*
        商品链接     */
    ItemUrl  *string `json:"item_url,omitempty" `

    /*
        卖家类型，0表示集市，1表示商城，3表示特价版     */
    UserType  *int64 `json:"user_type,omitempty" `

    /*
        卖家等级(字段等级SA)     */
    Ratesum  *int64 `json:"ratesum,omitempty" `

    /*
        店铺dsr 评分(字段等级SA)     */
    ShopDsr  *int64 `json:"shop_dsr,omitempty" `

    /*
        退款率是否低于行业均值(字段等级SA)     */
    IRfdRate  *bool `json:"i_rfd_rate,omitempty" `

    /*
        好评率是否高于行业均值(字段等级SA)     */
    HGoodRate  *bool `json:"h_good_rate,omitempty" `

    /*
        成交转化是否高于行业均值(字段等级SA)     */
    HPayRate30  *bool `json:"h_pay_rate30,omitempty" `

    /*
        30天销量；数据统计截止昨日非实时更新     */
    Volume  *int64 `json:"volume,omitempty" `

    /*
        是否加入消费者保障     */
    IsPrepay  *bool `json:"is_prepay,omitempty" `

    /*
        是否品牌精选，0不是，1是     */
    SuperiorBrand  *string `json:"superior_brand,omitempty" `

    /*
        商品库类型，支持多库类型输出，以英文逗号分隔“,”分隔，1:营销商品主推库，2. 内容商品库，如果值为空则不属于1，2这两种商品类型     */
    MaterialLibType  *string `json:"material_lib_type,omitempty" `

    /*
        pc宝贝详情(字段等级S)     */
    TmallDescUrl  *string `json:"tmall_desc_url,omitempty" `

    /*
        H5宝贝详情(字段等级S)     */
    TaobaoDescUrl  *string `json:"taobao_desc_url,omitempty" `

    /*
        商品信息-店铺名称     */
    ShopTitle  *string `json:"shop_title,omitempty" `

    /*
        商品信息-是否包邮     */
    FreeShipment  *bool `json:"free_shipment,omitempty" `

    /*
        商品信息-品牌名称     */
    BrandName  *string `json:"brand_name,omitempty" `

    /*
        商品信息-商品短标题     */
    ShortTitle  *string `json:"short_title,omitempty" `

    /*
        商品信息-商品白底图     */
    WhiteImage  *string `json:"white_image,omitempty" `

    /*
        卖家id(字段等级C)     */
    SellerId  *int64 `json:"seller_id,omitempty" `

    /*
        年销量     */
    AnnualVol  *string `json:"annual_vol,omitempty" `

}

func (s *TaobaoTbkItemInfoUpgradeGetBasicMapData) SetTitle(v string) *TaobaoTbkItemInfoUpgradeGetBasicMapData {
    s.Title = &v
    return s
}
func (s *TaobaoTbkItemInfoUpgradeGetBasicMapData) SetPictUrl(v string) *TaobaoTbkItemInfoUpgradeGetBasicMapData {
    s.PictUrl = &v
    return s
}
func (s *TaobaoTbkItemInfoUpgradeGetBasicMapData) SetSmallImages(v []string) *TaobaoTbkItemInfoUpgradeGetBasicMapData {
    s.SmallImages = &v
    return s
}
func (s *TaobaoTbkItemInfoUpgradeGetBasicMapData) SetLevelOneCategoryName(v string) *TaobaoTbkItemInfoUpgradeGetBasicMapData {
    s.LevelOneCategoryName = &v
    return s
}
func (s *TaobaoTbkItemInfoUpgradeGetBasicMapData) SetCategoryName(v string) *TaobaoTbkItemInfoUpgradeGetBasicMapData {
    s.CategoryName = &v
    return s
}
func (s *TaobaoTbkItemInfoUpgradeGetBasicMapData) SetProvcity(v string) *TaobaoTbkItemInfoUpgradeGetBasicMapData {
    s.Provcity = &v
    return s
}
func (s *TaobaoTbkItemInfoUpgradeGetBasicMapData) SetItemUrl(v string) *TaobaoTbkItemInfoUpgradeGetBasicMapData {
    s.ItemUrl = &v
    return s
}
func (s *TaobaoTbkItemInfoUpgradeGetBasicMapData) SetUserType(v int64) *TaobaoTbkItemInfoUpgradeGetBasicMapData {
    s.UserType = &v
    return s
}
func (s *TaobaoTbkItemInfoUpgradeGetBasicMapData) SetRatesum(v int64) *TaobaoTbkItemInfoUpgradeGetBasicMapData {
    s.Ratesum = &v
    return s
}
func (s *TaobaoTbkItemInfoUpgradeGetBasicMapData) SetShopDsr(v int64) *TaobaoTbkItemInfoUpgradeGetBasicMapData {
    s.ShopDsr = &v
    return s
}
func (s *TaobaoTbkItemInfoUpgradeGetBasicMapData) SetIRfdRate(v bool) *TaobaoTbkItemInfoUpgradeGetBasicMapData {
    s.IRfdRate = &v
    return s
}
func (s *TaobaoTbkItemInfoUpgradeGetBasicMapData) SetHGoodRate(v bool) *TaobaoTbkItemInfoUpgradeGetBasicMapData {
    s.HGoodRate = &v
    return s
}
func (s *TaobaoTbkItemInfoUpgradeGetBasicMapData) SetHPayRate30(v bool) *TaobaoTbkItemInfoUpgradeGetBasicMapData {
    s.HPayRate30 = &v
    return s
}
func (s *TaobaoTbkItemInfoUpgradeGetBasicMapData) SetVolume(v int64) *TaobaoTbkItemInfoUpgradeGetBasicMapData {
    s.Volume = &v
    return s
}
func (s *TaobaoTbkItemInfoUpgradeGetBasicMapData) SetIsPrepay(v bool) *TaobaoTbkItemInfoUpgradeGetBasicMapData {
    s.IsPrepay = &v
    return s
}
func (s *TaobaoTbkItemInfoUpgradeGetBasicMapData) SetSuperiorBrand(v string) *TaobaoTbkItemInfoUpgradeGetBasicMapData {
    s.SuperiorBrand = &v
    return s
}
func (s *TaobaoTbkItemInfoUpgradeGetBasicMapData) SetMaterialLibType(v string) *TaobaoTbkItemInfoUpgradeGetBasicMapData {
    s.MaterialLibType = &v
    return s
}
func (s *TaobaoTbkItemInfoUpgradeGetBasicMapData) SetTmallDescUrl(v string) *TaobaoTbkItemInfoUpgradeGetBasicMapData {
    s.TmallDescUrl = &v
    return s
}
func (s *TaobaoTbkItemInfoUpgradeGetBasicMapData) SetTaobaoDescUrl(v string) *TaobaoTbkItemInfoUpgradeGetBasicMapData {
    s.TaobaoDescUrl = &v
    return s
}
func (s *TaobaoTbkItemInfoUpgradeGetBasicMapData) SetShopTitle(v string) *TaobaoTbkItemInfoUpgradeGetBasicMapData {
    s.ShopTitle = &v
    return s
}
func (s *TaobaoTbkItemInfoUpgradeGetBasicMapData) SetFreeShipment(v bool) *TaobaoTbkItemInfoUpgradeGetBasicMapData {
    s.FreeShipment = &v
    return s
}
func (s *TaobaoTbkItemInfoUpgradeGetBasicMapData) SetBrandName(v string) *TaobaoTbkItemInfoUpgradeGetBasicMapData {
    s.BrandName = &v
    return s
}
func (s *TaobaoTbkItemInfoUpgradeGetBasicMapData) SetShortTitle(v string) *TaobaoTbkItemInfoUpgradeGetBasicMapData {
    s.ShortTitle = &v
    return s
}
func (s *TaobaoTbkItemInfoUpgradeGetBasicMapData) SetWhiteImage(v string) *TaobaoTbkItemInfoUpgradeGetBasicMapData {
    s.WhiteImage = &v
    return s
}
func (s *TaobaoTbkItemInfoUpgradeGetBasicMapData) SetSellerId(v int64) *TaobaoTbkItemInfoUpgradeGetBasicMapData {
    s.SellerId = &v
    return s
}
func (s *TaobaoTbkItemInfoUpgradeGetBasicMapData) SetAnnualVol(v string) *TaobaoTbkItemInfoUpgradeGetBasicMapData {
    s.AnnualVol = &v
    return s
}
