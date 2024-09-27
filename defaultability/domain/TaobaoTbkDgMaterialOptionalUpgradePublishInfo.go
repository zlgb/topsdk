package domain


type TaobaoTbkDgMaterialOptionalUpgradePublishInfo struct {
    /*
        商品信息-收入比率(%)；商品佣金比率+补贴比率     */
    IncomeRate  *string `json:"income_rate,omitempty" `

    /*
        前N件佣金信息-前N件佣金生效或预热时透出以下字段     */
    TopnInfo  *TaobaoTbkDgMaterialOptionalUpgradeTopNInfoDTO `json:"topn_info,omitempty" `

    /*
        链接-宝贝推广链接     */
    ClickUrl  *string `json:"click_url,omitempty" `

    /*
        链接-宝贝+券二合一页面链接  	     */
    CouponShareUrl  *string `json:"coupon_share_url,omitempty" `

    /*
        预热活动到手价对应的佣金比率     */
    FutureActivityCommissionRate  *string `json:"future_activity_commission_rate,omitempty" `

    /*
        预热价活动开始时间     */
    FutureActivityTime  *string `json:"future_activity_time,omitempty" `

    /*
        定向计划集合-仅支持联系商务或运营小二申请定向计划合集字段权限     */
    SpCampaignList  *[]TaobaoTbkDgMaterialOptionalUpgradeSpCampaign `json:"sp_campaign_list,omitempty" `

    /*
        榜单url     */
    RankPageUrl  *string `json:"rank_page_url,omitempty" `

    /*
        推广信息-商品信息-佣金类型。MKT表示营销计划，SP表示定向计划，COMMON表示通用计划，ZX表示自选计划     */
    CommissionType  *string `json:"commission_type,omitempty" `

    /*
        商品佣金信息     */
    IncomeInfo  *TaobaoTbkDgMaterialOptionalUpgradeFinalIncomeInfo `json:"income_info,omitempty" `

    /*
        额外奖励活动类型，如果一个商品有多个奖励类型，返回结果使用空格分割，0=预售单单奖励，1=618超级U选单单补     */
    CpaRewardType  *string `json:"cpa_reward_type,omitempty" `

    /*
        额外奖励活动金额，活动奖励金额的类型与cpa_reward_type字段对应，如果一个商品有多个奖励类型，返回结果使用空格分割     */
    CpaRewardAmount  *string `json:"cpa_reward_amount,omitempty" `

    /*
        商品是否包含定向计划     */
    IncludeDxjh  *string `json:"include_dxjh,omitempty" `

    /*
        两小时推广销量。 非实时，约半小时更新     */
    TwoHourPromotionSales  *int64 `json:"two_hour_promotion_sales,omitempty" `

    /*
        当天推广销量。 非实时，约1小时更新     */
    DailyPromotionSales  *int64 `json:"daily_promotion_sales,omitempty" `

}

func (s *TaobaoTbkDgMaterialOptionalUpgradePublishInfo) SetIncomeRate(v string) *TaobaoTbkDgMaterialOptionalUpgradePublishInfo {
    s.IncomeRate = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradePublishInfo) SetTopnInfo(v TaobaoTbkDgMaterialOptionalUpgradeTopNInfoDTO) *TaobaoTbkDgMaterialOptionalUpgradePublishInfo {
    s.TopnInfo = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradePublishInfo) SetClickUrl(v string) *TaobaoTbkDgMaterialOptionalUpgradePublishInfo {
    s.ClickUrl = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradePublishInfo) SetCouponShareUrl(v string) *TaobaoTbkDgMaterialOptionalUpgradePublishInfo {
    s.CouponShareUrl = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradePublishInfo) SetFutureActivityCommissionRate(v string) *TaobaoTbkDgMaterialOptionalUpgradePublishInfo {
    s.FutureActivityCommissionRate = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradePublishInfo) SetFutureActivityTime(v string) *TaobaoTbkDgMaterialOptionalUpgradePublishInfo {
    s.FutureActivityTime = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradePublishInfo) SetSpCampaignList(v []TaobaoTbkDgMaterialOptionalUpgradeSpCampaign) *TaobaoTbkDgMaterialOptionalUpgradePublishInfo {
    s.SpCampaignList = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradePublishInfo) SetRankPageUrl(v string) *TaobaoTbkDgMaterialOptionalUpgradePublishInfo {
    s.RankPageUrl = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradePublishInfo) SetCommissionType(v string) *TaobaoTbkDgMaterialOptionalUpgradePublishInfo {
    s.CommissionType = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradePublishInfo) SetIncomeInfo(v TaobaoTbkDgMaterialOptionalUpgradeFinalIncomeInfo) *TaobaoTbkDgMaterialOptionalUpgradePublishInfo {
    s.IncomeInfo = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradePublishInfo) SetCpaRewardType(v string) *TaobaoTbkDgMaterialOptionalUpgradePublishInfo {
    s.CpaRewardType = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradePublishInfo) SetCpaRewardAmount(v string) *TaobaoTbkDgMaterialOptionalUpgradePublishInfo {
    s.CpaRewardAmount = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradePublishInfo) SetIncludeDxjh(v string) *TaobaoTbkDgMaterialOptionalUpgradePublishInfo {
    s.IncludeDxjh = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradePublishInfo) SetTwoHourPromotionSales(v int64) *TaobaoTbkDgMaterialOptionalUpgradePublishInfo {
    s.TwoHourPromotionSales = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradePublishInfo) SetDailyPromotionSales(v int64) *TaobaoTbkDgMaterialOptionalUpgradePublishInfo {
    s.DailyPromotionSales = &v
    return s
}
