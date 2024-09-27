package domain


type TaobaoTbkDgMaterialTemporaryOptionalMapData struct {
    /*
        优惠券信息-优惠券开始时间     */
    CouponStartTime  *string `json:"coupon_start_time,omitempty" `

    /*
        优惠券信息-优惠券结束时间     */
    CouponEndTime  *string `json:"coupon_end_time,omitempty" `

    /*
        商品信息-定向计划信息     */
    InfoDxjh  *string `json:"info_dxjh,omitempty" `

    /*
        商品信息-淘客30天推广量     */
    TkTotalSales  *string `json:"tk_total_sales,omitempty" `

    /*
        商品信息-月支出佣金(该字段废弃，请勿再用)     */
    TkTotalCommi  *string `json:"tk_total_commi,omitempty" `

    /*
        优惠券信息-优惠券id     */
    CouponId  *string `json:"coupon_id,omitempty" `

    /*
        商品信息-宝贝id(该字段废弃，请勿再用)     */
    NumIid  *string `json:"num_iid,omitempty" `

    /*
        商品信息-商品标题     */
    Title  *string `json:"title,omitempty" `

    /*
        商品信息-商品主图     */
    PictUrl  *string `json:"pict_url,omitempty" `

    /*
        商品信息-商品小图列表     */
    SmallImages  *[]string `json:"small_images,omitempty" `

    /*
        商品信息-商品一口价格     */
    ReservePrice  *string `json:"reserve_price,omitempty" `

    /*
        折扣价（元） 若属于预售商品，付定金时间内，折扣价=预售价     */
    ZkFinalPrice  *string `json:"zk_final_price,omitempty" `

    /*
        店铺信息-卖家类型。0表示集市，1表示天猫     */
    UserType  *int64 `json:"user_type,omitempty" `

    /*
        商品信息-宝贝所在地     */
    Provcity  *string `json:"provcity,omitempty" `

    /*
        链接-宝贝地址     */
    ItemUrl  *string `json:"item_url,omitempty" `

    /*
        商品信息-是否包含营销计划     */
    IncludeMkt  *string `json:"include_mkt,omitempty" `

    /*
        商品信息-是否包含定向计划     */
    IncludeDxjh  *string `json:"include_dxjh,omitempty" `

    /*
        商品信息-佣金比率。1550表示15.5%     */
    CommissionRate  *string `json:"commission_rate,omitempty" `

    /*
        商品信息-30天销量（饿了么卡券信息-总销量）     */
    Volume  *int64 `json:"volume,omitempty" `

    /*
        店铺信息-店铺名称     */
    ShopTitle  *string `json:"shop_title,omitempty" `

    /*
        优惠券信息-优惠券总量     */
    CouponTotalCount  *int64 `json:"coupon_total_count,omitempty" `

    /*
        优惠券信息-优惠券剩余量     */
    CouponRemainCount  *int64 `json:"coupon_remain_count,omitempty" `

    /*
        优惠券信息-优惠券满减信息     */
    CouponInfo  *string `json:"coupon_info,omitempty" `

    /*
        商品信息-佣金类型。MKT表示营销计划，SP表示定向计划，COMMON表示通用计划     */
    CommissionType  *string `json:"commission_type,omitempty" `

    /*
        店铺信息-店铺dsr评分     */
    ShopDsr  *int64 `json:"shop_dsr,omitempty" `

    /*
        链接-宝贝+券二合一页面链接     */
    CouponShareUrl  *string `json:"coupon_share_url,omitempty" `

    /*
        链接-宝贝推广链接     */
    Url  *string `json:"url,omitempty" `

    /*
        商品信息-一级类目名称     */
    LevelOneCategoryName  *string `json:"level_one_category_name,omitempty" `

    /*
        商品信息-一级类目ID     */
    LevelOneCategoryId  *int64 `json:"level_one_category_id,omitempty" `

    /*
        商品信息-叶子类目名称     */
    CategoryName  *string `json:"category_name,omitempty" `

    /*
        商品信息-叶子类目id     */
    CategoryId  *int64 `json:"category_id,omitempty" `

    /*
        商品信息-商品短标题     */
    ShortTitle  *string `json:"short_title,omitempty" `

    /*
        商品信息-商品白底图     */
    WhiteImage  *string `json:"white_image,omitempty" `

    /*
        拼团专用-拼团结束时间     */
    Oetime  *string `json:"oetime,omitempty" `

    /*
        拼团专用-拼团开始时间     */
    Ostime  *string `json:"ostime,omitempty" `

    /*
        拼团专用-拼团几人团     */
    JddNum  *int64 `json:"jdd_num,omitempty" `

    /*
        拼团专用-拼团拼成价，单位元     */
    JddPrice  *string `json:"jdd_price,omitempty" `

    /*
        预售专用-预售数量     */
    UvSumPreSale  *int64 `json:"uv_sum_pre_sale,omitempty" `

    /*
        链接-物料块id(测试中请勿使用)     */
    XId  *string `json:"x_id,omitempty" `

    /*
        优惠券信息-优惠券起用门槛，满X元可用。如：满299元减20元     */
    CouponStartFee  *string `json:"coupon_start_fee,omitempty" `

    /*
        优惠券（元） 若属于预售商品，该优惠券付尾款可用，付定金不可用     */
    CouponAmount  *string `json:"coupon_amount,omitempty" `

    /*
        商品信息-宝贝描述(推荐理由)     */
    ItemDescription  *string `json:"item_description,omitempty" `

    /*
        店铺信息-卖家昵称     */
    Nick  *string `json:"nick,omitempty" `

    /*
        拼团专用-拼团一人价（原价)，单位元     */
    OrigPrice  *string `json:"orig_price,omitempty" `

    /*
        拼团专用-拼团库存数量     */
    TotalStock  *int64 `json:"total_stock,omitempty" `

    /*
        拼团专用-拼团已售数量     */
    SellNum  *int64 `json:"sell_num,omitempty" `

    /*
        拼团专用-拼团剩余库存     */
    Stock  *int64 `json:"stock,omitempty" `

    /*
        营销-天猫营销玩法     */
    TmallPlayActivityInfo  *string `json:"tmall_play_activity_info,omitempty" `

    /*
        商品信息-宝贝id     */
    ItemId  *string `json:"item_id,omitempty" `

    /*
        商品邮费     */
    RealPostFee  *string `json:"real_post_fee,omitempty" `

    /*
        锁住的佣金率     */
    LockRate  *string `json:"lock_rate,omitempty" `

    /*
        锁佣结束时间     */
    LockRateEndTime  *int64 `json:"lock_rate_end_time,omitempty" `

    /*
        锁佣开始时间     */
    LockRateStartTime  *int64 `json:"lock_rate_start_time,omitempty" `

    /*
        预售商品-优惠     */
    PresaleDiscountFeeText  *string `json:"presale_discount_fee_text,omitempty" `

    /*
        预售商品-付尾款结束时间（毫秒）     */
    PresaleTailEndTime  *int64 `json:"presale_tail_end_time,omitempty" `

    /*
        预售商品-付尾款开始时间（毫秒）     */
    PresaleTailStartTime  *int64 `json:"presale_tail_start_time,omitempty" `

    /*
        预售商品-付定金结束时间（毫秒）     */
    PresaleEndTime  *int64 `json:"presale_end_time,omitempty" `

    /*
        预售商品-付定金开始时间（毫秒）     */
    PresaleStartTime  *int64 `json:"presale_start_time,omitempty" `

    /*
        预售商品-定金（元）     */
    PresaleDeposit  *string `json:"presale_deposit,omitempty" `

    /*
        预售有礼-淘礼金发放时间     */
    YsylTljSendTime  *string `json:"ysyl_tlj_send_time,omitempty" `

    /*
        预售有礼-佣金比例（ 预售有礼活动享受的推广佣金比例，注：推广该活动有特殊分成规则，请详见：https://tbk.bbs.taobao.com/detail.html?appId=45301&postId=9334376 ）     */
    YsylCommissionRate  *string `json:"ysyl_commission_rate,omitempty" `

    /*
        预售有礼-预估淘礼金（元）     */
    YsylTljFace  *string `json:"ysyl_tlj_face,omitempty" `

    /*
        预售有礼-推广链接     */
    YsylClickUrl  *string `json:"ysyl_click_url,omitempty" `

    /*
        预售有礼-淘礼金使用结束时间     */
    YsylTljUseEndTime  *string `json:"ysyl_tlj_use_end_time,omitempty" `

    /*
        预售有礼-淘礼金使用开始时间     */
    YsylTljUseStartTime  *string `json:"ysyl_tlj_use_start_time,omitempty" `

    /*
        本地化-销售开始时间     */
    SaleBeginTime  *string `json:"sale_begin_time,omitempty" `

    /*
        本地化-销售结束时间     */
    SaleEndTime  *string `json:"sale_end_time,omitempty" `

    /*
        本地化-到门店距离（米）     */
    Distance  *string `json:"distance,omitempty" `

    /*
        本地化-可用店铺id     */
    UsableShopId  *string `json:"usable_shop_id,omitempty" `

    /*
        本地化-可用店铺名称     */
    UsableShopName  *string `json:"usable_shop_name,omitempty" `

    /*
        活动价     */
    SalePrice  *string `json:"sale_price,omitempty" `

    /*
        跨店满减信息     */
    KuadianPromotionInfo  *string `json:"kuadian_promotion_info,omitempty" `

    /*
        是否品牌精选，0不是，1是     */
    SuperiorBrand  *string `json:"superior_brand,omitempty" `

    /*
        比价场景专用，当系统检测到入参消费者ID购买当前商品会获得《天天开彩蛋》玩法的彩蛋时，该字段显示1，否则为0     */
    RewardInfo  *int64 `json:"reward_info,omitempty" `

    /*
        是否品牌快抢，0不是，1是     */
    IsBrandFlashSale  *string `json:"is_brand_flash_sale,omitempty" `

    /*
        本地化-扩展信息     */
    LocalizationExtend  *string `json:"localization_extend,omitempty" `

    /*
        物料评估-匹配分     */
    MatchScore  *string `json:"match_score,omitempty" `

    /*
        物料评估-收益分     */
    CommiScore  *string `json:"commi_score,omitempty" `

    /*
        是否是热门商品，0不是，1是     */
    HotFlag  *string `json:"hot_flag,omitempty" `

    /*
        前N件佣金信息-前N件佣金生效或预热时透出以下字段     */
    TopnInfo  *TaobaoTbkDgMaterialTemporaryOptionalStepRateDTO `json:"topn_info,omitempty" `

    /*
        百亿补贴信息     */
    BybtInfo  *TaobaoTbkDgMaterialTemporaryOptionalTopNInfoDTO `json:"bybt_info,omitempty" `

    /*
        商品入驻淘特后产生的所有销量量级，不特指某段具体时间     */
    TtSoldCount  *string `json:"tt_sold_count,omitempty" `

    /*
        猫超买返卡信息     */
    MaifanPromotion  *TaobaoTbkDgMaterialTemporaryOptionalMaifanPromotionDTO `json:"maifan_promotion,omitempty" `

    /*
        额外奖励活动类型，如果一个商品有多个奖励类型，返回结果使用空格分割，0=单单奖励，1=超级单单奖励     */
    CpaRewardType  *string `json:"cpa_reward_type,omitempty" `

}

func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetCouponStartTime(v string) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.CouponStartTime = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetCouponEndTime(v string) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.CouponEndTime = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetInfoDxjh(v string) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.InfoDxjh = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetTkTotalSales(v string) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.TkTotalSales = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetTkTotalCommi(v string) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.TkTotalCommi = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetCouponId(v string) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.CouponId = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetNumIid(v string) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.NumIid = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetTitle(v string) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.Title = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetPictUrl(v string) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.PictUrl = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetSmallImages(v []string) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.SmallImages = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetReservePrice(v string) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.ReservePrice = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetZkFinalPrice(v string) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.ZkFinalPrice = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetUserType(v int64) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.UserType = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetProvcity(v string) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.Provcity = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetItemUrl(v string) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.ItemUrl = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetIncludeMkt(v string) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.IncludeMkt = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetIncludeDxjh(v string) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.IncludeDxjh = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetCommissionRate(v string) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.CommissionRate = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetVolume(v int64) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.Volume = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetShopTitle(v string) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.ShopTitle = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetCouponTotalCount(v int64) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.CouponTotalCount = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetCouponRemainCount(v int64) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.CouponRemainCount = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetCouponInfo(v string) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.CouponInfo = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetCommissionType(v string) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.CommissionType = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetShopDsr(v int64) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.ShopDsr = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetCouponShareUrl(v string) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.CouponShareUrl = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetUrl(v string) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.Url = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetLevelOneCategoryName(v string) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.LevelOneCategoryName = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetLevelOneCategoryId(v int64) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.LevelOneCategoryId = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetCategoryName(v string) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.CategoryName = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetCategoryId(v int64) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.CategoryId = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetShortTitle(v string) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.ShortTitle = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetWhiteImage(v string) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.WhiteImage = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetOetime(v string) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.Oetime = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetOstime(v string) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.Ostime = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetJddNum(v int64) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.JddNum = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetJddPrice(v string) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.JddPrice = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetUvSumPreSale(v int64) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.UvSumPreSale = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetXId(v string) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.XId = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetCouponStartFee(v string) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.CouponStartFee = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetCouponAmount(v string) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.CouponAmount = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetItemDescription(v string) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.ItemDescription = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetNick(v string) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.Nick = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetOrigPrice(v string) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.OrigPrice = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetTotalStock(v int64) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.TotalStock = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetSellNum(v int64) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.SellNum = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetStock(v int64) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.Stock = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetTmallPlayActivityInfo(v string) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.TmallPlayActivityInfo = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetItemId(v string) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.ItemId = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetRealPostFee(v string) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.RealPostFee = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetLockRate(v string) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.LockRate = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetLockRateEndTime(v int64) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.LockRateEndTime = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetLockRateStartTime(v int64) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.LockRateStartTime = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetPresaleDiscountFeeText(v string) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.PresaleDiscountFeeText = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetPresaleTailEndTime(v int64) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.PresaleTailEndTime = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetPresaleTailStartTime(v int64) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.PresaleTailStartTime = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetPresaleEndTime(v int64) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.PresaleEndTime = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetPresaleStartTime(v int64) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.PresaleStartTime = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetPresaleDeposit(v string) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.PresaleDeposit = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetYsylTljSendTime(v string) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.YsylTljSendTime = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetYsylCommissionRate(v string) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.YsylCommissionRate = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetYsylTljFace(v string) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.YsylTljFace = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetYsylClickUrl(v string) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.YsylClickUrl = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetYsylTljUseEndTime(v string) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.YsylTljUseEndTime = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetYsylTljUseStartTime(v string) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.YsylTljUseStartTime = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetSaleBeginTime(v string) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.SaleBeginTime = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetSaleEndTime(v string) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.SaleEndTime = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetDistance(v string) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.Distance = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetUsableShopId(v string) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.UsableShopId = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetUsableShopName(v string) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.UsableShopName = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetSalePrice(v string) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.SalePrice = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetKuadianPromotionInfo(v string) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.KuadianPromotionInfo = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetSuperiorBrand(v string) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.SuperiorBrand = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetRewardInfo(v int64) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.RewardInfo = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetIsBrandFlashSale(v string) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.IsBrandFlashSale = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetLocalizationExtend(v string) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.LocalizationExtend = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetMatchScore(v string) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.MatchScore = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetCommiScore(v string) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.CommiScore = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetHotFlag(v string) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.HotFlag = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetTopnInfo(v TaobaoTbkDgMaterialTemporaryOptionalStepRateDTO) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.TopnInfo = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetBybtInfo(v TaobaoTbkDgMaterialTemporaryOptionalTopNInfoDTO) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.BybtInfo = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetTtSoldCount(v string) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.TtSoldCount = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetMaifanPromotion(v TaobaoTbkDgMaterialTemporaryOptionalMaifanPromotionDTO) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.MaifanPromotion = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalMapData) SetCpaRewardType(v string) *TaobaoTbkDgMaterialTemporaryOptionalMapData {
    s.CpaRewardType = &v
    return s
}
