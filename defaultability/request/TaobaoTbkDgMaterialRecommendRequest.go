package request


type TaobaoTbkDgMaterialRecommendRequest struct {
    /*
        页大小，默认20，1~100 defalutValue��20    */
    PageSize  *int64 `json:"page_size,omitempty" required:"false" `
    /*
        第几页，默认：1 defalutValue��1    */
    PageNo  *int64 `json:"page_no,omitempty" required:"false" `
    /*
        官方提供的物料Id；可以通过taobao.tbk.optimus.tou.material.ids.get接口获取公开的物料id列表或查看物料id汇总贴：https://market.m.taobao.com/app/qn/toutiao-new/index-pc.html#/detail/10628875?_k=gpov9a     */
    MaterialId  *int64 `json:"material_id" required:"true" `
    /*
        推广位id，mm_xxx_xxx_12345678三段式的最后一段数字（登录pub.alimama.com推广管理-推广位管理中查询）     */
    AdzoneId  *int64 `json:"adzone_id" required:"true" `
    /*
        渠道关系ID，仅适用于渠道推广场景     */
    RelationId  *int64 `json:"relation_id,omitempty" required:"false" `
    /*
        智能匹配-设备号类型：IMEI，或者IDFA，或者UTDID（UTDID不支持MD5加密），或者OAID；使用智能推荐请先签署协议https://pub.alimama.com/fourth/protocol/common.htm?key=hangye_laxin     */
    DeviceType  *string `json:"device_type,omitempty" required:"false" `
    /*
        智能匹配-设备号加密类型：MD5；使用智能推荐请先签署协议https://pub.alimama.com/fourth/protocol/common.htm?key=hangye_laxin     */
    DeviceEncrypt  *string `json:"device_encrypt,omitempty" required:"false" `
    /*
        智能匹配-设备号加密后的值（MD5加密需32位小写）；使用智能推荐请先签署协议https://pub.alimama.com/fourth/protocol/common.htm?key=hangye_laxin     */
    DeviceValue  *string `json:"device_value,omitempty" required:"false" `
    /*
        1-自购省，2-推广赚（代理模式专属ID，代理模式必填，非代理模式不用填写该字段）     */
    PromotionType  *string `json:"promotion_type,omitempty" required:"false" `
    /*
        会员运营ID     */
    SpecialId  *string `json:"special_id,omitempty" required:"false" `
    /*
        淘宝客新商品ID；用于相似商品推荐（注意：使用相似商品推荐material_id=13256必传，并请先签署协议https://pub.alimama.com/fourth/protocol/common.htm?key=hangye_laxin)；另关于《淘宝客新商品ID升级》参考白皮书：https://www.yuque.com/taobaolianmengguanfangxiaoer/zmig94/tfyt0pahmlpzu2ud     */
    ItemId  *string `json:"item_id,omitempty" required:"false" `
    /*
        选品库收藏夹id，获取收藏夹id参考文档：https://mos.m.taobao.com/union/page_20230109_175050_176?spm=a219t._portal_v2_pages_promo_goods_index_htm.0.0.7c2a75a5H2ER3N     */
    FavoritesId  *string `json:"favorites_id,omitempty" required:"false" `
}

func (s *TaobaoTbkDgMaterialRecommendRequest) SetPageSize(v int64) *TaobaoTbkDgMaterialRecommendRequest {
    s.PageSize = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendRequest) SetPageNo(v int64) *TaobaoTbkDgMaterialRecommendRequest {
    s.PageNo = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendRequest) SetMaterialId(v int64) *TaobaoTbkDgMaterialRecommendRequest {
    s.MaterialId = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendRequest) SetAdzoneId(v int64) *TaobaoTbkDgMaterialRecommendRequest {
    s.AdzoneId = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendRequest) SetRelationId(v int64) *TaobaoTbkDgMaterialRecommendRequest {
    s.RelationId = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendRequest) SetDeviceType(v string) *TaobaoTbkDgMaterialRecommendRequest {
    s.DeviceType = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendRequest) SetDeviceEncrypt(v string) *TaobaoTbkDgMaterialRecommendRequest {
    s.DeviceEncrypt = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendRequest) SetDeviceValue(v string) *TaobaoTbkDgMaterialRecommendRequest {
    s.DeviceValue = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendRequest) SetPromotionType(v string) *TaobaoTbkDgMaterialRecommendRequest {
    s.PromotionType = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendRequest) SetSpecialId(v string) *TaobaoTbkDgMaterialRecommendRequest {
    s.SpecialId = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendRequest) SetItemId(v string) *TaobaoTbkDgMaterialRecommendRequest {
    s.ItemId = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendRequest) SetFavoritesId(v string) *TaobaoTbkDgMaterialRecommendRequest {
    s.FavoritesId = &v
    return s
}

func (req *TaobaoTbkDgMaterialRecommendRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.PageNo != nil) {
        paramMap["page_no"] = *req.PageNo
    }
    if(req.MaterialId != nil) {
        paramMap["material_id"] = *req.MaterialId
    }
    if(req.AdzoneId != nil) {
        paramMap["adzone_id"] = *req.AdzoneId
    }
    if(req.RelationId != nil) {
        paramMap["relation_id"] = *req.RelationId
    }
    if(req.DeviceType != nil) {
        paramMap["device_type"] = *req.DeviceType
    }
    if(req.DeviceEncrypt != nil) {
        paramMap["device_encrypt"] = *req.DeviceEncrypt
    }
    if(req.DeviceValue != nil) {
        paramMap["device_value"] = *req.DeviceValue
    }
    if(req.PromotionType != nil) {
        paramMap["promotion_type"] = *req.PromotionType
    }
    if(req.SpecialId != nil) {
        paramMap["special_id"] = *req.SpecialId
    }
    if(req.ItemId != nil) {
        paramMap["item_id"] = *req.ItemId
    }
    if(req.FavoritesId != nil) {
        paramMap["favorites_id"] = *req.FavoritesId
    }
    return paramMap
}

func (req *TaobaoTbkDgMaterialRecommendRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}