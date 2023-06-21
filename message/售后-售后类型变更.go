package message

import "encoding/json"

// 文档：https://open.xiaohongshu.com/document/message/file/7/27

type MsgAfterSaleTransfer struct {
	ReturnsId   string  `json:"returnsId"`
	OrderId     string  `json:"orderId"`
	ReturnType  int64   `json:"returnType"`
	RequestFrom int64   `json:"request_from"`
	RefundFee   float64 `json:"refundFee"`
	UpdateTime  int64   `json:"updateTime"`
}

var _ MessageData = new(MsgAfterSaleTransfer)

func (MsgAfterSaleTransfer) MsgTag() string {
	return "msg_after_sale_transfer"
}

func (MsgAfterSaleTransfer) DecodeData(data string) (MsgAfterSaleTransfer, error) {
	var resp MsgAfterSaleTransfer
	err := json.Unmarshal([]byte(data), &resp)
	return resp, err
}
