package api

import "encoding/json"

type (
	ReqAfterSaleAuditReturns struct {
		BaseRequest
		ReqAfterSaleAuditReturnsData
	}

	ReqAfterSaleAuditReturnsReceiveInfo struct {
		Code                       string `json:"code"`
		Country                    string `json:"country"`
		City                       string `json:"city"`
		Province                   string `json:"province"`
		District                   string `json:"district"`
		Street                     string `json:"street"`
		ReceiverName               string `json:"receiverName"`
		ReceiverPhone              string `json:"receiverPhone"`
		SellerAddressRecordId      int    `json:"sellerAddressRecordId"`
		SellerAddressRecordVersion int    `json:"sellerAddressRecordVersion"`
	}

	ReqAfterSaleAuditReturnsData struct {
		ReturnsId    string                               `json:"returnsId"`
		Action       int                                  `json:"action"`
		Reason       int                                  `json:"reason,omitempty"`
		Description  string                               `json:"description,omitempty"`
		Message      string                               `json:"message,omitempty"`
		ReceiverInfo *ReqAfterSaleAuditReturnsReceiveInfo `json:"receiverInfo,omitempty"`
	}
)

var _ Request = new(ReqAfterSaleAuditReturns)

func (r ReqAfterSaleAuditReturns) Method() string {
	return "afterSale.auditReturns"
}
func (r ReqAfterSaleAuditReturns) Params(base BaseRequest) []byte {
	r.BaseRequest = base
	ret, _ := json.Marshal(r)
	return ret
}

type (
	RespAfterSaleAuditReturns struct {
		BaseResponse
	}
)

var _ Response = new(RespAfterSaleAuditReturns)

func (r RespAfterSaleAuditReturns) ErrorCode() int64 {
	return r.BaseResponse.ErrorCode
}
func (r RespAfterSaleAuditReturns) ErrorMsg() string {
	return r.BaseResponse.ErrorMsg
}
func (r RespAfterSaleAuditReturns) Success() bool {
	return r.BaseResponse.Success
}
