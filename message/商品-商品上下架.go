package message

import "encoding/json"

// 文档：https://open.xiaohongshu.com/document/message/file/5/19

type MsgItemBuyable struct {
	ItemId     string `json:"itemId"`
	UpdateTime int64  `json:"updateTime"`
}

var _ MessageData = new(MsgItemBuyable)

func (MsgItemBuyable) MsgTag() string {
	return "msg_item_buyable"
}

func (MsgItemBuyable) DecodeData(data string) (MsgItemBuyable, error) {
	var resp MsgItemBuyable
	err := json.Unmarshal([]byte(data), &resp)
	return resp, err
}
