package api

import "encoding/json"

// 文档：https://open.xiaohongshu.com/document/api?apiNavigationId=4&id=18&gatewayId=103&gatewayVersionId=1661&apiId=5698

type (
	ReqListAfterSaleApi struct {
		BaseRequest
		Status     int64 `json:"status"`
		PageNo     int64 `json:"pageNo"`
		PageSize   int64 `json:"pageSize"`
		StartTime  int64 `json:"startTime"`
		EndTime    int64 `json:"endTime"`
		TimeType   int64 `json:"timeType"`
		UseHasNext bool  `json:"useHasNext"`
		ReasonId   int64 `json:"reasonId"`
		ReturnType int64 `json:"returnType"`
	}
)

var _ Request = new(ReqListAfterSaleApi)

func (r ReqListAfterSaleApi) Method() string {
	return "afterSale.listAfterSaleApi"
}

func (r ReqListAfterSaleApi) Params(base BaseRequest) []byte {
	r.BaseRequest = base
	ret, _ := json.Marshal(r)
	return ret
}

type (
	RespListAfterSaleApi struct {
		BaseResponse
		Data RespListAfterSaleApiData `json:"data"`
	}
	RespListAfterSaleApiData struct {
		Total               int64 `json:"total"`
		PageNo              int64 `json:"pageNo"`
		PageSize            int64 `json:"pageSize"`
		HaxNext             bool  `json:"haxNext"`
		SimpleAfterSaleList []struct {
			ReturnsId                string  `json:"returnsId"`
			ReturnType               int64   `json:"returnType"`
			ReasonId                 int64   `json:"reasonId"`
			Reason                   string  `json:"reason"`
			Status                   int64   `json:"status"`
			SubStatus                int64   `json:"subStatus"`
			ReceiveAbnormalType      int64   `json:"receiveAbnormalType"`
			OrderId                  string  `json:"orderId"`
			ExchangeOrderId          string  `json:"exchangeOrderId"`
			UserId                   string  `json:"userId"`
			CreatedTime              int64   `json:"createdTime"`
			ReturnExpressNo          string  `json:"returnExpressNo"`
			ReturnExpressCompany     string  `json:"returnExpressCompany"`
			ReturnAddress            string  `json:"returnAddress"`
			ShipNeeded               int64   `json:"shipNeeded"`
			Refunded                 bool    `json:"refunded"`
			RefundStatus             int64   `json:"refundStatus"`
			AutoReceiveDeadline      int64   `json:"autoReceiveDeadline"`
			UseFastRefund            bool    `json:"useFastRefund"`
			UpdateTime               int64   `json:"updateTime"`
			ReturnExpressCompanyCode string  `json:"returnExpressCompanyCode"`
			ExpectedRefundAmount     float64 `json:"expectedRefundAmount"`
		} `json:"simpleAfterSaleList"`
		MaxPageNo int64 `json:"maxPageNo"`
	}
)

var _ Response = new(RespListAfterSaleApi)

func (r RespListAfterSaleApi) ErrorCode() int64 {
	return r.BaseResponse.ErrorCode
}

func (r RespListAfterSaleApi) ErrorMsg() string {
	return r.BaseResponse.ErrorMsg

}

func (r RespListAfterSaleApi) Success() bool {
	return r.BaseResponse.Success
}
