package api

import "encoding/json"

type (
	ReqAfterSaleConfirmReceive struct {
		BaseRequest
		ReturnsId   string `json:"returnsId"`
		Action      int64  `json:"action"`
		Reason      int64  `json:"reason,omitempty"`
		Description string `json:"description,omitempty"`
	}
)

var _ Request = new(ReqAfterSaleConfirmReceive)

func (r ReqAfterSaleConfirmReceive) Method() string {
	return "afterSale.confirmReceive"
}
func (r ReqAfterSaleConfirmReceive) Params(base BaseRequest) []byte {
	r.BaseRequest = base
	ret, _ := json.Marshal(r)
	return ret
}

type (
	RespAfterSaleConfirmReceive struct {
		BaseResponse
	}
)

var _ Response = new(RespAfterSaleConfirmReceive)

func (r RespAfterSaleConfirmReceive) ErrorCode() int64 {
	return r.BaseResponse.ErrorCode
}
func (r RespAfterSaleConfirmReceive) ErrorMsg() string {
	return r.BaseResponse.ErrorMsg
}
func (r RespAfterSaleConfirmReceive) Success() bool {
	return r.BaseResponse.Success
}
