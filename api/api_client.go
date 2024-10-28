package api

import (
	"errors"
	"fmt"
	"time"

	json "github.com/bytedance/sonic"
)

// API调用客户端
type ApiClient struct {
	AppId       string
	AppSecret   string
	AccessToken string
}

func (c *ApiClient) SetAccessToken(token string) *ApiClient {
	c.AccessToken = token
	return c
}

func (c *ApiClient) PostV2(req Request) (Response, error) {
	if c.AccessToken == "" {
		return nil, errors.New("access_token不能为空")
	}

	// 构造公共参数
	baseReq := BaseRequest{
		AppId:       c.AppId,
		Timestamp:   time.Now().Unix(),
		AccessToken: c.AccessToken,
		Version:     Version,
		Method:      req.Method(),
	}
	baseReq.Sign = MakeSign(baseReq.Method, baseReq.AppId, c.AppSecret, baseReq.Version, baseReq.Timestamp)

	respBody, err := Post(req.Params(baseReq))
	if err != nil {
		return nil, err
	}
	var resp Response
	if err = json.Unmarshal(respBody, &resp); err != nil {
		return nil, err
	}

	if resp.Success() && resp.ErrorCode() == SuccessCode {
		return resp, nil
	}

	errMsg := fmt.Sprintf("xhs: error_code=%d, error_msg=%s", resp.ErrorCode(), resp.ErrorMsg())
	return nil, errors.New(errMsg)
}

func (c *ApiClient) Post(req Request, resp Response) error {
	if c.AccessToken == "" {
		return errors.New("access_token不能为空")
	}

	// 构造公共参数
	baseReq := BaseRequest{
		AppId:       c.AppId,
		Timestamp:   time.Now().Unix(),
		AccessToken: c.AccessToken,
		Version:     Version,
		Method:      req.Method(),
	}
	baseReq.Sign = MakeSign(baseReq.Method, baseReq.AppId, c.AppSecret, baseReq.Version, baseReq.Timestamp)

	respBody, err := Post(req.Params(baseReq))
	if err != nil {
		return err
	}

	if err = json.Unmarshal(respBody, &resp); err != nil {
		return err
	}

	if resp.Success() && resp.ErrorCode() == SuccessCode {
		return nil
	}

	errMsg := fmt.Sprintf("xhs: error_code=%d, error_msg=%s", resp.ErrorCode(), resp.ErrorMsg())
	return errors.New(errMsg)
}
