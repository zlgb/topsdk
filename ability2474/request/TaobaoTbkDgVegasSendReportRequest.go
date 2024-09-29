package request


type TaobaoTbkDgVegasSendReportRequest struct {
    /*
        统计日期     */
    BizDate  *string `json:"biz_date" required:"true" `
    /*
        渠道关系id     */
    RelationId  *int64 `json:"relation_id,omitempty" required:"false" `
    /*
        已下线，后续不需要填写     */
    ActivityId  *int64 `json:"activity_id,omitempty" required:"false" `
    /*
        页码 defalutValue��1    */
    PageNo  *int64 `json:"page_no,omitempty" required:"false" `
    /*
        每页大小 defalutValue��10    */
    PageSize  *int64 `json:"page_size,omitempty" required:"false" `
    /*
        媒体推广pid     */
    Pid  *string `json:"pid,omitempty" required:"false" `
    /*
        查询维度，不填写默认是pid维度     */
    RptDim  *string `json:"rpt_dim,omitempty" required:"false" `
    /*
        查询红包类型，1-超级红包，2-福利购，3-签到红包，4-福利直降，5-幸运赢免单，不传时默认查询超级红包数据 defalutValue��1    */
    ActivityCategory  *int64 `json:"activity_category,omitempty" required:"false" `
}

func (s *TaobaoTbkDgVegasSendReportRequest) SetBizDate(v string) *TaobaoTbkDgVegasSendReportRequest {
    s.BizDate = &v
    return s
}
func (s *TaobaoTbkDgVegasSendReportRequest) SetRelationId(v int64) *TaobaoTbkDgVegasSendReportRequest {
    s.RelationId = &v
    return s
}
func (s *TaobaoTbkDgVegasSendReportRequest) SetActivityId(v int64) *TaobaoTbkDgVegasSendReportRequest {
    s.ActivityId = &v
    return s
}
func (s *TaobaoTbkDgVegasSendReportRequest) SetPageNo(v int64) *TaobaoTbkDgVegasSendReportRequest {
    s.PageNo = &v
    return s
}
func (s *TaobaoTbkDgVegasSendReportRequest) SetPageSize(v int64) *TaobaoTbkDgVegasSendReportRequest {
    s.PageSize = &v
    return s
}
func (s *TaobaoTbkDgVegasSendReportRequest) SetPid(v string) *TaobaoTbkDgVegasSendReportRequest {
    s.Pid = &v
    return s
}
func (s *TaobaoTbkDgVegasSendReportRequest) SetRptDim(v string) *TaobaoTbkDgVegasSendReportRequest {
    s.RptDim = &v
    return s
}
func (s *TaobaoTbkDgVegasSendReportRequest) SetActivityCategory(v int64) *TaobaoTbkDgVegasSendReportRequest {
    s.ActivityCategory = &v
    return s
}

func (req *TaobaoTbkDgVegasSendReportRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.BizDate != nil) {
        paramMap["biz_date"] = *req.BizDate
    }
    if(req.RelationId != nil) {
        paramMap["relation_id"] = *req.RelationId
    }
    if(req.ActivityId != nil) {
        paramMap["activity_id"] = *req.ActivityId
    }
    if(req.PageNo != nil) {
        paramMap["page_no"] = *req.PageNo
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.Pid != nil) {
        paramMap["pid"] = *req.Pid
    }
    if(req.RptDim != nil) {
        paramMap["rpt_dim"] = *req.RptDim
    }
    if(req.ActivityCategory != nil) {
        paramMap["activity_category"] = *req.ActivityCategory
    }
    return paramMap
}

func (req *TaobaoTbkDgVegasSendReportRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}