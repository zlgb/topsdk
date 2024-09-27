package domain


type TaobaoTbkDgMaterialRecommendSpCampaign struct {
    /*
        定向计划活动ID     */
    SpCid  *string `json:"sp_cid,omitempty" `

    /*
        定向计划名称     */
    SpName  *string `json:"sp_name,omitempty" `

    /*
        定向佣金率     */
    SpRate  *string `json:"sp_rate,omitempty" `

    /*
        定向是否锁佣，0=不锁佣 1=锁佣     */
    SpLockStatus  *string `json:"sp_lock_status,omitempty" `

    /*
        定向计划申请链接     */
    SpApplyLink  *string `json:"sp_apply_link,omitempty" `

    /*
        定向计划是否可用 1-可用 0-不可用     */
    SpStatus  *string `json:"sp_status,omitempty" `

}

func (s *TaobaoTbkDgMaterialRecommendSpCampaign) SetSpCid(v string) *TaobaoTbkDgMaterialRecommendSpCampaign {
    s.SpCid = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendSpCampaign) SetSpName(v string) *TaobaoTbkDgMaterialRecommendSpCampaign {
    s.SpName = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendSpCampaign) SetSpRate(v string) *TaobaoTbkDgMaterialRecommendSpCampaign {
    s.SpRate = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendSpCampaign) SetSpLockStatus(v string) *TaobaoTbkDgMaterialRecommendSpCampaign {
    s.SpLockStatus = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendSpCampaign) SetSpApplyLink(v string) *TaobaoTbkDgMaterialRecommendSpCampaign {
    s.SpApplyLink = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendSpCampaign) SetSpStatus(v string) *TaobaoTbkDgMaterialRecommendSpCampaign {
    s.SpStatus = &v
    return s
}
