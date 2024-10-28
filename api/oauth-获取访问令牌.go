package api

import json "github.com/bytedance/sonic"

type ReqGetAccessToken struct {
	BaseRequest
	Code string `json:"code"`
}

var _ Request = new(ReqGetAccessToken)

func (r ReqGetAccessToken) Method() string {
	return "oauth.getAccessToken"
}

func (r ReqGetAccessToken) Params(base BaseRequest) []byte {
	r.BaseRequest = base
	ret, _ := json.Marshal(r)
	return ret
}

type (
	RespGetAccessToken struct {
		BaseResponse
		Data RespGetAccessTokenData `json:"data"`
	}

	RespGetAccessTokenData struct {
		AccessToken           string `json:"accessToken"`
		AccessTokenExpiresAt  int64  `json:"accessTokenExpiresAt"`
		RefreshToken          string `json:"refreshToken"`
		RefreshTokenExpiresAt int64  `json:"refreshTokenExpiresAt"`
		SellerId              string `json:"sellerId"`
		SellerName            string `json:"sellerName"`
	}
)

var _ Response = new(RespGetAccessToken)

func (r RespGetAccessToken) ErrorCode() int64 {
	return r.BaseResponse.ErrorCode
}

func (r RespGetAccessToken) ErrorMsg() string {
	return r.BaseResponse.ErrorMsg

}

func (r RespGetAccessToken) Success() bool {
	return r.BaseResponse.Success
}
