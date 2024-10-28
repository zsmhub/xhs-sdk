package message

import json "github.com/bytedance/sonic"

type MsgFulfillmentSellerRemarkChange struct {
	OrderId        string `json:"orderId"`
	Operator       string `json:"operator"`
	SellerMarkNote string `json:"sellerMarkNote"`
	UpdateTime     int64  `json:"updateTime"`
}

var _ MessageData = new(MsgFulfillmentSellerRemarkChange)

func (MsgFulfillmentSellerRemarkChange) MsgTag() string {
	return "msg_fulfillment_seller_remark_change"
}

func (MsgFulfillmentSellerRemarkChange) DecodeData(data string) (MsgFulfillmentSellerRemarkChange, error) {
	var resp MsgFulfillmentSellerRemarkChange
	err := json.Unmarshal([]byte(data), &resp)
	return resp, err
}
