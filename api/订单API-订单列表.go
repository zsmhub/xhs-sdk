package api

import "encoding/json"

type ReqGetOrderList struct {
	BaseRequest
	StartTime   int64 `json:"startTime"`
	EndTime     int64 `json:"endTime"`
	TimeType    int64 `json:"timeType"`
	OrderType   int64 `json:"orderType"`
	OrderStatus int64 `json:"orderStatus"`
	PageNo      int64 `json:"pageNo"`
	PageSize    int64 `json:"pageSize"`
}

var _ Request = new(ReqGetOrderList)

func (r ReqGetOrderList) Method() string {
	return "order.getOrderList"
}

func (r ReqGetOrderList) Params(base BaseRequest) []byte {
	r.BaseRequest = base
	ret, _ := json.Marshal(r)
	return ret
}

type (
	RespGetOrderList struct {
		BaseResponse
		Data RespGetOrderListData `json:"data"`
	}
	RespGetOrderListData struct {
		Total     int64 `json:"total"`
		PageNo    int64 `json:"pageNo"`
		PageSize  int64 `json:"pageSize"`
		MaxPageNo int64 `json:"maxPageNo"`
		OrderList []struct {
			OrderId                 string `json:"orderId"`
			OrderType               int64  `json:"orderType"`
			OrderStatus             int64  `json:"orderStatus"`
			OrderAfterSalesStatus   int64  `json:"orderAfterSalesStatus"`
			CancelStatus            int64  `json:"cancelStatus"`
			CreatedTime             int64  `json:"createdTime"`
			PaidTime                int64  `json:"paidTime"`
			UpdateTime              int64  `json:"updateTime"`
			DeliveryTime            int64  `json:"deliveryTime"`
			CancelTime              int64  `json:"cancelTime"`
			FinishTime              int64  `json:"finishTime"`
			PromiseLastDeliveryTime int64  `json:"promiseLastDeliveryTime"`
			PlanInfoId              string `json:"planInfoId"`
			PlanInfoName            string `json:"planInfoName"`
			ReceiverCountryId       string `json:"receiverCountryId"`
			ReceiverCountryName     string `json:"receiverCountryName"`
			ReceiverProvinceId      string `json:"receiverProvinceId"`
			ReceiverProvinceName    string `json:"receiverProvinceName"`
			ReceiverCityId          string `json:"receiverCityId"`
			ReceiverCityName        string `json:"receiverCityName"`
			ReceiverDistrictId      string `json:"receiverDistrictId"`
			ReceiverDistrictName    string `json:"receiverDistrictName"`
			CustomerRemark          string `json:"customerRemark"`
			SellerRemark            string `json:"sellerRemark"`
			SellerRemarkFlag        int64  `json:"sellerRemarkFlag"`
			OriginalOrderId         string `json:"originalOrderId"`
			Logistics               string `json:"logistics"`
			TotalDepositAmount      int64  `json:"totalDepositAmount"`
			TotalMerchantDiscount   int64  `json:"totalMerchantDiscount"`
			TotalRedDiscount        int64  `json:"totalRedDiscount"`
			PaymentType             int64  `json:"paymentType"`
			UserId                  string `json:"userId"`
		} `json:"orderList"`
	}
)

var _ Response = new(RespGetOrderList)

func (r RespGetOrderList) ErrorCode() int64 {
	return r.BaseResponse.ErrorCode
}

func (r RespGetOrderList) ErrorMsg() string {
	return r.BaseResponse.ErrorMsg

}

func (r RespGetOrderList) Success() bool {
	return r.BaseResponse.Success
}
