package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoTbkOptimusTouMaterialIdsGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        返回结果
    */
    Data  []domain.TaobaoTbkOptimusTouMaterialIdsGetTouMaterials `json:"data,omitempty" `
    /*
        物料id总数
    */
    TotalCount  int64 `json:"total_count,omitempty" `
    /*
        结果描述信息
    */
    ResultMsg  string `json:"result_msg,omitempty" `
}
