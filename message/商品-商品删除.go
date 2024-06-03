package message

import "encoding/json"

type MsgSkuDelete struct {
	SkuId      string `json:"skuId"`
	UpdateTime int64  `json:"updateTime"`
}

var _ MessageData = new(MsgSkuDelete)

func (MsgSkuDelete) MsgTag() string {
	return "msg_sku_delete"
}

func (MsgSkuDelete) DecodeData(data string) (MsgSkuDelete, error) {
	var resp MsgSkuDelete
	err := json.Unmarshal([]byte(data), &resp)
	return resp, err
}
