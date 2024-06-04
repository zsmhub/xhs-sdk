package api

import "encoding/json"

type (
	ReqListAfterSaleInfos struct {
		BaseRequest
		PageNo      int64   `json:"pageNo"`
		PageSize    int64   `json:"pageSize"`
		OrderId     string  `json:"orderId"`
		Statuses    []int64 `json:"statuses"`
		ReturnTypes []int64 `json:"returnTypes"`
		StartTime   int64   `json:"startTime"`
		EndTime     int64   `json:"endTime"`
		TimeType    int64   `json:"timeType"`
	}
)

var _ Request = new(ReqListAfterSaleInfos)

func (r ReqListAfterSaleInfos) Method() string {
	return "afterSale.listAfterSaleInfos"
}

func (r ReqListAfterSaleInfos) Params(base BaseRequest) []byte {
	r.BaseRequest = base
	ret, _ := json.Marshal(r)
	return ret
}

type (
	RespListAfterSaleInfos struct {
		BaseResponse
		Data RespListAfterSaleInfosData `json:"data"`
	}
	RespListAfterSaleInfosData struct {
		TotalCount          int64 `json:"totalCount"`
		PageNo              int64 `json:"pageNo"`
		PageSize            int64 `json:"pageSize"`
		AfterSaleBasicInfos []struct {
			ReturnsId                string `json:"returnsId"`
			ReturnType               int64  `json:"returnType"`
			ReasonId                 int64  `json:"reasonId"`
			ReasonNameZh             string `json:"reasonNameZh"`
			Status                   int64  `json:"status"`
			UserId                   string `json:"userId"`
			OrderId                  string `json:"orderId"`
			ApplyTime                int64  `json:"applyTime"`
			UpdatedAt                int64  `json:"updatedAt"`
			ExpireTime               int64  `json:"expireTime"`
			Desc                     string `json:"desc"`
			ReturnsTag               int64  `json:"returnsTag"`
			ExpectedRefundAmountYuan int64  `json:"expectedRefundAmountYuan"`
		} `json:"afterSaleBasicInfos"`
		MaxPageNo int64 `json:"maxPageNo"`
	}
)

var _ Response = new(RespListAfterSaleInfos)

func (r RespListAfterSaleInfos) ErrorCode() int64 {
	return r.BaseResponse.ErrorCode
}

func (r RespListAfterSaleInfos) ErrorMsg() string {
	return r.BaseResponse.ErrorMsg

}

func (r RespListAfterSaleInfos) Success() bool {
	return r.BaseResponse.Success
}
