package response

import (
    "topsdk/ability865/domain"
)

type TaobaoOpenAccountTokenValidateResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        验证成功返回token中的信息
    */
    Data  domain.TaobaoOpenAccountTokenValidateOpenAccountTokenValidateResult `json:"data,omitempty" `
}
