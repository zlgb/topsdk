package request


type TaobaoTbkShopGetRequest struct {
    /*
        累计推广商品上限     */
    EndAuctionCount  *int64 `json:"end_auction_count,omitempty" required:"false" `
    /*
        淘客佣金比率上限，1~10000     */
    EndCommissionRate  *int64 `json:"end_commission_rate,omitempty" required:"false" `
    /*
        信用等级上限，1~20     */
    EndCredit  *int64 `json:"end_credit,omitempty" required:"false" `
    /*
        店铺商品总数上限     */
    EndTotalAction  *int64 `json:"end_total_action,omitempty" required:"false" `
    /*
        需返回的字段列表     */
    Fields  *string `json:"fields" required:"true" `
    /*
        是否商城的店铺，设置为true表示该是属于淘宝商城的店铺，设置为false或不设置表示不判断这个属性     */
    IsTmall  *bool `json:"is_tmall,omitempty" required:"false" `
    /*
        第几页，默认1，1~100 defalutValue��1    */
    PageNo  *int64 `json:"page_no,omitempty" required:"false" `
    /*
        页大小，默认20，1~100 defalutValue��20    */
    PageSize  *int64 `json:"page_size,omitempty" required:"false" `
    /*
        链接形式：1：PC，2：无线，默认：１ defalutValue��1    */
    Platform  *int64 `json:"platform,omitempty" required:"false" `
    /*
        查询词     */
    Q  *string `json:"q" required:"true" `
    /*
        排序_des（降序），排序_asc（升序），佣金比率（commission_rate）， 商品数量（auction_count），销售总数量（total_auction）     */
    Sort  *string `json:"sort,omitempty" required:"false" `
    /*
        累计推广商品下限     */
    StartAuctionCount  *int64 `json:"start_auction_count,omitempty" required:"false" `
    /*
        淘客佣金比率下限，1~10000     */
    StartCommissionRate  *int64 `json:"start_commission_rate,omitempty" required:"false" `
    /*
        信用等级下限，1~20     */
    StartCredit  *int64 `json:"start_credit,omitempty" required:"false" `
    /*
        店铺商品总数下限     */
    StartTotalAction  *int64 `json:"start_total_action,omitempty" required:"false" `
}

func (s *TaobaoTbkShopGetRequest) SetEndAuctionCount(v int64) *TaobaoTbkShopGetRequest {
    s.EndAuctionCount = &v
    return s
}
func (s *TaobaoTbkShopGetRequest) SetEndCommissionRate(v int64) *TaobaoTbkShopGetRequest {
    s.EndCommissionRate = &v
    return s
}
func (s *TaobaoTbkShopGetRequest) SetEndCredit(v int64) *TaobaoTbkShopGetRequest {
    s.EndCredit = &v
    return s
}
func (s *TaobaoTbkShopGetRequest) SetEndTotalAction(v int64) *TaobaoTbkShopGetRequest {
    s.EndTotalAction = &v
    return s
}
func (s *TaobaoTbkShopGetRequest) SetFields(v string) *TaobaoTbkShopGetRequest {
    s.Fields = &v
    return s
}
func (s *TaobaoTbkShopGetRequest) SetIsTmall(v bool) *TaobaoTbkShopGetRequest {
    s.IsTmall = &v
    return s
}
func (s *TaobaoTbkShopGetRequest) SetPageNo(v int64) *TaobaoTbkShopGetRequest {
    s.PageNo = &v
    return s
}
func (s *TaobaoTbkShopGetRequest) SetPageSize(v int64) *TaobaoTbkShopGetRequest {
    s.PageSize = &v
    return s
}
func (s *TaobaoTbkShopGetRequest) SetPlatform(v int64) *TaobaoTbkShopGetRequest {
    s.Platform = &v
    return s
}
func (s *TaobaoTbkShopGetRequest) SetQ(v string) *TaobaoTbkShopGetRequest {
    s.Q = &v
    return s
}
func (s *TaobaoTbkShopGetRequest) SetSort(v string) *TaobaoTbkShopGetRequest {
    s.Sort = &v
    return s
}
func (s *TaobaoTbkShopGetRequest) SetStartAuctionCount(v int64) *TaobaoTbkShopGetRequest {
    s.StartAuctionCount = &v
    return s
}
func (s *TaobaoTbkShopGetRequest) SetStartCommissionRate(v int64) *TaobaoTbkShopGetRequest {
    s.StartCommissionRate = &v
    return s
}
func (s *TaobaoTbkShopGetRequest) SetStartCredit(v int64) *TaobaoTbkShopGetRequest {
    s.StartCredit = &v
    return s
}
func (s *TaobaoTbkShopGetRequest) SetStartTotalAction(v int64) *TaobaoTbkShopGetRequest {
    s.StartTotalAction = &v
    return s
}

func (req *TaobaoTbkShopGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.EndAuctionCount != nil) {
        paramMap["end_auction_count"] = *req.EndAuctionCount
    }
    if(req.EndCommissionRate != nil) {
        paramMap["end_commission_rate"] = *req.EndCommissionRate
    }
    if(req.EndCredit != nil) {
        paramMap["end_credit"] = *req.EndCredit
    }
    if(req.EndTotalAction != nil) {
        paramMap["end_total_action"] = *req.EndTotalAction
    }
    if(req.Fields != nil) {
        paramMap["fields"] = *req.Fields
    }
    if(req.IsTmall != nil) {
        paramMap["is_tmall"] = *req.IsTmall
    }
    if(req.PageNo != nil) {
        paramMap["page_no"] = *req.PageNo
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.Platform != nil) {
        paramMap["platform"] = *req.Platform
    }
    if(req.Q != nil) {
        paramMap["q"] = *req.Q
    }
    if(req.Sort != nil) {
        paramMap["sort"] = *req.Sort
    }
    if(req.StartAuctionCount != nil) {
        paramMap["start_auction_count"] = *req.StartAuctionCount
    }
    if(req.StartCommissionRate != nil) {
        paramMap["start_commission_rate"] = *req.StartCommissionRate
    }
    if(req.StartCredit != nil) {
        paramMap["start_credit"] = *req.StartCredit
    }
    if(req.StartTotalAction != nil) {
        paramMap["start_total_action"] = *req.StartTotalAction
    }
    return paramMap
}

func (req *TaobaoTbkShopGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}