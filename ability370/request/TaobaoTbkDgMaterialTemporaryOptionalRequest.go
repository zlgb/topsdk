package request

import (
        "topsdk/ability370/domain"
        "topsdk/util"
    )

type TaobaoTbkDgMaterialTemporaryOptionalRequest struct {
    /*
        商品筛选(特定媒体支持)-店铺dsr评分。筛选大于等于当前设置的店铺dsr评分的商品0-50000之间     */
    StartDsr  *int64 `json:"start_dsr,omitempty" required:"false" `
    /*
        页大小，默认20，1~100 defalutValue��20    */
    PageSize  *int64 `json:"page_size,omitempty" required:"false" `
    /*
        第几页，默认：１ defalutValue��1    */
    PageNo  *int64 `json:"page_no,omitempty" required:"false" `
    /*
        链接形式：1：PC，2：无线，默认：１ defalutValue��1    */
    Platform  *int64 `json:"platform,omitempty" required:"false" `
    /*
        商品筛选-淘客佣金比率上限。如：1234表示12.34%     */
    EndTkRate  *int64 `json:"end_tk_rate,omitempty" required:"false" `
    /*
        商品筛选-淘客佣金比率下限。如：1234表示12.34%     */
    StartTkRate  *int64 `json:"start_tk_rate,omitempty" required:"false" `
    /*
        商品筛选-折扣价范围上限。单位：元     */
    EndPrice  *int64 `json:"end_price,omitempty" required:"false" `
    /*
        商品筛选-折扣价范围下限。单位：元     */
    StartPrice  *int64 `json:"start_price,omitempty" required:"false" `
    /*
        商品筛选-是否海外商品。true表示属于海外商品，false或不设置表示不限     */
    IsOverseas  *bool `json:"is_overseas,omitempty" required:"false" `
    /*
        商品筛选-是否天猫商品。true表示属于天猫商品，false或不设置表示不限     */
    IsTmall  *bool `json:"is_tmall,omitempty" required:"false" `
    /*
        排序_des（降序），排序_asc（升序），销量（total_sales），淘客佣金比率（tk_rate）， 累计推广量（tk_total_sales），总支出佣金（tk_total_commi），价格（price），匹配分（match）     */
    Sort  *string `json:"sort,omitempty" required:"false" `
    /*
        商品筛选-所在地     */
    Itemloc  *string `json:"itemloc,omitempty" required:"false" `
    /*
        商品筛选-后台类目ID。用,分割，最大10个，该ID可以通过taobao.itemcats.get接口获取到     */
    Cat  *string `json:"cat,omitempty" required:"false" `
    /*
        商品筛选-查询词     */
    Q  *string `json:"q,omitempty" required:"false" `
    /*
        不传时默认物料id=2836；如果直接对消费者投放，可使用官方个性化算法优化的搜索物料id=17004 defalutValue��2836    */
    MaterialId  *int64 `json:"material_id,omitempty" required:"false" `
    /*
        优惠券筛选-是否有优惠券。true表示该商品有优惠券，false或不设置表示不限     */
    HasCoupon  *bool `json:"has_coupon,omitempty" required:"false" `
    /*
        ip参数影响邮费获取，如果不传或者传入不准确，邮费无法精准提供     */
    Ip  *string `json:"ip,omitempty" required:"false" `
    /*
        mm_xxx_xxx_12345678三段式的最后一段数字     */
    AdzoneId  *int64 `json:"adzone_id" required:"true" `
    /*
        商品筛选-是否包邮。true表示包邮，false或不设置表示不限     */
    NeedFreeShipment  *bool `json:"need_free_shipment,omitempty" required:"false" `
    /*
        商品筛选-是否加入消费者保障。true表示加入，false或不设置表示不限     */
    NeedPrepay  *bool `json:"need_prepay,omitempty" required:"false" `
    /*
        商品筛选(特定媒体支持)-成交转化是否高于行业均值。True表示大于等于，false或不设置表示不限     */
    IncludePayRate30  *bool `json:"include_pay_rate_30,omitempty" required:"false" `
    /*
        商品筛选-好评率是否高于行业均值。True表示大于等于，false或不设置表示不限     */
    IncludeGoodRate  *bool `json:"include_good_rate,omitempty" required:"false" `
    /*
        商品筛选(特定媒体支持)-退款率是否低于行业均值。True表示大于等于，false或不设置表示不限     */
    IncludeRfdRate  *bool `json:"include_rfd_rate,omitempty" required:"false" `
    /*
        商品筛选-牛皮癣程度。取值：1不限，2无，3轻微 defalutValue��1    */
    NpxLevel  *int64 `json:"npx_level,omitempty" required:"false" `
    /*
        智能匹配-设备号加密类型：MD5     */
    DeviceEncrypt  *string `json:"device_encrypt,omitempty" required:"false" `
    /*
        智能匹配-设备号加密后的值（MD5加密需32位小写）     */
    DeviceValue  *string `json:"device_value,omitempty" required:"false" `
    /*
        智能匹配-设备类型：IMEI/IDFA/UTDID/OAID，（IOS推荐使用IDFA，android等推荐使用OAID）     */
    DeviceType  *string `json:"device_type,omitempty" required:"false" `
    /*
        商品筛选-KA媒体淘客佣金比率上限。如：1234表示12.34%     */
    EndKaTkRate  *int64 `json:"end_ka_tk_rate,omitempty" required:"false" `
    /*
        商品筛选-KA媒体淘客佣金比率下限。如：1234表示12.34%     */
    StartKaTkRate  *int64 `json:"start_ka_tk_rate,omitempty" required:"false" `
    /*
        锁佣结束时间     */
    LockRateEndTime  *int64 `json:"lock_rate_end_time,omitempty" required:"false" `
    /*
        锁佣开始时间     */
    LockRateStartTime  *int64 `json:"lock_rate_start_time,omitempty" required:"false" `
    /*
        本地化业务入参-LBS信息-经度     */
    Longitude  *string `json:"longitude,omitempty" required:"false" `
    /*
        本地化业务入参-LBS信息-纬度     */
    Latitude  *string `json:"latitude,omitempty" required:"false" `
    /*
        本地化业务入参-LBS信息-国标城市码，仅支持单个请求，请求饿了么卡券物料时，该字段必填。 （详细城市ID见：https://mo.m.taobao.com/page_2020010315120200508）     */
    CityCode  *string `json:"city_code,omitempty" required:"false" `
    /*
        商家id，仅支持饿了么卡券商家ID，支持批量请求1-100以内，多个商家ID使用英文逗号分隔     */
    SellerIds  *string `json:"seller_ids,omitempty" required:"false" `
    /*
        会员运营ID     */
    SpecialId  *string `json:"special_id,omitempty" required:"false" `
    /*
        渠道关系ID，仅适用于渠道推广场景     */
    RelationId  *string `json:"relation_id,omitempty" required:"false" `
    /*
        本地化业务入参-分页唯一标识，非首页的请求必传，值为上一页返回结果中的page_result_key字段值     */
    PageResultKey  *string `json:"page_result_key,omitempty" required:"false" `
    /*
        人群ID，仅适用于物料评估场景material_id=41377     */
    UcrowdId  *int64 `json:"ucrowd_id,omitempty" required:"false" `
    /*
        物料评估-商品列表     */
    UcrowdRankItems  *[]domain.TaobaoTbkDgMaterialTemporaryOptionalUcrowdrankitems `json:"ucrowd_rank_items,omitempty" required:"false" `
    /*
        是否获取前N件佣金信息	0否，1是，其他值否     */
    GetTopnRate  *int64 `json:"get_topn_rate,omitempty" required:"false" `
}

func (s *TaobaoTbkDgMaterialTemporaryOptionalRequest) SetStartDsr(v int64) *TaobaoTbkDgMaterialTemporaryOptionalRequest {
    s.StartDsr = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalRequest) SetPageSize(v int64) *TaobaoTbkDgMaterialTemporaryOptionalRequest {
    s.PageSize = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalRequest) SetPageNo(v int64) *TaobaoTbkDgMaterialTemporaryOptionalRequest {
    s.PageNo = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalRequest) SetPlatform(v int64) *TaobaoTbkDgMaterialTemporaryOptionalRequest {
    s.Platform = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalRequest) SetEndTkRate(v int64) *TaobaoTbkDgMaterialTemporaryOptionalRequest {
    s.EndTkRate = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalRequest) SetStartTkRate(v int64) *TaobaoTbkDgMaterialTemporaryOptionalRequest {
    s.StartTkRate = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalRequest) SetEndPrice(v int64) *TaobaoTbkDgMaterialTemporaryOptionalRequest {
    s.EndPrice = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalRequest) SetStartPrice(v int64) *TaobaoTbkDgMaterialTemporaryOptionalRequest {
    s.StartPrice = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalRequest) SetIsOverseas(v bool) *TaobaoTbkDgMaterialTemporaryOptionalRequest {
    s.IsOverseas = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalRequest) SetIsTmall(v bool) *TaobaoTbkDgMaterialTemporaryOptionalRequest {
    s.IsTmall = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalRequest) SetSort(v string) *TaobaoTbkDgMaterialTemporaryOptionalRequest {
    s.Sort = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalRequest) SetItemloc(v string) *TaobaoTbkDgMaterialTemporaryOptionalRequest {
    s.Itemloc = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalRequest) SetCat(v string) *TaobaoTbkDgMaterialTemporaryOptionalRequest {
    s.Cat = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalRequest) SetQ(v string) *TaobaoTbkDgMaterialTemporaryOptionalRequest {
    s.Q = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalRequest) SetMaterialId(v int64) *TaobaoTbkDgMaterialTemporaryOptionalRequest {
    s.MaterialId = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalRequest) SetHasCoupon(v bool) *TaobaoTbkDgMaterialTemporaryOptionalRequest {
    s.HasCoupon = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalRequest) SetIp(v string) *TaobaoTbkDgMaterialTemporaryOptionalRequest {
    s.Ip = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalRequest) SetAdzoneId(v int64) *TaobaoTbkDgMaterialTemporaryOptionalRequest {
    s.AdzoneId = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalRequest) SetNeedFreeShipment(v bool) *TaobaoTbkDgMaterialTemporaryOptionalRequest {
    s.NeedFreeShipment = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalRequest) SetNeedPrepay(v bool) *TaobaoTbkDgMaterialTemporaryOptionalRequest {
    s.NeedPrepay = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalRequest) SetIncludePayRate30(v bool) *TaobaoTbkDgMaterialTemporaryOptionalRequest {
    s.IncludePayRate30 = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalRequest) SetIncludeGoodRate(v bool) *TaobaoTbkDgMaterialTemporaryOptionalRequest {
    s.IncludeGoodRate = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalRequest) SetIncludeRfdRate(v bool) *TaobaoTbkDgMaterialTemporaryOptionalRequest {
    s.IncludeRfdRate = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalRequest) SetNpxLevel(v int64) *TaobaoTbkDgMaterialTemporaryOptionalRequest {
    s.NpxLevel = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalRequest) SetDeviceEncrypt(v string) *TaobaoTbkDgMaterialTemporaryOptionalRequest {
    s.DeviceEncrypt = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalRequest) SetDeviceValue(v string) *TaobaoTbkDgMaterialTemporaryOptionalRequest {
    s.DeviceValue = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalRequest) SetDeviceType(v string) *TaobaoTbkDgMaterialTemporaryOptionalRequest {
    s.DeviceType = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalRequest) SetEndKaTkRate(v int64) *TaobaoTbkDgMaterialTemporaryOptionalRequest {
    s.EndKaTkRate = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalRequest) SetStartKaTkRate(v int64) *TaobaoTbkDgMaterialTemporaryOptionalRequest {
    s.StartKaTkRate = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalRequest) SetLockRateEndTime(v int64) *TaobaoTbkDgMaterialTemporaryOptionalRequest {
    s.LockRateEndTime = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalRequest) SetLockRateStartTime(v int64) *TaobaoTbkDgMaterialTemporaryOptionalRequest {
    s.LockRateStartTime = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalRequest) SetLongitude(v string) *TaobaoTbkDgMaterialTemporaryOptionalRequest {
    s.Longitude = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalRequest) SetLatitude(v string) *TaobaoTbkDgMaterialTemporaryOptionalRequest {
    s.Latitude = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalRequest) SetCityCode(v string) *TaobaoTbkDgMaterialTemporaryOptionalRequest {
    s.CityCode = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalRequest) SetSellerIds(v string) *TaobaoTbkDgMaterialTemporaryOptionalRequest {
    s.SellerIds = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalRequest) SetSpecialId(v string) *TaobaoTbkDgMaterialTemporaryOptionalRequest {
    s.SpecialId = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalRequest) SetRelationId(v string) *TaobaoTbkDgMaterialTemporaryOptionalRequest {
    s.RelationId = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalRequest) SetPageResultKey(v string) *TaobaoTbkDgMaterialTemporaryOptionalRequest {
    s.PageResultKey = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalRequest) SetUcrowdId(v int64) *TaobaoTbkDgMaterialTemporaryOptionalRequest {
    s.UcrowdId = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalRequest) SetUcrowdRankItems(v []domain.TaobaoTbkDgMaterialTemporaryOptionalUcrowdrankitems) *TaobaoTbkDgMaterialTemporaryOptionalRequest {
    s.UcrowdRankItems = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalRequest) SetGetTopnRate(v int64) *TaobaoTbkDgMaterialTemporaryOptionalRequest {
    s.GetTopnRate = &v
    return s
}

func (req *TaobaoTbkDgMaterialTemporaryOptionalRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.StartDsr != nil) {
        paramMap["start_dsr"] = *req.StartDsr
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.PageNo != nil) {
        paramMap["page_no"] = *req.PageNo
    }
    if(req.Platform != nil) {
        paramMap["platform"] = *req.Platform
    }
    if(req.EndTkRate != nil) {
        paramMap["end_tk_rate"] = *req.EndTkRate
    }
    if(req.StartTkRate != nil) {
        paramMap["start_tk_rate"] = *req.StartTkRate
    }
    if(req.EndPrice != nil) {
        paramMap["end_price"] = *req.EndPrice
    }
    if(req.StartPrice != nil) {
        paramMap["start_price"] = *req.StartPrice
    }
    if(req.IsOverseas != nil) {
        paramMap["is_overseas"] = *req.IsOverseas
    }
    if(req.IsTmall != nil) {
        paramMap["is_tmall"] = *req.IsTmall
    }
    if(req.Sort != nil) {
        paramMap["sort"] = *req.Sort
    }
    if(req.Itemloc != nil) {
        paramMap["itemloc"] = *req.Itemloc
    }
    if(req.Cat != nil) {
        paramMap["cat"] = *req.Cat
    }
    if(req.Q != nil) {
        paramMap["q"] = *req.Q
    }
    if(req.MaterialId != nil) {
        paramMap["material_id"] = *req.MaterialId
    }
    if(req.HasCoupon != nil) {
        paramMap["has_coupon"] = *req.HasCoupon
    }
    if(req.Ip != nil) {
        paramMap["ip"] = *req.Ip
    }
    if(req.AdzoneId != nil) {
        paramMap["adzone_id"] = *req.AdzoneId
    }
    if(req.NeedFreeShipment != nil) {
        paramMap["need_free_shipment"] = *req.NeedFreeShipment
    }
    if(req.NeedPrepay != nil) {
        paramMap["need_prepay"] = *req.NeedPrepay
    }
    if(req.IncludePayRate30 != nil) {
        paramMap["include_pay_rate_30"] = *req.IncludePayRate30
    }
    if(req.IncludeGoodRate != nil) {
        paramMap["include_good_rate"] = *req.IncludeGoodRate
    }
    if(req.IncludeRfdRate != nil) {
        paramMap["include_rfd_rate"] = *req.IncludeRfdRate
    }
    if(req.NpxLevel != nil) {
        paramMap["npx_level"] = *req.NpxLevel
    }
    if(req.DeviceEncrypt != nil) {
        paramMap["device_encrypt"] = *req.DeviceEncrypt
    }
    if(req.DeviceValue != nil) {
        paramMap["device_value"] = *req.DeviceValue
    }
    if(req.DeviceType != nil) {
        paramMap["device_type"] = *req.DeviceType
    }
    if(req.EndKaTkRate != nil) {
        paramMap["end_ka_tk_rate"] = *req.EndKaTkRate
    }
    if(req.StartKaTkRate != nil) {
        paramMap["start_ka_tk_rate"] = *req.StartKaTkRate
    }
    if(req.LockRateEndTime != nil) {
        paramMap["lock_rate_end_time"] = *req.LockRateEndTime
    }
    if(req.LockRateStartTime != nil) {
        paramMap["lock_rate_start_time"] = *req.LockRateStartTime
    }
    if(req.Longitude != nil) {
        paramMap["longitude"] = *req.Longitude
    }
    if(req.Latitude != nil) {
        paramMap["latitude"] = *req.Latitude
    }
    if(req.CityCode != nil) {
        paramMap["city_code"] = *req.CityCode
    }
    if(req.SellerIds != nil) {
        paramMap["seller_ids"] = *req.SellerIds
    }
    if(req.SpecialId != nil) {
        paramMap["special_id"] = *req.SpecialId
    }
    if(req.RelationId != nil) {
        paramMap["relation_id"] = *req.RelationId
    }
    if(req.PageResultKey != nil) {
        paramMap["page_result_key"] = *req.PageResultKey
    }
    if(req.UcrowdId != nil) {
        paramMap["ucrowd_id"] = *req.UcrowdId
    }
    if(req.UcrowdRankItems != nil) {
        paramMap["ucrowd_rank_items"] = util.ConvertStructList(*req.UcrowdRankItems)
    }
    if(req.GetTopnRate != nil) {
        paramMap["get_topn_rate"] = *req.GetTopnRate
    }
    return paramMap
}

func (req *TaobaoTbkDgMaterialTemporaryOptionalRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}