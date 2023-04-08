package message

import "encoding/json"

// 文档：https://open.xiaohongshu.com/document/message/file/5/18

type MsgItemCreate struct {
	ItemId     string `json:"itemId"`
	UpdateTime int64  `json:"updateTime"`
}

var _ MessageData = new(MsgItemCreate)

func (MsgItemCreate) MsgTag() string {
	return "msg_item_create"
}

func (MsgItemCreate) DecodeData(data string) (MsgItemCreate, error) {
	var resp MsgItemCreate
	err := json.Unmarshal([]byte(data), &resp)
	return resp, err
}
