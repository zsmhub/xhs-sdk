package api

import "encoding/json"

// 文档：https://open.xiaohongshu.com/document/api?apiNavigationId=3&id=6&gatewayId=103&gatewayVersionId=1661&apiId=27235

type (
	ReqOrderDeliver struct {
		BaseRequest
		OrderId            string   `json:"orderId"`
		ExpressNo          string   `json:"expressNo"`
		ExpressCompanyCode string   `json:"expressCompanyCode"`
		ExpressCompanyName string   `json:"expressCompanyName"`
		DeliveringTime     int64    `json:"deliveringTime"`
		Unpack             bool     `json:"unpack"`
		SkuIdList          []string `json:"skuIdList"`
	}
)

var _ Request = new(ReqOrderDeliver)

func (r ReqOrderDeliver) Method() string {
	return "order.orderDeliver"
}

func (r ReqOrderDeliver) Params(base BaseRequest) []byte {
	r.BaseRequest = base
	ret, _ := json.Marshal(r)
	return ret
}

type (
	RespOrderDeliver struct {
		BaseResponse
		Data string `json:"data"`
	}
)

var _ Response = new(RespOrderDeliver)

func (r RespOrderDeliver) ErrorCode() int64 {
	return r.BaseResponse.ErrorCode
}

func (r RespOrderDeliver) ErrorMsg() string {
	return r.BaseResponse.ErrorMsg
}

func (r RespOrderDeliver) Success() bool {
	return r.BaseResponse.Success
}
