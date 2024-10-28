package message

import json "github.com/bytedance/sonic"

type MsgFulfillmentDeliverTimeChange struct {
	OrderId       string `json:"orderId"`
	OpenAddressId int64  `json:"openAddressId"`
	UpdateTime    int64  `json:"updateTime"`
}

var _ MessageData = new(MsgFulfillmentDeliverTimeChange)

func (MsgFulfillmentDeliverTimeChange) MsgTag() string {
	return "msg_fulfillment_delivery_time_change"
}

func (MsgFulfillmentDeliverTimeChange) DecodeData(data string) (MsgFulfillmentDeliverTimeChange, error) {
	var resp MsgFulfillmentDeliverTimeChange
	err := json.Unmarshal([]byte(data), &resp)
	return resp, err
}
