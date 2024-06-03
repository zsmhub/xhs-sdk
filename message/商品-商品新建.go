package message

import "encoding/json"

type MsgSkuCreate struct {
	SkuId      string `json:"skuId"`
	UpdateTime int64  `json:"updateTime"`
}

var _ MessageData = new(MsgSkuCreate)

func (MsgSkuCreate) MsgTag() string {
	return "msg_sku_create"
}

func (MsgSkuCreate) DecodeData(data string) (MsgSkuCreate, error) {
	var resp MsgSkuCreate
	err := json.Unmarshal([]byte(data), &resp)
	return resp, err
}
