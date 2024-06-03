package api

import "encoding/json"

type ReqGetOrderDetail struct {
	BaseRequest
	OrderId string `json:"orderId"`
}

var _ Request = new(ReqGetOrderDetail)

func (r ReqGetOrderDetail) Method() string {
	return "order.getOrderDetail"
}

func (r ReqGetOrderDetail) Params(base BaseRequest) []byte {
	r.BaseRequest = base
	ret, _ := json.Marshal(r)
	return ret
}

type (
	RespGetOrderDetail struct {
		BaseResponse
		Data RespGetOrderDetailData `json:"data"`
	}
	RespGetOrderDetailData struct {
		OrderId                  string         `json:"orderId"`
		OrderType                int64          `json:"orderType"`
		OrderStatus              int64          `json:"orderStatus"`
		OrderAfterSalesStatus    int64          `json:"orderAfterSalesStatus"`
		CancelStatus             int64          `json:"cancelStatus"`
		CreatedTime              int64          `json:"createdTime"`
		PaidTime                 int64          `json:"paidTime"`
		UpdateTime               int64          `json:"updateTime"`
		DeliveryTime             int64          `json:"deliveryTime"`
		CancelTime               int64          `json:"cancelTime"`
		FinishTime               int64          `json:"finishTime"`
		PromiseLastDeliveryTime  int64          `json:"promiseLastDeliveryTime"`
		PlanInfoId               string         `json:"planInfoId"`
		PlanInfoName             string         `json:"planInfoName"`
		ReceiverCountryId        string         `json:"receiverCountryId"`
		ReceiverCountryName      string         `json:"receiverCountryName"`
		ReceiverProvinceId       string         `json:"receiverProvinceId"`
		ReceiverProvinceName     string         `json:"receiverProvinceName"`
		ReceiverCityId           string         `json:"receiverCityId"`
		ReceiverCityName         string         `json:"receiverCityName"`
		ReceiverDistrictId       string         `json:"receiverDistrictId"`
		ReceiverDistrictName     string         `json:"receiverDistrictName"`
		CustomerRemark           string         `json:"customerRemark"`
		SellerRemark             string         `json:"sellerRemark"`
		SellerRemarkFlag         int64          `json:"sellerRemarkFlag"`
		PresaleDeliveryStartTime int64          `json:"presaleDeliveryStartTime"`
		PresaleDeliveryEndTime   int64          `json:"presaleDeliveryEndTime"`
		SkuList                  []OrderSkuList `json:"skuList"`
		OriginalOrderId          string         `json:"originalOrderId"`
		TotalNetWeightAmount     int64          `json:"totalNetWeightAmount"`
		TotalPayAmount           int64          `json:"totalPayAmount"`
		TotalShippingFree        int64          `json:"totalShippingFree"`
		Unpack                   bool           `json:"unpack"`
		ExpressTrackingNo        string         `json:"expressTrackingNo"`
		ExpressCompanyCode       string         `json:"expressCompanyCode"`
		ReceiverName             string         `json:"receiverName"`
		ReceiverPhone            string         `json:"receiverPhone"`
		ReceiverAddress          string         `json:"receiverAddress"`
		BoundExtendInfo          struct {
			PayNo          string   `json:"payNo"`
			PayChannel     string   `json:"payChannel"`
			ProductValue   float64  `json:"productValue"`
			PayAmount      float64  `json:"payAmount"`
			TaxAmount      float64  `json:"taxAmount"`
			ShippingFee    float64  `json:"shippingFee"`
			DiscountAmount float64  `json:"discountAmount"`
			ZoneCodes      []string `json:"zoneCodes"`
		} `json:"boundExtendInfo"`
		TransferExtendInfo struct {
			InternationalExpressNo string  `json:"internationalExpressNo"`
			OrderDeclaredAmount    float64 `json:"orderDeclaredAmount"`
			PaintMarker            string  `json:"paintMarker"`
			CollectionPlace        string  `json:"collectionPlace"`
			ThreeSegmentCode       string  `json:"threeSegmentCode"`
		} `json:"transferExtendInfo"`
		OpenAddressId           string `json:"openAddressId"`
		SimpleDeliveryOrderList []struct {
			DeliveryOrderIndex int64    `json:"deliveryOrderIndex"`
			Status             int64    `json:"status"`
			ExpressTrackingNo  string   `json:"expressTrackingNo"`
			ExpressCompanyCode string   `json:"expressCompanyCode"`
			SkuIdList          []string `json:"skuIdList"`
		} `json:"simpleDeliveryOrderList"`
		Logistics string `json:"logistics"`
	}
	OrderSkuList struct {
		SkuId                 string               `json:"skuId"`
		SkuName               string               `json:"skuName"`
		Erpcode               string               `json:"erpcode"`
		SkuSpec               string               `json:"skuSpec"`
		SkuImage              string               `json:"skuImage"`
		SkuQuantity           int64                `json:"skuQuantity"`
		SkuDetailList         []OrderSkuDetailList `json:"skuDetailList"`
		TotalPaidAmount       int64                `json:"totalPaidAmount"`
		TotalMerchantDiscount int64                `json:"totalMerchantDiscount"`
		TotalRedDiscount      int64                `json:"totalRedDiscount"`
		TotalTaxAmount        int64                `json:"totalTaxAmount"`
		TotalNetWeight        int64                `json:"totalNetWeight"`
		SkuTag                int64                `json:"skuTag"`
		IsChannel             bool                 `json:"isChannel"`
		Channel               bool                 `json:"channel"`
	}
	OrderSkuDetailList struct {
		SkuId                  string `json:"skuId"`
		ErpCode                string `json:"erpCode"`
		Barcode                string `json:"barcode"`
		ScSkuCode              string `json:"scSkuCode"`
		Quantity               int64  `json:"quantity"`
		RegisterName           string `json:"registerName"`
		SkuName                string `json:"skuName"`
		PricePerSku            int64  `json:"pricePerSku"`
		TaxPerSku              int64  `json:"taxPerSku"`
		PaidAmountPerSku       int64  `json:"paidAmountPerSku"`
		DepositAmountPerSku    int64  `json:"depositAmountPerSku"`
		MerchantDiscountPerSku int64  `json:"merchantDiscountPerSku"`
		RedDiscountPerSku      int64  `json:"redDiscountPerSku"`
		RawPricePerSku         int64  `json:"rawPricePerSku"`
	}
)

var _ Response = new(RespGetOrderDetail)

func (r RespGetOrderDetail) ErrorCode() int64 {
	return r.BaseResponse.ErrorCode
}

func (r RespGetOrderDetail) ErrorMsg() string {
	return r.BaseResponse.ErrorMsg

}

func (r RespGetOrderDetail) Success() bool {
	return r.BaseResponse.Success
}
