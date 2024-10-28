package api

import json "github.com/bytedance/sonic"

type ReqGetZones struct {
	BaseRequest
	Code                 string `json:"code"`
	Name                 string `json:"name"`
	Upper                string `json:"upper"`
	FilterNonContinental string `json:"filterNonContinental"`
}

var _ Request = new(ReqGetZones)

func (r ReqGetZones) Method() string {
	return "common.getZones"
}

func (r ReqGetZones) Params(base BaseRequest) []byte {
	r.BaseRequest = base
	ret, _ := json.Marshal(r)
	return ret
}

type (
	RespGetZones struct {
		BaseResponse
		Data []RespGetZonesData `json:"data"`
	}
	RespGetZonesData struct {
		Name          string `json:"name"`
		ShortName     string `json:"shortName"`
		Code          string `json:"code"`
		Upper         string `json:"upper"`
		InitialPinyin string `json:"initialPinyin"`
		Pinyin        string `json:"pinyin"`
		IsActive      string `json:"isActive"`
	}
)

var _ Response = new(RespGetZones)

func (r RespGetZones) ErrorCode() int64 {
	return r.BaseResponse.ErrorCode
}

func (r RespGetZones) ErrorMsg() string {
	return r.BaseResponse.ErrorMsg

}

func (r RespGetZones) Success() bool {
	return r.BaseResponse.Success
}
