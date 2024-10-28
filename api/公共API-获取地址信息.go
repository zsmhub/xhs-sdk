package api

import json "github.com/bytedance/sonic"

type ReqGetNestZone struct {
	BaseRequest
}

var _ Request = new(ReqGetNestZone)

func (r ReqGetNestZone) Method() string {
	return "common.getNestZone"
}

func (r ReqGetNestZone) Params(base BaseRequest) []byte {
	r.BaseRequest = base
	ret, _ := json.Marshal(r)
	return ret
}

type (
	RespGetNestZone struct {
		BaseResponse
		Data RespGetNestZoneData `json:"data"`
	}
	RespGetNestZoneData struct {
		ProvinceZoneList []ZoneList `json:"provinceZoneList"`
	}
	Zone struct {
		Id         string `json:"id"`
		Code       string `json:"code"`
		Name       string `json:"name"`
		Upper      string `json:"upper"`
		Zipcode    string `json:"zipcode"`
		IsDeactive bool   `json:"isDeactive"`
	}
	ZoneList struct {
		Zone
		Zones []ZoneList `json:"zones"`
	}
)

var _ Response = new(RespGetNestZone)

func (r RespGetNestZone) ErrorCode() int64 {
	return r.BaseResponse.ErrorCode
}

func (r RespGetNestZone) ErrorMsg() string {
	return r.BaseResponse.ErrorMsg

}

func (r RespGetNestZone) Success() bool {
	return r.BaseResponse.Success
}
