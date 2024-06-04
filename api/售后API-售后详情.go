package api

import "encoding/json"

type (
	ReqGetAfterSaleInfo struct {
		BaseRequest
		ReturnsId           string `json:"returnsId"`
		NeedNegotiateRecord bool   `json:"needNegotiateRecord"`
	}
)

var _ Request = new(ReqGetAfterSaleInfo)

func (r ReqGetAfterSaleInfo) Method() string {
	return "afterSale.getAfterSaleInfo"
}

func (r ReqGetAfterSaleInfo) Params(base BaseRequest) []byte {
	r.BaseRequest = base
	ret, _ := json.Marshal(r)
	return ret
}

type (
	RespGetAfterSaleInfo struct {
		BaseResponse
		Data RespGetAfterSaleInfoData `json:"data"`
	}

	RespGetAfterSaleInfoData struct {
		AfterSaleInfo    *RespGetAfterSaleInfoDataAfterSaleInfo    `json:"afterSaleInfo"`
		LogisticsInfo    *RespGetAfterSaleInfoDataLogisticsInfo    `json:"logisticsInfo"`
		NegotiateRecords *RespGetAfterSaleInfoDataNegotiateRecords `json:"negotiateRecords"`
	}

	RespGetAfterSaleInfoDataAfterSaleInfo struct {
		ReturnsId                string                      `json:"returnsId"`
		ReturnType               int64                       `json:"returnType"`
		ReasonId                 int64                       `json:"reasonId"`
		ReasonNameZh             string                      `json:"reasonNameZh"`
		Status                   int64                       `json:"status"`
		UserId                   string                      `json:"userId"`
		OrderId                  string                      `json:"orderId"`
		ApplyTime                int64                       `json:"applyTime"`
		UpdatedAt                int64                       `json:"updatedAt"`
		ExpireTime               int64                       `json:"expireTime"`
		ReturnAddress            *AfterSaleInfoReturnAddress `json:"return_address"`
		ProofPhotos              []string                    `json:"proofPhotos"`
		Desc                     string                      `json:"desc"`
		SupportCarriageInsurance bool                        `json:"supportCarriageInsurance"`
		OpenAddressId            string                      `json:"openAddressId"`
		Skus                     []AfterSaleInfoSKU          `json:"skus"`
		ExchangeSKUs             []AfterSaleInfoSKU          `json:"exchangeSKUs"`
		CloseReasonZh            string                      `json:"closeReasonZh"`
		ReturnsTag               int64                       `json:"returnsTag"`
		AppliedShipFeeAmountYuan float64                     `json:"appliedShipFeeAmountYuan"`
		AppliedSkusAmountYuan    float64                     `json:"appliedSkusAmountYuan"`
		ExpectedRefundAmountYuan float64                     `json:"expectedRefundAmountYuan"`
		RefundAmountYuan         float64                     `json:"refundAmountYuan"`
		RefundStatus             int64                       `json:"refundStatus"`
		CargoStatus              int64                       `json:"cargoStatus"`
	}
	AfterSaleInfoReturnAddress struct {
		Province    string `json:"province"`
		City        string `json:"city"`
		County      string `json:"county"`
		Town        string `json:"town"`
		Street      string `json:"street"`
		Phone       string `json:"phone"`
		Name        string `json:"name"`
		FullAddress string `json:"fullAddress"`
	}
	AfterSaleInfoVariants struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	}
	AfterSaleInfoSKU struct {
		SkuId                  string                  `json:"skuId"`
		SkuName                string                  `json:"skuName"`
		Image                  string                  `json:"image"`
		Price                  float64                 `json:"price"`
		BoughtCount            int64                   `json:"boughtCount"`
		AppliedCount           int64                   `json:"appliedCount"`
		AppliedTotalAmountYuan float64                 `json:"appliedTotalAmountYuan"`
		Scskucode              string                  `json:"scskucode"`
		Barcode                string                  `json:"barcode"`
		Variants               []AfterSaleInfoVariants `json:"variants"`
		SkuERPCode             string                  `json:"skuERPCode"`
	}

	RespGetAfterSaleInfoDataLogisticsInfo struct {
		AfterSale struct {
			ExpressNo          string `json:"expressNo"`
			ExpressCompanyCode string `json:"expressCompanyCode"`
			ExpressCompanyName string `json:"expressCompanyName"`
			FillExpressNoTime  int64  `json:"fillExpressNoTime"`
			ExpressSignTime    int64  `json:"expressSignTime"`
		} `json:"after_sale"`
		Exchange struct {
			ExpressNo          string `json:"expressNo"`
			ExpressCompanyCode string `json:"expressCompanyCode"`
			ExpressCompanyName string `json:"expressCompanyName"`
			FillExpressNoTime  int64  `json:"fillExpressNoTime"`
		} `json:"exchange"`
		Order struct {
			ExpressNo          string `json:"expressNo"`
			ExpressCompanyCode string `json:"expressCompanyCode"`
			ExpressCompanyName string `json:"expressCompanyName"`
		} `json:"order"`
	}

	RespGetAfterSaleInfoDataNegotiateRecords struct {
		Title            string                       `json:"title"`
		OperatorRoleName string                       `json:"operatorRoleName"`
		OperatorRole     int64                        `json:"operatorRole"`
		Attributes       []NegotiateRecordsAttributes `json:"attributes"`
		Time             string                       `json:"time"`
	}
	NegotiateRecordsAttributes struct {
		Type   int64    `json:"type"`
		Key    int64    `json:"key"`
		Desc   string   `json:"desc"`
		Photos []string `json:"photos"`
	}
)

var _ Response = new(RespGetAfterSaleInfo)

func (r RespGetAfterSaleInfo) ErrorCode() int64 {
	return r.BaseResponse.ErrorCode
}

func (r RespGetAfterSaleInfo) ErrorMsg() string {
	return r.BaseResponse.ErrorMsg

}

func (r RespGetAfterSaleInfo) Success() bool {
	return r.BaseResponse.Success
}
