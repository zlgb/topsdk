package domain


type TaobaoTbkDgVegasSendReportRightsSendRptDTO struct {
    /*
        渠道关系id的发放数据     */
    RelationRptList  *[]TaobaoTbkDgVegasSendReportRightsSendRelationRptDto `json:"relation_rpt_list,omitempty" `

    /*
        pid的发放数据     */
    PidRptList  *[]TaobaoTbkDgVegasSendReportRightsSendRelationRptDto `json:"pid_rpt_list,omitempty" `

}

func (s *TaobaoTbkDgVegasSendReportRightsSendRptDTO) SetRelationRptList(v []TaobaoTbkDgVegasSendReportRightsSendRelationRptDto) *TaobaoTbkDgVegasSendReportRightsSendRptDTO {
    s.RelationRptList = &v
    return s
}
func (s *TaobaoTbkDgVegasSendReportRightsSendRptDTO) SetPidRptList(v []TaobaoTbkDgVegasSendReportRightsSendRelationRptDto) *TaobaoTbkDgVegasSendReportRightsSendRptDTO {
    s.PidRptList = &v
    return s
}
