package request


type TaobaoTbkItemInfoUpgradeGetRequest struct {
    /*
        商品ID。多个用","分割，一次最多查询20个     */
    ItemId  *string `json:"item_id" required:"true" `
    /*
        ip地址，影响邮费获取，如果不传或者传入不准确，邮费无法精准提供     */
    Ip  *string `json:"ip,omitempty" required:"false" `
    /*
        1-动态ID转链场景，2-消费者比价场景，3-商品库导购场景（不填默认为1）；场景id使用说明参考《淘宝客新商品ID升级》白皮书：https://www.yuque.com/taobaolianmengguanfangxiaoer/zmig94/tfyt0pahmlpzu2ud     */
    BizSceneId  *string `json:"biz_scene_id,omitempty" required:"false" `
    /*
        1-自购省，2-推广赚（代理模式专属ID，代理模式必填，非代理模式不用填写该字段）     */
    PromotionType  *string `json:"promotion_type,omitempty" required:"false" `
    /*
        渠道关系ID，仅适用于渠道推广场景     */
    RelationId  *string `json:"relation_id,omitempty" required:"false" `
    /*
        商品库服务账户(场景id3权限对应的memberid）     */
    ManageItemPubId  *int64 `json:"manage_item_pub_id,omitempty" required:"false" `
    /*
        是否获取单品淘礼金剩余数量，0-否，1-是，默认否(仅开通淘礼金权限媒体可查)     */
    GetTljInfo  *int64 `json:"get_tlj_info,omitempty" required:"false" `
}

func (s *TaobaoTbkItemInfoUpgradeGetRequest) SetItemId(v string) *TaobaoTbkItemInfoUpgradeGetRequest {
    s.ItemId = &v
    return s
}
func (s *TaobaoTbkItemInfoUpgradeGetRequest) SetIp(v string) *TaobaoTbkItemInfoUpgradeGetRequest {
    s.Ip = &v
    return s
}
func (s *TaobaoTbkItemInfoUpgradeGetRequest) SetBizSceneId(v string) *TaobaoTbkItemInfoUpgradeGetRequest {
    s.BizSceneId = &v
    return s
}
func (s *TaobaoTbkItemInfoUpgradeGetRequest) SetPromotionType(v string) *TaobaoTbkItemInfoUpgradeGetRequest {
    s.PromotionType = &v
    return s
}
func (s *TaobaoTbkItemInfoUpgradeGetRequest) SetRelationId(v string) *TaobaoTbkItemInfoUpgradeGetRequest {
    s.RelationId = &v
    return s
}
func (s *TaobaoTbkItemInfoUpgradeGetRequest) SetManageItemPubId(v int64) *TaobaoTbkItemInfoUpgradeGetRequest {
    s.ManageItemPubId = &v
    return s
}
func (s *TaobaoTbkItemInfoUpgradeGetRequest) SetGetTljInfo(v int64) *TaobaoTbkItemInfoUpgradeGetRequest {
    s.GetTljInfo = &v
    return s
}

func (req *TaobaoTbkItemInfoUpgradeGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ItemId != nil) {
        paramMap["item_id"] = *req.ItemId
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
    if(req.GetTljInfo != nil) {
        paramMap["get_tlj_info"] = *req.GetTljInfo
    }
    return paramMap
}

func (req *TaobaoTbkItemInfoUpgradeGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}