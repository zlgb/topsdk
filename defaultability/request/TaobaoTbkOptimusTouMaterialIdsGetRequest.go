package request

import (
        "topsdk/defaultability/domain"
        "topsdk/util"
    )

type TaobaoTbkOptimusTouMaterialIdsGetRequest struct {
    /*
        请求结构     */
    MaterialQuery  *domain.TaobaoTbkOptimusTouMaterialIdsGetMaterialQuery `json:"material_query,omitempty" required:"false" `
}

func (s *TaobaoTbkOptimusTouMaterialIdsGetRequest) SetMaterialQuery(v domain.TaobaoTbkOptimusTouMaterialIdsGetMaterialQuery) *TaobaoTbkOptimusTouMaterialIdsGetRequest {
    s.MaterialQuery = &v
    return s
}

func (req *TaobaoTbkOptimusTouMaterialIdsGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.MaterialQuery != nil) {
        paramMap["material_query"] = util.ConvertStruct(*req.MaterialQuery)
    }
    return paramMap
}

func (req *TaobaoTbkOptimusTouMaterialIdsGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}