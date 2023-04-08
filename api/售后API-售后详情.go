package api

import "encoding/json"

// 文档：https://open.xiaohongshu.com/document/api?apiNavigationId=4&id=19&gatewayId=103&gatewayVersionId=1661&apiId=5699

type (
	ReqGetAfterSaleDetail struct {
		BaseRequest
		AfterSaleId string `json:"afterSaleId"`
	}
)

var _ Request = new(ReqGetAfterSaleDetail)

func (r ReqGetAfterSaleDetail) Method() string {
	return "afterSale.getAfterSaleDetail"
}

func (r ReqGetAfterSaleDetail) Params(base BaseRequest) []byte {
	r.BaseRequest = base
	ret, _ := json.Marshal(r)
	return ret
}

type (
	RespGetAfterSaleDetail struct {
		BaseResponse
		Data RespGetAfterSaleDetailData `json:"data"`
	}
	RespGetAfterSaleDetailData struct {
		ReturnsId            string   `json:"returnsId"`
		ReturnType           int64    `json:"returnType"`
		ReasonId             int64    `json:"reasonId"`
		Reason               string   `json:"reason"`
		Status               int64    `json:"status"`
		SubStatus            int64    `json:"subStatus"`
		ReceiveAbnormalType  int64    `json:"receiveAbnormalType"`
		OrderId              string   `json:"orderId"`
		ExchangeOrderId      string   `json:"exchangeOrderId"`
		UserId               string   `json:"userId"`
		CreatedAt            int64    `json:"createdAt"`
		ReturnExpressNo      string   `json:"returnExpressNo"`
		ReturnExpressCompany string   `json:"returnExpressCompany"`
		ReturnAddress        string   `json:"returnAddress"`
		ShipNeeded           int64    `json:"shipNeeded"`
		Refunded             bool     `json:"refunded"`
		RefundStatus         int64    `json:"refundStatus"`
		AutoReceiveDeadline  int64    `json:"autoReceiveDeadline"`
		UseFastRefund        bool     `json:"useFastRefund"`
		ProofPhotos          []string `json:"proofPhotos"`
		Desc                 string   `json:"desc"`
		Note                 string   `json:"note"`
		RefundTime           int64    `json:"refundTime"`
		FillExpressTime      int64    `json:"fillExpressTime"`
		ExpressSignTime      int64    `json:"expressSignTime"`
		Skus                 []struct {
			SkuId             string  `json:"skuId"`
			SkuName           string  `json:"skuName"`
			Image             string  `json:"image"`
			Price             float64 `json:"price"`
			BoughtCount       int64   `json:"boughtCount"`
			AppliedCount      int64   `json:"appliedCount"`
			ReturnedCount     int64   `json:"returnedCount"`
			RefundedCount     int64   `json:"refundedCount"`
			ReturnPrice       float64 `json:"returnPrice"`
			ExchangeSkuId     string  `json:"exchangeSkuId"`
			ExchangeSkuName   string  `json:"exchangeSkuName"`
			ExchangeSkuImage  string  `json:"exchangeSkuImage"`
			ScSkucode         string  `json:"scSkucode"`
			Barcode           string  `json:"barcode"`
			ExchangeScskuCode string  `json:"exchangeScskuCode"`
			ExchangeBarcode   string  `json:"exchangeBarcode"`
		} `json:"skus"`
		RefundFee                float64 `json:"refundFee"`
		ReturnExpressRefundable  bool    `json:"returnExpressRefundable"`
		ReturnExpressRefunded    bool    `json:"returnExpressRefunded"`
		ExpectRefundFee          float64 `json:"expectRefundFee"`
		UpdatedAt                int64   `json:"updatedAt"`
		ReturnExpressCompanyCode string  `json:"returnExpressCompanyCode"`
		OpenAddressId            string  `json:"openAddressId"`
		ExpectedRefundAmount     float64 `json:"expectedRefundAmount"`
	}
)

var _ Response = new(RespGetAfterSaleDetail)

func (r RespGetAfterSaleDetail) ErrorCode() int64 {
	return r.BaseResponse.ErrorCode
}

func (r RespGetAfterSaleDetail) ErrorMsg() string {
	return r.BaseResponse.ErrorMsg

}

func (r RespGetAfterSaleDetail) Success() bool {
	return r.BaseResponse.Success
}
