package domain


type TaobaoTbkDgMaterialRecommendBasicMapData struct {
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
        商品信息-一级类目名称     */
    LevelOneCategoryName  *string `json:"level_one_category_name,omitempty" `

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
        年销量     */
    AnnualVol  *string `json:"annual_vol,omitempty" `

    /*
        商品信息-商品小图列表     */
    SmallImages  *[]string `json:"small_images,omitempty" `

    /*
        商品邮费     */
    RealPostFee  *string `json:"real_post_fee,omitempty" `

}

func (s *TaobaoTbkDgMaterialRecommendBasicMapData) SetTitle(v string) *TaobaoTbkDgMaterialRecommendBasicMapData {
    s.Title = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendBasicMapData) SetShortTitle(v string) *TaobaoTbkDgMaterialRecommendBasicMapData {
    s.ShortTitle = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendBasicMapData) SetPictUrl(v string) *TaobaoTbkDgMaterialRecommendBasicMapData {
    s.PictUrl = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendBasicMapData) SetWhiteImage(v string) *TaobaoTbkDgMaterialRecommendBasicMapData {
    s.WhiteImage = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendBasicMapData) SetLevelOneCategoryId(v int64) *TaobaoTbkDgMaterialRecommendBasicMapData {
    s.LevelOneCategoryId = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendBasicMapData) SetLevelOneCategoryName(v string) *TaobaoTbkDgMaterialRecommendBasicMapData {
    s.LevelOneCategoryName = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendBasicMapData) SetCategoryId(v int64) *TaobaoTbkDgMaterialRecommendBasicMapData {
    s.CategoryId = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendBasicMapData) SetCategoryName(v string) *TaobaoTbkDgMaterialRecommendBasicMapData {
    s.CategoryName = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendBasicMapData) SetSellerId(v int64) *TaobaoTbkDgMaterialRecommendBasicMapData {
    s.SellerId = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendBasicMapData) SetUserType(v int64) *TaobaoTbkDgMaterialRecommendBasicMapData {
    s.UserType = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendBasicMapData) SetShopTitle(v string) *TaobaoTbkDgMaterialRecommendBasicMapData {
    s.ShopTitle = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendBasicMapData) SetVolume(v int64) *TaobaoTbkDgMaterialRecommendBasicMapData {
    s.Volume = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendBasicMapData) SetSubTitle(v string) *TaobaoTbkDgMaterialRecommendBasicMapData {
    s.SubTitle = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendBasicMapData) SetBrandName(v string) *TaobaoTbkDgMaterialRecommendBasicMapData {
    s.BrandName = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendBasicMapData) SetAnnualVol(v string) *TaobaoTbkDgMaterialRecommendBasicMapData {
    s.AnnualVol = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendBasicMapData) SetSmallImages(v []string) *TaobaoTbkDgMaterialRecommendBasicMapData {
    s.SmallImages = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendBasicMapData) SetRealPostFee(v string) *TaobaoTbkDgMaterialRecommendBasicMapData {
    s.RealPostFee = &v
    return s
}
