package request


type TaobaoTbkItemInfoGetRequest struct {
    /*
        商品ID串，用,分割，最大40个     */
    NumIids  *string `json:"num_iids" required:"true" `
    /*
        链接形式：1：PC，2：无线，默认：１ defalutValue��1    */
    Platform  *int64 `json:"platform,omitempty" required:"false" `
    /*
        ip地址，影响邮费获取，如果不传或者传入不准确，邮费无法精准提供     */
    Ip  *string `json:"ip,omitempty" required:"false" `
    /*
        1-动态ID转链场景，2-消费者比价场景，3-商品库导购场景（不填默认为1）     */
    BizSceneId  *string `json:"biz_scene_id,omitempty" required:"false" `
    /*
        1-自购省，2-推广赚（代理模式专属ID，代理模式必填，非代理模式不用填写该字段）     */
    PromotionType  *string `json:"promotion_type,omitempty" required:"false" `
    /*
        渠道关系ID     */
    RelationId  *string `json:"relation_id,omitempty" required:"false" `
    /*
        商品库服务账户(场景id3权限对应的memberid）     */
    ManageItemPubId  *int64 `json:"manage_item_pub_id,omitempty" required:"false" `
}

func (s *TaobaoTbkItemInfoGetRequest) SetNumIids(v string) *TaobaoTbkItemInfoGetRequest {
    s.NumIids = &v
    return s
}
func (s *TaobaoTbkItemInfoGetRequest) SetPlatform(v int64) *TaobaoTbkItemInfoGetRequest {
    s.Platform = &v
    return s
}
func (s *TaobaoTbkItemInfoGetRequest) SetIp(v string) *TaobaoTbkItemInfoGetRequest {
    s.Ip = &v
    return s
}
func (s *TaobaoTbkItemInfoGetRequest) SetBizSceneId(v string) *TaobaoTbkItemInfoGetRequest {
    s.BizSceneId = &v
    return s
}
func (s *TaobaoTbkItemInfoGetRequest) SetPromotionType(v string) *TaobaoTbkItemInfoGetRequest {
    s.PromotionType = &v
    return s
}
func (s *TaobaoTbkItemInfoGetRequest) SetRelationId(v string) *TaobaoTbkItemInfoGetRequest {
    s.RelationId = &v
    return s
}
func (s *TaobaoTbkItemInfoGetRequest) SetManageItemPubId(v int64) *TaobaoTbkItemInfoGetRequest {
    s.ManageItemPubId = &v
    return s
}

func (req *TaobaoTbkItemInfoGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.NumIids != nil) {
        paramMap["num_iids"] = *req.NumIids
    }
    if(req.Platform != nil) {
        paramMap["platform"] = *req.Platform
    }
    if(req.Ip != nil) {
        paramMap["ip"] = *req.Ip
    }
    if(req.BizSceneId != nil) {
        paramMap["biz_scene_id"] = *req.BizSceneId
    }
    if(req.PromotionType != nil) {
        paramMap["promotion_type"] = *req.PromotionType
    }
    if(req.RelationId != nil) {
        paramMap["relation_id"] = *req.RelationId
    }
    if(req.ManageItemPubId != nil) {
        paramMap["manage_item_pub_id"] = *req.ManageItemPubId
    }
    return paramMap
}

func (req *TaobaoTbkItemInfoGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}