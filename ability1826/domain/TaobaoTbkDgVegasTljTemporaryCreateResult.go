package domain


type TaobaoTbkDgVegasTljTemporaryCreateResult struct {
    /*
        model     */
    Model  *TaobaoTbkDgVegasTljTemporaryCreateRightsInstanceCreateResult `json:"model,omitempty" `

    /*
        msgCode     */
    MsgCode  *string `json:"msg_code,omitempty" `

    /*
        msgInfo     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        success     */
    Success  *bool `json:"success,omitempty" `

}

func (s *TaobaoTbkDgVegasTljTemporaryCreateResult) SetModel(v TaobaoTbkDgVegasTljTemporaryCreateRightsInstanceCreateResult) *TaobaoTbkDgVegasTljTemporaryCreateResult {
    s.Model = &v
    return s
}
func (s *TaobaoTbkDgVegasTljTemporaryCreateResult) SetMsgCode(v string) *TaobaoTbkDgVegasTljTemporaryCreateResult {
    s.MsgCode = &v
    return s
}
func (s *TaobaoTbkDgVegasTljTemporaryCreateResult) SetMsgInfo(v string) *TaobaoTbkDgVegasTljTemporaryCreateResult {
    s.MsgInfo = &v
    return s
}
func (s *TaobaoTbkDgVegasTljTemporaryCreateResult) SetSuccess(v bool) *TaobaoTbkDgVegasTljTemporaryCreateResult {
    s.Success = &v
    return s
}
