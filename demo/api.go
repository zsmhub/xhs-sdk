package demo

import (
	"encoding/json"
	"fmt"
	xhs_sdk "github.com/zsmhub/xhs-sdk"
	"github.com/zsmhub/xhs-sdk/api"
	"time"
)

// 调用 小红书开放平台API 示例
func GetOAuthUrl(redirectUri, state string) {
	authUrl := xhs_sdk.NewXhsOAuthClient(AppKey, AppSecret).GetOAuthUrl(redirectUri, state)
	fmt.Printf("authUrl=%s\n", authUrl)
}

func GetAccessToken() {
	accessToken, err := xhs_sdk.NewXhsOAuthClient(AppKey, AppSecret).GetAccessToken(Code)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", accessToken)
}

func RefreshAccessToken() {
	accessToken, err := xhs_sdk.NewXhsOAuthClient(AppKey, AppSecret).RefreshAccessToken(RefreshToken)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", accessToken)
}

func GetExpressCompanyList() {
	var resp api.RespGetExpressCompanyList
	err := xhs_sdk.NewXhsClient(AppKey, AppSecret).SetAccessToken(AccessToken).Post(api.ReqGetExpressCompanyList{}, &resp)
	if err != nil {
		fmt.Println(err)
		return
	}
	ret, _ := json.Marshal(resp)
	fmt.Printf("%s\n", string(ret))
}

func GetZones() {
	var resp api.RespGetZones
	err := xhs_sdk.NewXhsClient(AppKey, AppSecret).SetAccessToken(AccessToken).Post(api.ReqGetZones{}, &resp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", resp)
}

func GetNestZone() {
	var resp api.RespGetNestZone
	err := xhs_sdk.NewXhsClient(AppKey, AppSecret).SetAccessToken(AccessToken).Post(api.ReqGetNestZone{}, &resp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", resp)
}

func GetOrderList() {
	var (
		now = time.Now()
		req = api.ReqGetOrderList{
			StartTime:   now.AddDate(0, 0, -1).Unix(),
			EndTime:     now.Unix(),
			TimeType:    xhs_sdk.OrderTimeType1,
			OrderType:   0,
			OrderStatus: 0,
			PageNo:      0,
			PageSize:    0,
		}
		resp api.RespGetOrderList
	)
	err := xhs_sdk.NewXhsClient(AppKey, AppSecret).SetAccessToken(AccessToken).Post(req, &resp)
	if err != nil {
		fmt.Println(err)
		return
	}
	ret, _ := json.Marshal(resp)
	fmt.Printf("%s\n", string(ret))
}

func GetOrderDetail() {
	var (
		req = api.ReqGetOrderDetail{
			OrderId: "P680835140254578241",
		}
		resp api.RespGetOrderDetail
	)
	err := xhs_sdk.NewXhsClient(AppKey, AppSecret).SetAccessToken(AccessToken).Post(req, &resp)
	if err != nil {
		fmt.Println(err)
		return
	}
	ret, _ := json.Marshal(resp)
	fmt.Printf("%s\n", string(ret))
}

func GetOrderReceiverInfo() {
	var (
		req = api.ReqGetOrderReceiverInfo{
			ReceiverQueries: []api.ReceiverQuery{
				{
					OrderId:       "P679967014036803561",
					OpenAddressId: "1ab009592728cc8a813ad1751aed0742",
				},
			},
			IsReturn: false,
		}
		resp api.RespGetOrderReceiverInfo
	)
	err := xhs_sdk.NewXhsClient(AppKey, AppSecret).SetAccessToken(AccessToken).Post(req, &resp)
	if err != nil {
		fmt.Println(err)
		return
	}
	ret, _ := json.Marshal(resp)
	fmt.Printf("%s\n", string(ret))
}

func BatchDecrypt() {
	var (
		req = api.ReqBatchDecrypt{
			BaseInfos: []api.BaseInfo{
				{
					DataTag:       "P679967014036803561",
					EncryptedData: "#n7hSLUW/WbyCapKBdmPTMPQMEMK8jMA1OfCC+6xi3M0=#n7hSLUW/WbyCapKBdmPTMFGCfwbFFQPiWQZdDDdNwqeLqUyVAEbNERmOw6S6ItW4HMjDAgM5DRYeKy9Afhm2rzJs3cU7Whd0S+5Aa+ZncFs=#2##",
				},
				{
					DataTag:       "P679967014036803561",
					EncryptedData: "#n7hSLUW/WbyCapKBdmPTMDA/KEFoYD+xDLLUdFfANd8=#n7hSLUW/WbyCapKBdmPTMFGCfwbFFQPiWQZdDDdNwqeLqUyVAEbNERmOw6S6ItW4TPaqnWzi5ceHtdUzCS+QUpXa9uhe3zhEq5s+FgppUVg=#3##",
				},
				{
					DataTag:       "P679967014036803561",
					EncryptedData: "#n7hSLUW/WbyCapKBdmPTMMJkWMc3EV7b4lY4EcBipU45y8I2GU799DebtB+K4G6MWs8wUgIKDpSaQhGoxxmvEi6UdCJAIq7j3YH6WvMkavrTgFwb3rfJjSTpCk0+ZRi+#n7hSLUW/WbyCapKBdmPTMFGCfwbFFQPiWQZdDDdNwqeLqUyVAEbNERmOw6S6ItW4HMjDAgM5DRYeKy9Afhm2r5BzTq+Nbdi32Tu/sUV4A+KjBQ8TXPfZ9ygC6ZR3eDCrw960sG8cCr/biBu9xFC4p9l1b4VINAMnArNCXs9bD78=#1##",
				},
			},
			ActionType: xhs_sdk.DecryptActionType1,
			AppUserId:  "",
		}
		resp api.RespBatchDecrypt
	)
	err := xhs_sdk.NewXhsClient(AppKey, AppSecret).SetAccessToken(AccessToken).Post(req, &resp)
	if err != nil {
		fmt.Println(err)
		return
	}
	ret, _ := json.Marshal(resp)
	fmt.Printf("%s\n", string(ret))
}

func GetAfterSaleDetail() {
	var (
		req = api.ReqGetAfterSaleDetail{
			AfterSaleId: "R4889848165315631",
		}
		resp api.RespGetAfterSaleDetail
	)
	err := xhs_sdk.NewXhsClient(AppKey, AppSecret).SetAccessToken(AccessToken).Post(req, &resp)
	if err != nil {
		fmt.Println(err)
		return
	}
	ret, _ := json.Marshal(resp)
	fmt.Printf("%s\n", string(ret))
}
