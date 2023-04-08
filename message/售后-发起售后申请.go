package message

import "encoding/json"

// 文档：https://open.xiaohongshu.com/document/message/file/7/26

type MsgAfterSaleCreate struct {
	ReturnsId   string  `json:"returnsId"`
	PackageId   string  `json:"packageId"`
	ReturnType  int64   `json:"returnType"`
	RequestFrom int64   `json:"request_from"`
	RefundFee   float64 `json:"refundFee"`
	UpdateTime  int64   `json:"updateTime"`
}

var _ MessageData = new(MsgAfterSaleCreate)

func (MsgAfterSaleCreate) MsgTag() string {
	return "msg_after_sale_create"
}

func (MsgAfterSaleCreate) DecodeData(data string) (MsgAfterSaleCreate, error) {
	var resp MsgAfterSaleCreate
	err := json.Unmarshal([]byte(data), &resp)
	return resp, err
}
