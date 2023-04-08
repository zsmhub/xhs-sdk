package message

import "encoding/json"

// 文档：https://open.xiaohongshu.com/document/message/file/5/21

type MsgItemDelete struct {
	ItemId     string `json:"itemId"`
	UpdateTime int64  `json:"updateTime"`
}

var _ MessageData = new(MsgItemDelete)

func (MsgItemDelete) MsgTag() string {
	return "msg_item_delete"
}

func (MsgItemDelete) DecodeData(data string) (MsgItemDelete, error) {
	var resp MsgItemDelete
	err := json.Unmarshal([]byte(data), &resp)
	return resp, err
}
