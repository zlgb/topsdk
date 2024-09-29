package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoTbkDgGeneralLinkConvertResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        TbkLinkDTO
    */
    Data  domain.TaobaoTbkDgGeneralLinkConvertTbkLinkDTO `json:"data,omitempty" `
    /*
        见错误码描述
    */
    BizErrorDesc  int64 `json:"biz_error_desc,omitempty" `
    /*
        见错误码描述
    */
    ResultMsg  string `json:"result_msg,omitempty" `
}
