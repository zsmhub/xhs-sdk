package api

import "encoding/json"

type (
	ReqGetOrderReceiverInfo struct {
		BaseRequest
		ReceiverQueries []ReceiverQuery `json:"receiverQueries"`
		IsReturn        bool            `json:"isReturn"`
	}
	ReceiverQuery struct {
		OrderId       string `json:"orderId"`
		OpenAddressId string `json:"openAddressId"`
	}
)

var _ Request = new(ReqGetOrderReceiverInfo)

func (r ReqGetOrderReceiverInfo) Method() string {
	return "order.getOrderReceiverInfo"
}

func (r ReqGetOrderReceiverInfo) Params(base BaseRequest) []byte {
	r.BaseRequest = base
	ret, _ := json.Marshal(r)
	return ret
}

type (
	RespGetOrderReceiverInfo struct {
		BaseResponse
		Data RespGetOrderReceiverInfoData `json:"data"`
	}
	RespGetOrderReceiverInfoData struct {
		ReceiverInfos []struct {
			OrderId              string `json:"orderId"`
			Matched              bool   `json:"matched"`
			ReceiverProvinceName string `json:"receiverProvinceName"`
			ReceiverCityName     string `json:"receiverCityName"`
			ReceiverDistrictName string `json:"receiverDistrictName"`
			ReceiverTownName     string `json:"receiverTownName"`
			ReceiverName         string `json:"receiverName"`
			ReceiverPhone        string `json:"receiverPhone"`
			ReceiverAddress      string `json:"receiverAddress"`
		} `json:"receiverInfos"`
	}
)

var _ Response = new(RespGetOrderReceiverInfo)

func (r RespGetOrderReceiverInfo) ErrorCode() int64 {
	return r.BaseResponse.ErrorCode
}

func (r RespGetOrderReceiverInfo) ErrorMsg() string {
	return r.BaseResponse.ErrorMsg

}

func (r RespGetOrderReceiverInfo) Success() bool {
	return r.BaseResponse.Success
}
