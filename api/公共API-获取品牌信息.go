package api

import json "github.com/bytedance/sonic"

type ReqBrandSearch struct {
	BaseRequest
	CategoryId string `json:"categoryId"`
	Keyword    string `json:"keyword"`
	PageNo     int64  `json:"pageNo"`
	PageSize   int64  `json:"pageSize"`
}

var _ Request = new(ReqBrandSearch)

func (r ReqBrandSearch) Method() string {
	return "common.brandSearch"
}

func (r ReqBrandSearch) Params(base BaseRequest) []byte {
	r.BaseRequest = base
	ret, _ := json.Marshal(r)
	return ret
}

type (
	RespBrandSearch struct {
		BaseResponse
		Data RespBrandSearchData `json:"data"`
	}
	RespBrandSearchData struct {
		Brands []struct {
			Name   string `json:"name"`
			EnName string `json:"enName"`
			Id     string `json:"id"`
			Image  string `json:"image"`
		} `json:"brands"`
	}
)

var _ Response = new(RespBrandSearch)

func (r RespBrandSearch) ErrorCode() int64 {
	return r.BaseResponse.ErrorCode
}

func (r RespBrandSearch) ErrorMsg() string {
	return r.BaseResponse.ErrorMsg

}

func (r RespBrandSearch) Success() bool {
	return r.BaseResponse.Success
}
