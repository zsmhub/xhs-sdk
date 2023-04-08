package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"time"
)

type OAuthClient struct {
	AppId     string
	AppSecret string
}

// 获取授权码链接
func (c *OAuthClient) GetOAuthUrl(redirectUri, state string) string {
	return fmt.Sprintf("https://ark.xiaohongshu.com/ark/authorization?appId=%s&redirectUri=%s&state=%s",
		c.AppId, url.QueryEscape(redirectUri), state)
}

// 获取访问令牌
func (c *OAuthClient) GetAccessToken(code string) (RespGetAccessToken, error) {
	var (
		req  ReqGetAccessToken
		resp RespGetAccessToken
	)

	req.Code = code

	// 构造公共参数
	baseReq := BaseRequest{
		AppId:     c.AppId,
		Timestamp: time.Now().Unix(),
		Version:   Version,
		Method:    req.Method(),
	}
	baseReq.Sign = MakeSign(baseReq.Method, baseReq.AppId, c.AppSecret, baseReq.Version, baseReq.Timestamp)

	respBody, err := Post(req.Params(baseReq))
	if err != nil {
		return resp, err
	}

	if err = json.Unmarshal(respBody, &resp); err != nil {
		return resp, err
	}

	if resp.Success() && resp.ErrorCode() == SuccessCode {
		return resp, nil
	}

	return resp, errors.New(resp.ErrorMsg())
}

// 刷新访问令牌
func (c *OAuthClient) RefreshAccessToken(refreshToken string) (RespGetAccessToken, error) {
	var (
		req  ReqRefreshAccessToken
		resp RespGetAccessToken
	)

	req.RefreshToken = refreshToken

	// 构造公共参数
	baseReq := BaseRequest{
		AppId:     c.AppId,
		Timestamp: time.Now().Unix(),
		Version:   Version,
		Method:    req.Method(),
	}
	baseReq.Sign = MakeSign(baseReq.Method, baseReq.AppId, c.AppSecret, baseReq.Version, baseReq.Timestamp)

	respBody, err := Post(req.Params(baseReq))
	if err != nil {
		return resp, err
	}

	if err = json.Unmarshal(respBody, &resp); err != nil {
		return resp, err
	}

	if resp.Success() && resp.ErrorCode() == SuccessCode {
		return resp, nil
	}

	return resp, errors.New(resp.ErrorMsg())
}
