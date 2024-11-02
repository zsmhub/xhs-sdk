package api

import "encoding/json"

type (
	ReqAfterSaleRejectReasons struct {
		BaseRequest
		ReturnsId string `json:"returnsId"`
		//拒绝原因类型 1：审核拒绝 2：收货拒绝
		RejectReasonType int64 `json:"rejectReasonType"`
	}
)

var _ Request = new(ReqAfterSaleRejectReasons)

func (r ReqAfterSaleRejectReasons) Method() string {
	return "afterSale.rejectReasons"
}
func (r ReqAfterSaleRejectReasons) Params(base BaseRequest) []byte {
	r.BaseRequest = base
	ret, _ := json.Marshal(r)
	return ret
}

type (
	RespAfterSaleRejectReasons struct {
		BaseResponse
		Data []RespAfterSaleRejectReasonsItem `json:"data"`
	}

	RespAfterSaleRejectReasonsItem struct {
		ReasonType int64  `json:"reasonType"`
		ReasonId   int64  `json:"reasonId"`
		ReasonName string `json:"reasonName"`
	}
)

var _ Response = new(RespAfterSaleRejectReasons)

func (r RespAfterSaleRejectReasons) ErrorCode() int64 {
	return r.BaseResponse.ErrorCode
}
func (r RespAfterSaleRejectReasons) ErrorMsg() string {
	return r.BaseResponse.ErrorMsg
}
func (r RespAfterSaleRejectReasons) Success() bool {
	return r.BaseResponse.Success
}
