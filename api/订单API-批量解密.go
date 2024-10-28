package api

import json "github.com/bytedance/sonic"

type (
	ReqBatchDecrypt struct {
		BaseRequest
		BaseInfos  []BaseInfo `json:"baseInfos"`
		ActionType string     `json:"actionType"` // 操作类型 1-单个查看订单明文，2-批量解密打单，3-批量解密后面向三方的数据下发，4-其他场景
		AppUserId  string     `json:"appUserId"`  // 三方操作id，服务商自定义，解密接口必填
	}
	BaseInfo struct {
		DataTag       string `json:"dataTag"` // 标签，订单场景为orderId
		EncryptedData string `json:"encryptedData"`
	}
)

var _ Request = new(ReqBatchDecrypt)

func (r ReqBatchDecrypt) Method() string {
	return "data.batchDecrypt"
}

func (r ReqBatchDecrypt) Params(base BaseRequest) []byte {
	r.BaseRequest = base
	ret, _ := json.Marshal(r)
	return ret
}

type (
	RespBatchDecrypt struct {
		BaseResponse
		Data RespBatchDecryptData `json:"data"`
	}
	RespBatchDecryptData struct {
		DataInfoList []struct {
			DataTag       string `json:"dataTag"`
			EncryptedData string `json:"encryptedData"`
			DecryptedData string `json:"decryptedData"`
			ErrorCode     int64  `json:"errorCode"`
			ErrorMsg      string `json:"errorMsg"`
		} `json:"dataInfoList"`
	}
)

var _ Response = new(RespBatchDecrypt)

func (r RespBatchDecrypt) ErrorCode() int64 {
	return r.BaseResponse.ErrorCode
}

func (r RespBatchDecrypt) ErrorMsg() string {
	return r.BaseResponse.ErrorMsg

}

func (r RespBatchDecrypt) Success() bool {
	return r.BaseResponse.Success
}
