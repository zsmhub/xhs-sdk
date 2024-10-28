package xhs_sdk

import (
	"github.com/mryee2023/xhs-sdk/api"
)

// 小红书开放平台 sdk 调用入口

// 业务接口调用客户端
func NewXhsClient(appId, appSecret string) *api.ApiClient {
	return &api.ApiClient{
		AppId:       appId,
		AppSecret:   appSecret,
		AccessToken: "",
	}
}

// 授权接口调用客户端，如获取AccessToken
func NewXhsOAuthClient(appId, appSecret string) *api.OAuthClient {
	return &api.OAuthClient{
		AppId:     appId,
		AppSecret: appSecret,
	}
}
