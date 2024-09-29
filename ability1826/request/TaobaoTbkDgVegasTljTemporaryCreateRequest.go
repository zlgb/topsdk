package request

import (
        "topsdk/util"
    )

type TaobaoTbkDgVegasTljTemporaryCreateRequest struct {
    /*
        妈妈广告位Id     */
    AdzoneId  *int64 `json:"adzone_id" required:"true" `
    /*
        必须传入0     */
    SecurityLevel  *int64 `json:"security_level,omitempty" required:"false" `
    /*
        使用开始日期。相对时间，无需填写，以用户领取时间作为使用开始时间。绝对时间，格式 yyyy-MM-dd，例如，2019-01-29，表示从2019-01-29 00:00:00 开始     */
    UseStartTime  *string `json:"use_start_time,omitempty" required:"false" `
    /*
        结束日期的模式,1:相对时间，2:绝对时间     */
    UseEndTimeMode  *int64 `json:"use_end_time_mode,omitempty" required:"false" `
    /*
        使用结束日期。如果是结束时间模式为相对时间，时间格式为1-7直接的整数, 例如，1（相对领取时间1天）； 如果是绝对时间，格式为yyyy-MM-dd，例如，2019-01-29，表示到2019-01-29 23:59:59结束     */
    UseEndTime  *string `json:"use_end_time,omitempty" required:"false" `
    /*
        发放截止时间     */
    SendEndTime  *util.LocalTime `json:"send_end_time,omitempty" required:"false" `
    /*
        发放开始时间     */
    SendStartTime  *util.LocalTime `json:"send_start_time" required:"true" `
    /*
        单个淘礼金面额，支持两位小数，单位元     */
    PerFace  *string `json:"per_face" required:"true" `
    /*
        必须设置为true，默认开启安全     */
    SecuritySwitch  *bool `json:"security_switch" required:"true" `
    /*
        单用户累计中奖次数上限     */
    UserTotalWinNumLimit  *int64 `json:"user_total_win_num_limit" required:"true" `
    /*
        淘礼金名称，最大10个字符     */
    Name  *string `json:"name" required:"true" `
    /*
        淘礼金总个数     */
    TotalNum  *int64 `json:"total_num" required:"true" `
    /*
        宝贝ID或营销ID     */
    ItemId  *string `json:"item_id" required:"true" `
    /*
        CPS佣金类型 defalutValue��MKT    */
    CampaignType  *string `json:"campaign_type,omitempty" required:"false" `
}

func (s *TaobaoTbkDgVegasTljTemporaryCreateRequest) SetAdzoneId(v int64) *TaobaoTbkDgVegasTljTemporaryCreateRequest {
    s.AdzoneId = &v
    return s
}
func (s *TaobaoTbkDgVegasTljTemporaryCreateRequest) SetSecurityLevel(v int64) *TaobaoTbkDgVegasTljTemporaryCreateRequest {
    s.SecurityLevel = &v
    return s
}
func (s *TaobaoTbkDgVegasTljTemporaryCreateRequest) SetUseStartTime(v string) *TaobaoTbkDgVegasTljTemporaryCreateRequest {
    s.UseStartTime = &v
    return s
}
func (s *TaobaoTbkDgVegasTljTemporaryCreateRequest) SetUseEndTimeMode(v int64) *TaobaoTbkDgVegasTljTemporaryCreateRequest {
    s.UseEndTimeMode = &v
    return s
}
func (s *TaobaoTbkDgVegasTljTemporaryCreateRequest) SetUseEndTime(v string) *TaobaoTbkDgVegasTljTemporaryCreateRequest {
    s.UseEndTime = &v
    return s
}
func (s *TaobaoTbkDgVegasTljTemporaryCreateRequest) SetSendEndTime(v util.LocalTime) *TaobaoTbkDgVegasTljTemporaryCreateRequest {
    s.SendEndTime = &v
    return s
}
func (s *TaobaoTbkDgVegasTljTemporaryCreateRequest) SetSendStartTime(v util.LocalTime) *TaobaoTbkDgVegasTljTemporaryCreateRequest {
    s.SendStartTime = &v
    return s
}
func (s *TaobaoTbkDgVegasTljTemporaryCreateRequest) SetPerFace(v string) *TaobaoTbkDgVegasTljTemporaryCreateRequest {
    s.PerFace = &v
    return s
}
func (s *TaobaoTbkDgVegasTljTemporaryCreateRequest) SetSecuritySwitch(v bool) *TaobaoTbkDgVegasTljTemporaryCreateRequest {
    s.SecuritySwitch = &v
    return s
}
func (s *TaobaoTbkDgVegasTljTemporaryCreateRequest) SetUserTotalWinNumLimit(v int64) *TaobaoTbkDgVegasTljTemporaryCreateRequest {
    s.UserTotalWinNumLimit = &v
    return s
}
func (s *TaobaoTbkDgVegasTljTemporaryCreateRequest) SetName(v string) *TaobaoTbkDgVegasTljTemporaryCreateRequest {
    s.Name = &v
    return s
}
func (s *TaobaoTbkDgVegasTljTemporaryCreateRequest) SetTotalNum(v int64) *TaobaoTbkDgVegasTljTemporaryCreateRequest {
    s.TotalNum = &v
    return s
}
func (s *TaobaoTbkDgVegasTljTemporaryCreateRequest) SetItemId(v string) *TaobaoTbkDgVegasTljTemporaryCreateRequest {
    s.ItemId = &v
    return s
}
func (s *TaobaoTbkDgVegasTljTemporaryCreateRequest) SetCampaignType(v string) *TaobaoTbkDgVegasTljTemporaryCreateRequest {
    s.CampaignType = &v
    return s
}

func (req *TaobaoTbkDgVegasTljTemporaryCreateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.AdzoneId != nil) {
        paramMap["adzone_id"] = *req.AdzoneId
    }
    if(req.SecurityLevel != nil) {
        paramMap["security_level"] = *req.SecurityLevel
    }
    if(req.UseStartTime != nil) {
        paramMap["use_start_time"] = *req.UseStartTime
    }
    if(req.UseEndTimeMode != nil) {
        paramMap["use_end_time_mode"] = *req.UseEndTimeMode
    }
    if(req.UseEndTime != nil) {
        paramMap["use_end_time"] = *req.UseEndTime
    }
    if(req.SendEndTime != nil) {
        paramMap["send_end_time"] = *req.SendEndTime
    }
    if(req.SendStartTime != nil) {
        paramMap["send_start_time"] = *req.SendStartTime
    }
    if(req.PerFace != nil) {
        paramMap["per_face"] = *req.PerFace
    }
    if(req.SecuritySwitch != nil) {
        paramMap["security_switch"] = *req.SecuritySwitch
    }
    if(req.UserTotalWinNumLimit != nil) {
        paramMap["user_total_win_num_limit"] = *req.UserTotalWinNumLimit
    }
    if(req.Name != nil) {
        paramMap["name"] = *req.Name
    }
    if(req.TotalNum != nil) {
        paramMap["total_num"] = *req.TotalNum
    }
    if(req.ItemId != nil) {
        paramMap["item_id"] = *req.ItemId
    }
    if(req.CampaignType != nil) {
        paramMap["campaign_type"] = *req.CampaignType
    }
    return paramMap
}

func (req *TaobaoTbkDgVegasTljTemporaryCreateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}