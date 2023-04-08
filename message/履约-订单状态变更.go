package message

import "encoding/json"

// 文档：https://open.xiaohongshu.com/document/message/file/6/22

type MsgFulfillmentStatusChange struct {
	PackageId     string `json:"packageId"`
	UpdateTime    int64  `json:"updateTime"`
	PackageStatus int64  `json:"packageStatus"`
}

var _ MessageData = new(MsgFulfillmentStatusChange)

func (MsgFulfillmentStatusChange) MsgTag() string {
	return "msg_fulfillment_status_change"
}

func (MsgFulfillmentStatusChange) DecodeData(data string) (MsgFulfillmentStatusChange, error) {
	var resp MsgFulfillmentStatusChange
	err := json.Unmarshal([]byte(data), &resp)
	return resp, err
}
