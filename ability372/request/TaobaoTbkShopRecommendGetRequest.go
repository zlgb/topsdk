package request


type TaobaoTbkShopRecommendGetRequest struct {
    /*
        返回数量，默认20，最大值40     */
    Count  *int64 `json:"count,omitempty" required:"false" `
    /*
        需返回的字段列表     */
    Fields  *string `json:"fields" required:"true" `
    /*
        链接形式：1：PC，2：无线，默认：１     */
    Platform  *int64 `json:"platform,omitempty" required:"false" `
    /*
        卖家Id     */
    UserId  *int64 `json:"user_id" required:"true" `
}

func (s *TaobaoTbkShopRecommendGetRequest) SetCount(v int64) *TaobaoTbkShopRecommendGetRequest {
    s.Count = &v
    return s
}
func (s *TaobaoTbkShopRecommendGetRequest) SetFields(v string) *TaobaoTbkShopRecommendGetRequest {
    s.Fields = &v
    return s
}
func (s *TaobaoTbkShopRecommendGetRequest) SetPlatform(v int64) *TaobaoTbkShopRecommendGetRequest {
    s.Platform = &v
    return s
}
func (s *TaobaoTbkShopRecommendGetRequest) SetUserId(v int64) *TaobaoTbkShopRecommendGetRequest {
    s.UserId = &v
    return s
}

func (req *TaobaoTbkShopRecommendGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Count != nil) {
        paramMap["count"] = *req.Count
    }
    if(req.Fields != nil) {
        paramMap["fields"] = *req.Fields
    }
    if(req.Platform != nil) {
        paramMap["platform"] = *req.Platform
    }
    if(req.UserId != nil) {
        paramMap["user_id"] = *req.UserId
    }
    return paramMap
}

func (req *TaobaoTbkShopRecommendGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}