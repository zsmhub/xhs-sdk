package message

import "encoding/json"

// 文档：https://open.xiaohongshu.com/document/message/file/6/24

type MsgFulfillmentSellerRemarkChange struct {
	PackageId      string `json:"packageId"`
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
