package api

import json "github.com/bytedance/sonic"

type ReqRefreshAccessToken struct {
	BaseRequest
	RefreshToken string `json:"refreshToken"`
}

var _ Request = new(ReqRefreshAccessToken)

func (r ReqRefreshAccessToken) Method() string {
	return "oauth.refreshToken"
}

func (r ReqRefreshAccessToken) Params(base BaseRequest) []byte {
	r.BaseRequest = base
	ret, _ := json.Marshal(r)
	return ret
}

type RespRefreshAccessToken struct {
	BaseResponse
	Data RespGetAccessTokenData `json:"data"`
}

var _ Response = new(RespRefreshAccessToken)

func (r RespRefreshAccessToken) ErrorCode() int64 {
	return r.BaseResponse.ErrorCode
}

func (r RespRefreshAccessToken) ErrorMsg() string {
	return r.BaseResponse.ErrorMsg

}

func (r RespRefreshAccessToken) Success() bool {
	return r.BaseResponse.Success
}
