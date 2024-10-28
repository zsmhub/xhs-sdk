package api

import json "github.com/bytedance/sonic"

type (
	ReqBatchDesensitise struct {
		BaseRequest
		BaseInfos  []BaseInfo `json:"baseInfos"`
		ActionType string     `json:"actionType"`
		AppUserId  string     `json:"appUserId"`
	}
)

var _ Request = new(ReqBatchDesensitise)

func (r ReqBatchDesensitise) Method() string {
	return "data.batchDesensitise"
}

func (r ReqBatchDesensitise) Params(base BaseRequest) []byte {
	r.BaseRequest = base
	ret, _ := json.Marshal(r)
	return ret
}

type (
	RespBatchDesensitise struct {
		BaseResponse
		Data RespBatchDesensitiseData `json:"data"`
	}
	RespBatchDesensitiseData struct {
		DesensitiseInfoList []struct {
			DataTag          string `json:"dataTag"`
			EncryptedData    string `json:"encryptedData"`
			DesensitisedData string `json:"desensitisedData"`
			ErrorCode        int64  `json:"errorCode"`
			ErrorMsg         string `json:"errorMsg"`
		} `json:"desensitiseInfoList"`
	}
)

var _ Response = new(RespBatchDesensitise)

func (r RespBatchDesensitise) ErrorCode() int64 {
	return r.BaseResponse.ErrorCode
}

func (r RespBatchDesensitise) ErrorMsg() string {
	return r.BaseResponse.ErrorMsg

}

func (r RespBatchDesensitise) Success() bool {
	return r.BaseResponse.Success
}
