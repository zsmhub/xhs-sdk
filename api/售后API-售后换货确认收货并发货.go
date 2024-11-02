package api

import "encoding/json"

type (
	ReqAfterSaleReceiverAndShip struct {
		BaseRequest
		ReturnsId          string `json:"returnsId"`
		ExpressCompanyCode string `json:"expressCompanyCode"`
		ExpressNo          string `json:"expressNo"`
	}
)

var _ Request = new(ReqAfterSaleReceiverAndShip)

func (r ReqAfterSaleReceiverAndShip) Method() string {
	return "afterSale.receiveAndShip"
}
func (r ReqAfterSaleReceiverAndShip) Params(base BaseRequest) []byte {
	r.BaseRequest = base
	ret, _ := json.Marshal(r)
	return ret
}

type (
	RespAfterSaleReceiverAndShip struct {
		BaseResponse
	}
)

var _ Response = new(RespAfterSaleReceiverAndShip)

func (r RespAfterSaleReceiverAndShip) ErrorCode() int64 {
	return r.BaseResponse.ErrorCode
}
func (r RespAfterSaleReceiverAndShip) ErrorMsg() string {
	return r.BaseResponse.ErrorMsg
}
func (r RespAfterSaleReceiverAndShip) Success() bool {
	return r.BaseResponse.Success
}
