package message

import (
	json "github.com/bytedance/sonic"
)

type MessageData interface {
	MsgTag() string
}

type Message struct {
	MsgTag   string `json:"msgTag"`
	SellerId string `json:"sellerId"`
	Data     string `json:"data"`
}

func (m *Message) DecodeData(data MessageData) error {
	return json.Unmarshal([]byte(m.Data), &data)
}

type MessageResp struct {
	Success   bool   `json:"success"`
	ErrorCode int    `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
}
