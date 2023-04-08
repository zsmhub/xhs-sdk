package message

import "encoding/json"

// 文档：https://open.xiaohongshu.com/document/message/file/5/20

type MsgItemAuditReject struct {
	ItemId     string `json:"itemId"`
	UpdateTime int64  `json:"updateTime"`
}

var _ MessageData = new(MsgItemAuditReject)

func (MsgItemAuditReject) MsgTag() string {
	return "msg_item_audit_reject"
}

func (MsgItemAuditReject) DecodeData(data string) (MsgItemAuditReject, error) {
	var resp MsgItemAuditReject
	err := json.Unmarshal([]byte(data), &resp)
	return resp, err
}
