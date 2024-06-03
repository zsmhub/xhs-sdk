package message

import "encoding/json"

type MsgFulfillmentReceiverChange struct {
	OrderId       string `json:"orderId"`
	OpenAddressId string `json:"openAddressId"`
	UpdateTime    int64  `json:"updateTime"`
}

var _ MessageData = new(MsgFulfillmentReceiverChange)

func (MsgFulfillmentReceiverChange) MsgTag() string {
	return "msg_fulfillment_receiver_change"
}

func (MsgFulfillmentReceiverChange) DecodeData(data string) (MsgFulfillmentReceiverChange, error) {
	var resp MsgFulfillmentReceiverChange
	err := json.Unmarshal([]byte(data), &resp)
	return resp, err
}
