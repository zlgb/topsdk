package response

import (
    "topsdk/ability414/domain"
)

type TaobaoTbkOrderRefundGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        返回数据
    */
    Data  domain.TaobaoTbkOrderRefundGetOrderPage `json:"data,omitempty" `
    /*
        接口返回值信息，跟rpc架构保持一致
    */
    ResultCode  int64 `json:"result_code,omitempty" `
    /*
        业务错误信息
    */
    BizErrorDesc  string `json:"biz_error_desc,omitempty" `
    /*
        业务错误码 101, 102,103
    */
    BizErrorCode  int64 `json:"biz_error_code,omitempty" `
    /*
        返回信息
    */
    ResultMsg  string `json:"result_msg,omitempty" `
}
