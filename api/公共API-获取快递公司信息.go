package api

import "encoding/json"

// 文档：https://open.xiaohongshu.com/document/api?apiNavigationId=2&id=16&gatewayId=103&gatewayVersionId=1661&apiId=5746

type ReqGetExpressCompanyList struct {
	BaseRequest
}

var _ Request = new(ReqGetExpressCompanyList)

func (r ReqGetExpressCompanyList) Method() string {
	return "common.getExpressCompanyList"
}

func (r ReqGetExpressCompanyList) Params(base BaseRequest) []byte {
	r.BaseRequest = base
	ret, _ := json.Marshal(r)
	return ret
}

type (
	RespGetExpressCompanyList struct {
		BaseResponse
		Data RespGetExpressCompanyListData `json:"data"`
	}
	RespGetExpressCompanyListData struct {
		ExpressCompanyInfos []struct {
			ExpressCompanyId   int64  `json:"expressCompanyId"`
			ExpressCompanyCode string `json:"expressCompanyCode"`
			ExpressCompanyName string `json:"expressCompanyName"`
			Comment            string `json:"comment"`
		} `json:"expressCompanyInfos"`
	}
)

var _ Response = new(RespGetExpressCompanyList)

func (r RespGetExpressCompanyList) ErrorCode() int64 {
	return r.BaseResponse.ErrorCode
}

func (r RespGetExpressCompanyList) ErrorMsg() string {
	return r.BaseResponse.ErrorMsg

}

func (r RespGetExpressCompanyList) Success() bool {
	return r.BaseResponse.Success
}
